// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package cocoa

/*
#cgo CFLAGS: -I./../../ -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include <Cocoa/Cocoa.h>
*/
import "C"
import "unsafe"

//export _GoDestroySignal
func _GoDestroySignal(window unsafe.Pointer) {
    Logger.Println("_GoDestroySignal")
    ptr := uintptr(window)
    if callback,ok := destroySignalCallbacks[ptr]; ok {
        delete(destroySignalCallbacks, ptr)
        callback()
    } else {
        Logger.Println("WARNING: _GoDestroySignal failed, callback not found")
    }

}
