// CEF C API example
// Project website: https://github.com/cztomczak/cefcapi

#pragma once

#include "handlers/cef_base.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"

// ----------------------------------------------------------------------------
// struct cef_life_span_handler_t
// ----------------------------------------------------------------------------

///
// Implement this structure to handle events related to browser life span. The
// functions of this structure will be called on the UI thread unless otherwise
// indicated.
///

// NOTE: There are many more callbacks in cef_life_span_handler,
//       but only on_before_close is implemented here.

///
// Called just before a browser is destroyed. Release all references to the
// browser object and do not attempt to execute any functions on the browser
// object after this callback returns. This callback will be the last
// notification that references |browser|. See do_close() documentation for
// additional usage information.
///
void CEF_CALLBACK on_before_close(struct _cef_life_span_handler_t* self,
                                  struct _cef_browser_t* browser) {
    DEBUG_CALLBACK("on_before_close\n");
    // TODO: Check how many browsers do exist and quit message
    //       loop only when last browser is closed. Otherwise
    //       closing a popup window will exit app while main
    //       window shouldn't be closed.
    cef_quit_message_loop();
}

void initialize_cef_life_span_handler(cef_life_span_handler_t* handler) {
    DEBUG_CALLBACK("initialize_cef_life_span_handler\n");
    handler->base.size = sizeof(cef_life_span_handler_t);
    initialize_cef_base_ref_counted((cef_base_ref_counted_t*)handler);
    // callbacks - there are many, but implementing only one
    handler->on_before_close = on_before_close;
}
