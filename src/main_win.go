// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package main

import "./cef"
import "./wingui"
import "os"
import "fmt"
import "syscall"

func main() {
    hInstance, e := wingui.GetModuleHandle(nil)
    if e != nil { wingui.AbortErrNo("GetModuleHandle", e) }
    
    fmt.Printf("cef2go: ExecuteProcess()\n")
    fmt.Printf("cef2go: process args = %v\n", os.Args)
    cef.ExecuteProcess(hInstance)
    
    settings := cef.Settings{}
    settings.CachePath = "webcache" // Set to empty to disable
    settings.LogSeverity = cef.LOGSEVERITY_DEFAULT // LOGSEVERITY_VERBOSE
    fmt.Printf("cef2go: Initialize()\n")
    cef.Initialize(settings)
    
    wndproc := syscall.NewCallback(WndProc)
    fmt.Printf("cef2go: CreateWindow()\n")
    var hwnd syscall.Handle = wingui.CreateWindow("cef2go example", wndproc)

    fmt.Printf("cef2go: CreateBrowser()\n")
    browserSettings := cef.BrowserSettings{}
    // TODO: It should be executable's directory used,
    // and not working directory.
    url, _ := os.Getwd()
    url = "file://" + url + "/example.html"
    fmt.Printf("cef2go: url = %v\n", url)
    cef.CreateBrowser(hwnd, browserSettings, url)
    cef.WindowResized(hwnd)

    fmt.Printf("cef2go: RunMessageLoop()\n")
    cef.RunMessageLoop()

    fmt.Printf("cef2go: Shutdown()\n")
    cef.Shutdown()

    os.Exit(0)
}

func WndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (rc uintptr) {
    switch msg {
    case wingui.WM_CREATE:
        rc = wingui.DefWindowProc(hwnd, msg, wparam, lparam)
    case wingui.WM_SIZE:
        cef.WindowResized(hwnd)
    case wingui.WM_CLOSE:
        wingui.DestroyWindow(hwnd)
    case wingui.WM_DESTROY:
        fmt.Printf("cef2go: QuitMessageLoop()\n")
        cef.QuitMessageLoop()
    default:
        rc = wingui.DefWindowProc(hwnd, msg, wparam, lparam)
    }
    return
}
