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
)

func Test_Initialize(t *testing.T) {
    settings := cef.Settings{}
    init := cef.Initialize(settings)
    if init != 1 {
        t.Error("Initialize() returned: %d", init)
    }
}
func Test_Shutdown(t *testing.T) {
    cef.Shutdown()
}
