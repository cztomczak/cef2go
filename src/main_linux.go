// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package main

import (
    "cef"
    "unsafe"
    "os"
    "log"
)

var Logger *log.Logger = log.New(os.Stdout, "[main] ", log.Lshortfile)

func main() {
    // TODO: It should be executable's directory used,
    // and not working directory.
    cwd, _ := os.Getwd()

    cef.ExecuteProcess(nil)

    settings := cef.Settings{}
    settings.CachePath = cwd + "/webcache" // Set to empty to disable
    settings.LogSeverity = cef.LOGSEVERITY_DEFAULT // LOGSEVERITY_VERBOSE
    settings.LogFile = cwd + "/debug.log"
    cef.Initialize(settings)

    // TODO: create GTK window here. If you pass nil then
    // CEF will create window of its own.
    var hwnd unsafe.Pointer = nil

    browserSettings := cef.BrowserSettings{}
    url := "file://" + cwd + "/example.html"
    cef.CreateBrowser(unsafe.Pointer(hwnd), browserSettings, url)

    cef.RunMessageLoop()
    cef.Shutdown()
    os.Exit(0)
}
