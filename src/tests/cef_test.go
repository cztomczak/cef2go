// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

// cef_test package tests the cef package. The tests
// had to be put to a separate package to speed up
// the running of tests. When the test was inside the
// cef package, then the cef package was recompiled
// each time, even though it was installed in GOPATH
// before running the test. And this took time and
// and slowed down the build process significantly.
package cef_test

import (
    "testing"
    "cef"
    "log"
    "os"
)

var g_logger *log.Logger = log.New(os.Stdout, "[cef_test] ", log.Lshortfile)

func Test_WorkingDirectory(t *testing.T) {
    // Change working directory while running tests, otherwise
    // CEF may have troubles finding the resource pak files.
    // os.Chdir("./../../Release")
}

func Test_ExecuteProcess(t *testing.T) {
    g_logger.Println("Test_ExecuteProcess")
    // If called for the browser process it will return 
    // immediately with a value of -1
    code := cef.ExecuteProcess(nil)
    g_logger.Println("ExecuteProcess returned:", code)
}

func Test_Initialize(t *testing.T) {
    g_logger.Println("Test_Initialize")
    settings := cef.Settings{}
    init := cef.Initialize(settings)
    g_logger.Println("Initialize() returned:", init)
    if init != 1 {
        t.Errorf("Initialize() returned: %d", init)
    }
}

func Test_Shutdown(t *testing.T) {
    g_logger.Println("Test_Shutdown")
    cef.Shutdown()
}
