// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package main

import (
    "cef"
    "wingui"
    "os"
    "syscall"
    "unsafe"
    "log"
    "time"
)

var Logger *log.Logger = log.New(os.Stdout, "[main] ", log.Lshortfile)

func main() {
    hInstance, e := wingui.GetModuleHandle(nil)
    if e != nil { wingui.AbortErrNo("GetModuleHandle", e) }
    
    cef.ExecuteProcess(unsafe.Pointer(hInstance))
    
    settings := cef.Settings{}
    settings.CachePath = "webcache" // Set to empty to disable
    settings.LogSeverity = cef.LOGSEVERITY_DEFAULT // LOGSEVERITY_VERBOSE
    cef.Initialize(settings)
    
    wndproc := syscall.NewCallback(WndProc)
    Logger.Println("CreateWindow")
    hwnd := wingui.CreateWindow("cef2go example", wndproc)

    browserSettings := cef.BrowserSettings{}
    // TODO: It should be executable's directory used
    // rather than working directory.
    url, _ := os.Getwd()
    url = "file://" + url + "/example.html"
    cef.CreateBrowser(unsafe.Pointer(hwnd), browserSettings, url)

    // It should be enough to call WindowResized after 10ms,
    // though to be sure let's extend it to 100ms.
    time.AfterFunc(time.Millisecond * 100, func(){
        cef.WindowResized(unsafe.Pointer(hwnd))
    })

    cef.RunMessageLoop()
    cef.Shutdown()
    os.Exit(0)
}

func WndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (rc uintptr) {
    switch msg {
    case wingui.WM_CREATE:
        rc = wingui.DefWindowProc(hwnd, msg, wparam, lparam)
    case wingui.WM_SIZE:
        cef.WindowResized(unsafe.Pointer(hwnd))
    case wingui.WM_CLOSE:
        wingui.DestroyWindow(hwnd)
    case wingui.WM_DESTROY:
        cef.QuitMessageLoop()
    default:
        rc = wingui.DefWindowProc(hwnd, msg, wparam, lparam)
    }
    return
}
