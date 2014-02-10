// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package main

import (
    "cef"
    "gtk"
    "os"
    "log"
)

var Logger *log.Logger = log.New(os.Stdout, "[main] ", log.Lshortfile)

func main() {
    // TODO: It should be executable's directory used
    // rather than working directory.
    cwd, _ := os.Getwd()

    // CEF subprocesses.
    cef.ExecuteProcess(nil)

    // CEF initialize.
    settings := cef.Settings{}
    settings.CachePath = cwd + "/webcache" // Set to empty to disable
    settings.LogSeverity = cef.LOGSEVERITY_DEFAULT // LOGSEVERITY_VERBOSE
    settings.LogFile = cwd + "/debug.log"
    cef.Initialize(settings)

    // Create GTK window.
    gtk.Initialize()
    window := gtk.CreateWindow("cef2go example", 1024, 768)
    gtk.ConnectDestroySignal(window, OnDestroyWindow)

    // Create browser.
    browserSettings := cef.BrowserSettings{}
    url := "file://" + cwd + "/example.html"
    cef.CreateBrowser(window, browserSettings, url)

    // CEF loop and shutdown.
    cef.RunMessageLoop()
    cef.Shutdown()
    os.Exit(0)
}

func OnDestroyWindow() {
    cef.QuitMessageLoop()
}
