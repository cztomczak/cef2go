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

// WinProc called by windows to notify us of all windows events we might be interested in.
func WndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (rc uintptr) {
    switch msg {
    case WM_CREATE:
        rc = DefWindowProc(hwnd, msg, wparam, lparam)
    /*
    case WM_COMMAND:
        switch syscall.Handle(lparam) {
        case bh:
            e := PostMessage(hwnd, WM_CLOSE, 0, 0)
            if e != nil {
                AbortErrNo("PostMessage", e)
            }
        default:
            rc = DefWindowProc(hwnd, msg, wparam, lparam)
        }
    */
    case WM_CLOSE:
        DestroyWindow(hwnd)
    case WM_DESTROY:
        PostQuitMessage(0)
    default:
        rc = DefWindowProc(hwnd, msg, wparam, lparam)
    }
    //fmt.Printf("WndProc(0x%08x, %d, 0x%08x, 0x%08x) (%d)\n", hwnd, msg, wparam, lparam, rc)
    return
}

func CreateWindow(title string) (hwnd syscall.Handle) {
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

    // Create callback
    wproc := syscall.NewCallback(WndProc)

    // RegisterClassEx
    wcname := syscall.StringToUTF16Ptr("myWindowClass")
    var wc Wndclassex
    wc.Size = uint32(unsafe.Sizeof(wc))
    wc.WndProc = wproc
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
