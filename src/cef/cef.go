// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package cef

/*
CEF capi fixes
--------------

In cef_string.h:
this => typedef cef_string_utf16_t cef_string_t;
to => #define cef_string_t cef_string_utf16_t

*/

/*
#cgo CFLAGS: -I./../../
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
*/
import "C"
import "unsafe"
import "os"

type Settings struct {
    CachePath string
    LogSeverity int
}

type BrowserSettings struct {
}

var g_mainArgs C.struct__cef_main_args_t
var g_app C.cef_app_t // needs reference counting
var g_clientHandler C.struct__cef_client_t // needs reference counting

// Sandbox is disabled. Including the "cef_sandbox.lib"
// library results in lots of GCC warnings/errors. It is
// compatible only with VS 2010. It would be required to
// build it using GCC. Add -lcef_sandbox to LDFLAGS.
// capi doesn't expose sandbox functions, you need do add
// these before import "C":
// void* cef_sandbox_info_create();
// void cef_sandbox_info_destroy(void* sandbox_info);
var g_sandboxInfo unsafe.Pointer

const (
    LOGSEVERITY_DEFAULT = C.LOGSEVERITY_DEFAULT
    LOGSEVERITY_VERBOSE = C.LOGSEVERITY_VERBOSE
    LOGSEVERITY_INFO = C.LOGSEVERITY_INFO
    LOGSEVERITY_WARNING = C.LOGSEVERITY_WARNING
    LOGSEVERITY_ERROR = C.LOGSEVERITY_ERROR
    LOGSEVERITY_ERROR_REPORT = C.LOGSEVERITY_ERROR_REPORT
    LOGSEVERITY_DISABLE = C.LOGSEVERITY_DISABLE
)

func ExecuteProcess(appHandle unsafe.Pointer) {
    FillMainArgs(&g_mainArgs, appHandle)

    // Sandbox info needs to be passed to both cef_execute_process()
    // and cef_initialize().
    // OFF: g_sandboxInfo = C.cef_sandbox_info_create()

    var exitCode C.int = C.cef_execute_process(&g_mainArgs, nil,
            g_sandboxInfo)
    if (exitCode >= 0) {
        os.Exit(int(exitCode))
    }
}

func Initialize(settings Settings) int {
    var cefSettings C.struct__cef_settings_t

    // cache_path
    var cachePath *C.char = C.CString(settings.CachePath)
    defer C.free(unsafe.Pointer(cachePath))
    C.cef_string_from_utf8(cachePath, C.strlen(cachePath),
            &cefSettings.cache_path)

    // log_severity
    cefSettings.log_severity =
            (C.cef_log_severity_t)(C.int(settings.LogSeverity))

    ret := C.cef_initialize(&g_mainArgs, &cefSettings, nil, g_sandboxInfo)
    return int(ret)
}

func CreateBrowser(hwnd unsafe.Pointer, settings BrowserSettings, 
        url string) {
    // windowInfo
    var windowInfo C.cef_window_info_t
    FillWindowInfo(&windowInfo, hwnd)
    
    // url
    var cefUrl C.cef_string_t
    var charUrl *C.char = C.CString(url)
    defer C.free(unsafe.Pointer(charUrl))
    C.cef_string_from_utf8(charUrl, C.strlen(charUrl), &cefUrl)

    // create browser
    var cefSettings C.struct__cef_browser_settings_t
    
    // TODO: reference counting, see:
    // https://code.google.com/p/chromiumembedded/wiki/UsingTheCAPI
    // --
    // What about C struct alignment issues in Go?
    // var refcnt unsafe.Pointer = unsafe.Pointer(&g_clientHandler)
    // refcnt += (unsafe.Pointer)(unsafe.Sizeof(g_clientHandler[0]))
    // C.InterlockedIncrement(refcnt)
    // C.InterlockedIncrement(&g_clientHandler{0})
    // C.InterlockedDecrement()
    // --

    // Must call synchronously so that a call to WindowResize()
    // works, after this function returns.
    C.cef_browser_host_create_browser_sync(&windowInfo, nil, &cefUrl,
            &cefSettings, nil)
}

func RunMessageLoop() {
    C.cef_run_message_loop()
}

func QuitMessageLoop() {
    C.cef_quit_message_loop()
}

func Shutdown() {
    C.cef_shutdown()
    // OFF: cef_sandbox_info_destroy(g_sandboxInfo)
}
