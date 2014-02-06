// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package wingui

import (
    "fmt"
    "os"
    "syscall"
    "unsafe"
)

// some help functions

func abortf(format string, a ...interface{}) {
    fmt.Fprintf(os.Stdout, format, a...)
    os.Exit(1)
}

func AbortErrNo(funcname string, err error) {
    errno, _ := err.(syscall.Errno)
    abortf("%s failed: %d %s\n", funcname, uint32(errno), err)
}

// global vars

var (
    mh syscall.Handle
    bh syscall.Handle
)

func CreateWindow(title string, wndproc uintptr) (hwnd syscall.Handle) {
    var e error

    // GetModuleHandle
    mh, e = GetModuleHandle(nil)
    if e != nil {
        AbortErrNo("GetModuleHandle", e)
    }

    // Get icon we're going to use.
    myicon, e := LoadIcon(0, IDI_APPLICATION)
    if e != nil {
        AbortErrNo("LoadIcon", e)
    }

    // Get cursor we're going to use.
    mycursor, e := LoadCursor(0, IDC_ARROW)
    if e != nil {
        AbortErrNo("LoadCursor", e)
    }

    // RegisterClassEx
    wcname := syscall.StringToUTF16Ptr("myWindowClass")
    var wc Wndclassex
    wc.Size = uint32(unsafe.Sizeof(wc))
    wc.WndProc = wndproc
    wc.Instance = mh
    wc.Icon = myicon
    wc.Cursor = mycursor
    wc.Background = COLOR_BTNFACE + 1
    wc.MenuName = nil
    wc.ClassName = wcname
    wc.IconSm = myicon
    if _, e := RegisterClassEx(&wc); e != nil {
        AbortErrNo("RegisterClassEx", e)
    }

    // CreateWindowEx
    wh, e := CreateWindowEx(
        0,
        wcname,
        syscall.StringToUTF16Ptr(title),
        WS_OVERLAPPEDWINDOW,
        CW_USEDEFAULT, CW_USEDEFAULT, CW_USEDEFAULT, CW_USEDEFAULT,
        0, 0, mh, 0)
    if e != nil {
        AbortErrNo("CreateWindowEx", e)
    }
    //fmt.Printf("main window handle is %x\n", wh)

    // ShowWindow
    ShowWindow(wh, SW_SHOWDEFAULT)

    // UpdateWindow
    if e := UpdateWindow(wh); e != nil {
        AbortErrNo("UpdateWindow", e)
    }

    hwnd = wh

    return

    /*
    // Process all windows messages until WM_QUIT.
    var m Msg
    for {
        r, e := GetMessage(&m, 0, 0, 0)
        if e != nil {
            AbortErrNo("GetMessage", e)
        }
        if r == 0 {
            // WM_QUIT received -> get out
            break
        }
        TranslateMessage(&m)
        DispatchMessage(&m)
    }
    return int(m.Wparam)
    */
}


/*
func main() {
    rc := rungui()
    os.Exit(rc)
}
*/
