// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cefcapi

#pragma once

#include <stdio.h>
#include "include/capi/cef_base_capi.h"

// Set to 1 to check if add_ref() and release()
// are called and to track the total number of calls.
// add_ref will be printed as "+", release as "-".
#define DEBUG_REFERENCE_COUNTING 0

// Print only the first execution of the callback,
// ignore the subsequent.
#define DEBUG_CALLBACK(x) { \
    static int first_call = 1; \
    if (first_call == 1) { \
        first_call = 0; \
        printf(x); \
    } \
}

// ----------------------------------------------------------------------------
// cef_base_ref_counted_t
// ----------------------------------------------------------------------------

///
// Structure defining the reference count implementation functions.
// All framework structures must include the cef_base_ref_counted_t 
// structure first.
///

///
// Increment the reference count.
///
void CEF_CALLBACK add_ref(cef_base_ref_counted_t* self) {
    DEBUG_CALLBACK("cef_base_ref_counted_t.add_ref\n");
    if (DEBUG_REFERENCE_COUNTING)
        printf("+");
}

///
// Decrement the reference count.  Delete this object when no references
// remain.
///
int CEF_CALLBACK release(cef_base_ref_counted_t* self) {
    DEBUG_CALLBACK("cef_base_ref_counted_t.release\n");
    if (DEBUG_REFERENCE_COUNTING)
        printf("-");
    return 1;
}

///
// Returns the current number of references.
///
int CEF_CALLBACK has_one_ref(cef_base_ref_counted_t* self) {
    DEBUG_CALLBACK("cef_base_ref_counted_t.has_one_ref\n");
    if (DEBUG_REFERENCE_COUNTING)
        printf("=");
    return 1;
}

void initialize_cef_base_ref_counted(cef_base_ref_counted_t* base) {
    printf("initialize_cef_base_ref_counted\n");
    // Check if "size" member was set.
    size_t size = base->size;
    // Let's print the size in case sizeof was used
    // on a pointer instead of a structure. In such
    // case the number will be very high.
    printf("cef_base_ref_counted_t.size = %lu\n", (unsigned long)size);
    if (size <= 0) {
        printf("FATAL: initialize_cef_base failed, size member not set\n");
        _exit(1);
    }
    base->add_ref = add_ref;
    base->release = release;
    base->has_one_ref = has_one_ref;
}
