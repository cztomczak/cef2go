// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

// DESCRIPTION:
// ----------------------------------------------------------------------------
// This example shows how to run an internal web server while
// embedding Chromium browser. You can communicate with Go from
// javascript using XMLHttpRequests.
// ----------------------------------------------------------------------------

package main

import (
    "fmt"
    "net/http"
    "runtime"
)

// Imports from "main_windows.go"
import (
    "cef"
    "wingui"
    "os"
    "syscall"
    "unsafe"
    "log"
    "time"
)

func handler(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, "This is Go server talking.<br>")
    fmt.Fprintf(w, "Time on the server: %v<br>",
            time.Now().Format("Jan 2, 2006 at 3:04pm (MST)"))
    fmt.Fprintf(w, "Go version: %v<br>", runtime.Version())
    fmt.Fprintf(w, "<br>")
    if req.URL.Path == "/" {
        fmt.Fprintf(w, "Try <a href=/test.go>/test.go</a><br>")
    } else if req.URL.Path == "/test.go" {
        fmt.Fprintf(w, "<b>You accessed /test.go</b><br>")
    }
}

func run_http_server() {
    http.HandleFunc("/", handler)
    listen_at := "127.0.0.1:54007"
    fmt.Printf("Running http server at %s\n", listen_at)
    http.ListenAndServe(listen_at, nil)
}

// ----------------------------------------------------------------------------
// The code below copied from "main_windows.go" with the following changes:
// 1. Added a call to run_http_server() at the beginning of main.
// 2. Changed url at browser creation to "http://127.0.0.1:54007/"
// 3. Imports were moved to the top of the file
// ----------------------------------------------------------------------------

var Logger *log.Logger = log.New(os.Stdout, "[main] ", log.Lshortfile)

func main() {
    go run_http_server()

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
    url = "http://127.0.0.1:54007/"
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
