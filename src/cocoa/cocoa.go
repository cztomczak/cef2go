// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package cocoa

/*
#cgo CFLAGS: -I./../../ -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include <stdlib.h>
#include <string.h>
#include <Cocoa/Cocoa.h>
#include <mach-o/dyld.h>
void InitializeApp() {
    [NSAutoreleasePool new];
    [NSApplication sharedApplication];
    [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
}
@interface WindowDelegate : NSObject <NSWindowDelegate> {
    @private
        NSView* view_;
}
@property (nonatomic, assign) NSView* view;
@end
@implementation WindowDelegate
@synthesize view = view_;
- (void)windowWillClose:(NSNotification *)notification {
    [NSAutoreleasePool new];
    printf("NSWindowDelegate::windowWillClose\n");
    _GoDestroySignal((__bridge void*)view_);
}
@end
void* CreateWindow(char* title, int width, int height) {
    [NSAutoreleasePool new];
    WindowDelegate* delegate = [[WindowDelegate alloc] init];
    id window = [[NSWindow alloc]
              initWithContentRect:NSMakeRect(0, 0, width, height)
              styleMask:(NSTitledWindowMask |
                         NSClosableWindowMask |
                         NSMiniaturizableWindowMask |
                         NSResizableWindowMask |
                         NSUnifiedTitleAndToolbarWindowMask )
              backing:NSBackingStoreBuffered
              defer:NO];
    delegate.view = [window contentView];
    [window setDelegate:delegate];
    [window cascadeTopLeftFromPoint:NSMakePoint(20,20)];
    [window setTitle:[NSString stringWithUTF8String:title]];
    [window makeKeyAndOrderFront:nil];
    return (__bridge void*) [window contentView];
}
void ActivateApp() {
    [NSAutoreleasePool new];
    [NSApp activateIgnoringOtherApps:YES];
}
*/
import "C"
import "unsafe"
import (
    "os"
    "log"
    "path/filepath"
)

var Logger *log.Logger = log.New(os.Stdout, "[cocoa] ", log.Lshortfile)

func InitializeApp() {
    C.InitializeApp()
}

func CreateWindow(title string, width int, height int) unsafe.Pointer {
    Logger.Println("CreateWindow")
    csTitle := C.CString(title)
    defer C.free(unsafe.Pointer(csTitle))
    window := C.CreateWindow(csTitle, C.int(width), C.int(height))
    return window
}

func ActivateApp() {
    C.ActivateApp()
}

type DestroyCallback func()
var destroySignalCallbacks map[uintptr]DestroyCallback =
        make(map[uintptr]DestroyCallback)

func ConnectDestroySignal(window unsafe.Pointer, callback DestroyCallback) {
    Logger.Println("ConnectDestroySignal")
    ptr := uintptr(window)
    destroySignalCallbacks[ptr] = callback
}

func GetExecutableDir() string {
    var path []C.char = make([]C.char, 1024)
    var size C.uint32_t = 1024
    if (C._NSGetExecutablePath(&path[0], &size) == 0) {
        return filepath.Dir(C.GoString(&path[0]))
    }
    return "."
}
