package main

import "./wingui"
import "os"

func main() {
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

    os.Exit(int(m.Wparam))
}
