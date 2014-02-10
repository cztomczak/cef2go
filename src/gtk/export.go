// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package gtk

//#include "gtk/gtk.h"
import "C"
import "unsafe"

//export _GoDestroySignal
func _GoDestroySignal(widget *C.GtkWidget, data C.gpointer) {
    Logger.Println("_GoDestroySignal")
    ptr := uintptr(unsafe.Pointer(widget))
    callback := destroySignalCallbacks[ptr]
    delete(destroySignalCallbacks, ptr)
    callback()
}
