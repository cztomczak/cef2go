// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package cef

/*
#cgo CFLAGS: -I./../../ -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include <stdlib.h>
#include <string.h>
#include "include/capi/cef_app_capi.h"
#include <Cocoa/Cocoa.h>
NSRect GetWindowBounds(void* view) {
    [NSAutoreleasePool new];
    NSView* nsView = (__bridge NSView*)view;
    return [nsView bounds];
}
*/
import "C"
import "unsafe"

import (
    "os"
)

var _Argv []*C.char = make([]*C.char, len(os.Args))

func FillMainArgs(mainArgs *C.struct__cef_main_args_t,
        appHandle unsafe.Pointer) {
    // On Mac appHandle is nil.
    Logger.Println("FillMainArgs, argc=", len(os.Args))
    for i, arg := range os.Args {
        _Argv[C.int(i)] = C.CString(arg)
    }
    mainArgs.argc = C.int(len(os.Args))
    mainArgs.argv = &_Argv[0]
}

func FillWindowInfo(windowInfo *C.cef_window_info_t, hwnd unsafe.Pointer) {
    Logger.Println("FillWindowInfo")
    
    // Setting title isn't required for the CEF inner window.
    // --
    // var windowName *C.char = C.CString("TODO Darwin example")
    // defer C.free(unsafe.Pointer(windowName))
    // C.cef_string_from_utf8(windowName, C.strlen(windowName),
    //        &windowInfo.window_name)

    var bounds C.NSRect = C.GetWindowBounds(hwnd)

    windowInfo.x = C.int(bounds.origin.x)
    windowInfo.y = C.int(bounds.origin.y)
    windowInfo.width = C.int(bounds.size.width)
    windowInfo.height = C.int(bounds.size.height)

    // parent
    windowInfo.parent_view = hwnd
}
