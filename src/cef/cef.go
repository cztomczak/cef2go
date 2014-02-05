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
#cgo LDFLAGS: -L./../../Release -llibcef -lcef_sandbox
#include <stdlib.h>
#include "include/capi/cef_app_capi.h"
*/
import "C"
import "unsafe"

type Settings struct {
    CachePath string
    LogSeverity int
}

const (
    LOGSEVERITY_DEFAULT = C.LOGSEVERITY_DEFAULT
    LOGSEVERITY_VERBOSE = C.LOGSEVERITY_VERBOSE
    LOGSEVERITY_INFO = C.LOGSEVERITY_INFO
    LOGSEVERITY_WARNING = C.LOGSEVERITY_WARNING
    LOGSEVERITY_ERROR = C.LOGSEVERITY_ERROR
    LOGSEVERITY_ERROR_REPORT = C.LOGSEVERITY_ERROR_REPORT
    LOGSEVERITY_DISABLE = C.LOGSEVERITY_DISABLE
)

func Initialize(settings Settings) {
    var mainArgs C.struct__cef_main_args_t
    var cefSettings C.struct__cef_settings_t
    var app C.cef_app_t
    var sandbox unsafe.Pointer

    // cache_path
    var cachePath *C.char = C.CString(settings.CachePath)
    defer C.free(unsafe.Pointer(cachePath))
    C.cef_string_from_utf8(cachePath, C.strlen(cachePath),
            &cefSettings.cache_path)

    // log_severity
    cefSettings.log_severity =
            (C.cef_log_severity_t)(C.int(settings.LogSeverity))

    C.cef_initialize(&mainArgs, &cefSettings, &app, sandbox)
}
func Shutdown() {
    C.cef_shutdown()
}
