// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package cef

/*
#cgo CFLAGS: -I./../../
#cgo LDFLAGS: -L./../../Release -llibcef
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
*/
import "C"
import "unsafe"
import (
    "syscall"
)

func FillMainArgs(mainArgs* C.struct__cef_main_args_t,
        appHandle syscall.Handle) {
    mainArgs.instance = (C.HINSTANCE)(unsafe.Pointer(appHandle))
}

func FillWindowInfo(windowInfo* C.cef_window_info_t, hwnd syscall.Handle) {
    var rect C.RECT
    C.GetWindowRect((C.HWND)(unsafe.Pointer(hwnd)),
            (*C.struct_tagRECT)(unsafe.Pointer(&rect)))
    windowInfo.style = C.WS_CHILD | C.WS_CLIPCHILDREN | C.WS_CLIPSIBLINGS |
            C.WS_TABSTOP | C.WS_VISIBLE
    windowInfo.parent_window = (C.HWND)(unsafe.Pointer(hwnd))
    windowInfo.x = C.int(rect.left)
    windowInfo.y = C.int(rect.top)
    windowInfo.width = C.int(rect.right - rect.left)
    windowInfo.height = C.int(rect.bottom - rect.top)
}

func WindowResized(hwnd syscall.Handle) {
    var rect C.RECT;
    C.GetClientRect(
            (C.HWND)(unsafe.Pointer(hwnd)),
            (*C.struct_tagRECT)(unsafe.Pointer(&rect)))
    var hdwp C.HDWP = C.BeginDeferWindowPos(1)
    var cefHwnd C.HWND = C.GetWindow(
            (C.HWND)(unsafe.Pointer(hwnd)), C.GW_CHILD)
    hdwp = C.DeferWindowPos(hdwp, cefHwnd, nil,
            C.int(rect.left), C.int(rect.top),
            C.int(rect.right - rect.left),
            C.int(rect.bottom - rect.top),
            C.SWP_NOZORDER)
    C.EndDeferWindowPos(hdwp)
}
