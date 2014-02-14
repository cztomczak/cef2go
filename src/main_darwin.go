// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package main

import (
    "cef"
    "cocoa"
    "os"
    "log"
)

var Logger *log.Logger = log.New(os.Stdout, "[main] ", log.Lshortfile)

func main() {
    // Executable's directory
    exeDir := cocoa.GetExecutableDir()

    // CEF subprocesses.
    cef.ExecuteProcess(nil)

    // Initialize CEF.
    settings := cef.Settings{}
    settings.CachePath = exeDir + "/webcache" // Set to empty to disable
    settings.LogSeverity = cef.LOGSEVERITY_DEFAULT // LOGSEVERITY_VERBOSE
    settings.LogFile = exeDir + "/debug.log"
    //settings.LocalesDirPath = cwd + "/cef.framework/Resources"
    //settings.ResourcesDirPath = cwd + "/cef.framework/Resources"
    cef.Initialize(settings)

    // Create Window using Cocoa API.
    cocoa.InitializeApp()
    window := cocoa.CreateWindow("cef2go example", 1024, 768)
    cocoa.ConnectDestroySignal(window, OnDestroyWindow)
    cocoa.ActivateApp()

    // Create browser.
    browserSettings := cef.BrowserSettings{}
    url := "file://" + exeDir + "/example.html"
    cef.CreateBrowser(window, browserSettings, url)

    // CEF loop and shutdown.
    cef.RunMessageLoop()
    cef.Shutdown()
    os.Exit(0)
}

func OnDestroyWindow() {
    cef.QuitMessageLoop()
}
