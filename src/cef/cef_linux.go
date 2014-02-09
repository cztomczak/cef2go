// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package cef

/*
#cgo CFLAGS: -I./../../
#cgo LDFLAGS: -L./../../Release -lcef
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
*/
import "C"
import "unsafe"
import (
    "os"
)

func FillMainArgs(mainArgs *C.struct__cef_main_args_t,
        appHandle unsafe.Pointer) {
    // On Linux appHandle is nil.
    // Converting os.Args to C equivalent argc/argv.
    var char *C.char
    charSize := int(unsafe.Sizeof(char))
    argv := C.malloc(C.size_t(charSize * len(os.Args)))
    for i, arg := range os.Args {
        // These C strings are not supposed to be freed,
        // do not call C.free().
        charArg := C.CString(arg)
        ptr := unsafe.Pointer(uintptr(argv) + uintptr(charSize * i))
        ptrSize := C.size_t(unsafe.Sizeof(ptr))
        C.memcpy(ptr, unsafe.Pointer(charArg), ptrSize)
    }
    mainArgs.argc = C.int(len(os.Args))
    mainArgs.argv = (**C.char)(argv)
}
