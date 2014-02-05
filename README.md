HTML 5 based GUI toolkit for the Go language
=====================================================

cef2go is an open source project founded by [Czarek Tomczak]
(http://www.linkedin.com/in/czarektomczak) in 2014
to provide Go bindings for the [Chromium Embedded Framework]
(https://code.google.com/p/chromiumembedded/) (CEF).
cef2go can act as a GUI toolkit, allowing you to create an HTML 5
based GUI in your application. Or you can just provide browser
capabilities to your application.

Supported platforms: Windows (Linux should appear soon).

Currently the cef2go example creates just a simple window with
the Chromium browser embedded. More advanced bindings are in
plans, and that includes javascript bindings and callbacks, so
that you can have bidirectional communication between Go and
Javascript.

cef2go is licensed under the New BSD License (BSD 3-clause),
see the LICENSE file.

See also CEF bindings for Mac OS X:
https://github.com/adieu/go-cef
