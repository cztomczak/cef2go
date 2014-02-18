// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cefcapi

#pragma once

#include "handlers/cef_base.h"
#include "include/capi/cef_app_capi.h"

// ----------------------------------------------------------------------------
// cef_app_t
// ----------------------------------------------------------------------------

///
// Implement this structure to provide handler implementations. Methods will be
// called by the process and/or thread indicated.
///

///
// Provides an opportunity to view and/or modify command-line arguments before
// processing by CEF and Chromium. The |process_type| value will be NULL for
// the browser process. Do not keep a reference to the cef_command_line_t
// object passed to this function. The CefSettings.command_line_args_disabled
// value can be used to start with an NULL command-line object. Any values
// specified in CefSettings that equate to command-line arguments will be set
// before this function is called. Be cautious when using this function to
// modify command-line arguments for non-browser processes as this may result
// in undefined behavior including crashes.
///
void CEF_CALLBACK on_before_command_line_processing(
        struct _cef_app_t* self, const cef_string_t* process_type,
        struct _cef_command_line_t* command_line) {
    DEBUG_CALLBACK("on_before_command_line_processing\n");
}

///
// Provides an opportunity to register custom schemes. Do not keep a reference
// to the |registrar| object. This function is called on the main thread for
// each process and the registered schemes should be the same across all
// processes.
///
void CEF_CALLBACK on_register_custom_schemes(
        struct _cef_app_t* self,
        struct _cef_scheme_registrar_t* registrar) {
    DEBUG_CALLBACK("on_register_custom_schemes\n");
}

///
// Return the handler for resource bundle events. If
// CefSettings.pack_loading_disabled is true (1) a handler must be returned.
// If no handler is returned resources will be loaded from pack files. This
// function is called by the browser and render processes on multiple threads.
///
struct _cef_resource_bundle_handler_t*
        CEF_CALLBACK get_resource_bundle_handler(struct _cef_app_t* self) {
    DEBUG_CALLBACK("get_resource_bundle_handler\n");
    return NULL;
}

///
// Return the handler for functionality specific to the browser process. This
// function is called on multiple threads in the browser process.
///
struct _cef_browser_process_handler_t* 
        CEF_CALLBACK get_browser_process_handler(struct _cef_app_t* self) {
    DEBUG_CALLBACK("get_browser_process_handler\n");
    return NULL;
}

///
// Return the handler for functionality specific to the render process. This
// function is called on the render process main thread.
///
struct _cef_render_process_handler_t*
        CEF_CALLBACK get_render_process_handler(struct _cef_app_t* self) {
    DEBUG_CALLBACK("get_render_process_handler\n");
    return NULL;
}

void initialize_app_handler(cef_app_t* app) {
    printf("initialize_app_handler\n");
    app->base.size = sizeof(cef_app_t);
    initialize_cef_base((cef_base_t*)app);
    // callbacks
    app->on_before_command_line_processing = on_before_command_line_processing;
    app->on_register_custom_schemes = on_register_custom_schemes;
    app->get_resource_bundle_handler = get_resource_bundle_handler;
    app->get_browser_process_handler = get_browser_process_handler;
    app->get_render_process_handler = get_render_process_handler;
}
