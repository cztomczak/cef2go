// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package main

import "./cef"
import "./wingui"
import "os"

func main() {
    settings := cef.Settings{}
    settings.CachePath = "webcache" // Set to empty to disable
    settings.LogSeverity = cef.LOGSEVERITY_DEFAULT // LOGSEVERITY_VERBOSE
    cef.Initialize(settings)
    wingui.CreateWindow("cef2go example")

    // Process all windows messages until WM_QUIT.
    var m wingui.Msg
    for {
        r, e := wingui.GetMessage(&m, 0, 0, 0)
        if e != nil {
            wingui.AbortErrNo("GetMessage", e)
        }
        if r == 0 {
            // WM_QUIT received -> get out
            break
        }
        wingui.TranslateMessage(&m)
        wingui.DispatchMessage(&m)
    }

    cef.Shutdown()
    os.Exit(int(m.Wparam))
}
