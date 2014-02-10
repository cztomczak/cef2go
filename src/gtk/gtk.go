// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package gtk

/*
#cgo pkg-config: --libs --cflags gtk+-2.0
#include <gtk/gtk.h>
#include <stdlib.h>
#include <string.h>
static inline GtkWindow* ToGtkWindow(GtkWidget* w) { return GTK_WINDOW(w); }
static inline GtkContainer* ToGtkContainer(GtkWidget* w) { return GTK_CONTAINER(w); }
void TerminationSignal(int signatl) { cef_quit_message_loop(); }
void ConnectTerminationSignal() {
    signal(SIGINT, TerminationSignal);
    signal(SIGTERM, TerminationSignal);
}
void DestroySignal(GtkWidget* widget, gpointer data) { 
    _GoDestroySignal(widget, data);
}
void ConnectDestroySignal(GtkWidget* window) {
    g_signal_connect(G_OBJECT(window), "destroy",
            G_CALLBACK(DestroySignal), NULL);
}
*/
import "C"
import "unsafe"
import (
    "os"
    "log"
)

var Logger *log.Logger = log.New(os.Stdout, "[gtk] ", log.Lshortfile)

func Initialize() {
    C.gtk_init(nil, nil)
    C.ConnectTerminationSignal()
}

func CreateWindow(title string, width int, height int) unsafe.Pointer {
    Logger.Println("CreateWindow")
    
    // Create window.
    window := C.gtk_window_new(C.GTK_WINDOW_TOPLEVEL)
    
    // Default size.
    C.gtk_window_set_default_size(C.ToGtkWindow(window),
            C.gint(width), C.gint(height))
    
    // Center.
    C.gtk_window_set_position(C.ToGtkWindow(window), C.GTK_WIN_POS_CENTER)
    
    // Title.
    csTitle := C.CString(title)
    defer C.free(unsafe.Pointer(csTitle))
    C.gtk_window_set_title(C.ToGtkWindow(window), (*C.gchar)(csTitle))
    
    // TODO: focus
    // g_signal_connect(window, "focus", G_CALLBACK(&HandleFocus), NULL);

    // CEF requires a container. Embedding browser in a top
    // level window fails.
    vbox := C.gtk_vbox_new(0, 0)
    C.gtk_container_add(C.ToGtkContainer(window), vbox)
    
    // Show.
    C.gtk_widget_show_all(window)

    return unsafe.Pointer(vbox)
}

type DestroyCallback func()
var destroySignalCallbacks map[uintptr]DestroyCallback =
        make(map[uintptr]DestroyCallback)

func ConnectDestroySignal(window unsafe.Pointer, callback DestroyCallback) {
    Logger.Println("ConnectDestroySignal")
    ptr := uintptr(window)
    destroySignalCallbacks[ptr] = callback
    C.ConnectDestroySignal((*C.GtkWidget)(window))
}

