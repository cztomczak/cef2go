cef2go - HTML 5 based GUI toolkit for the Go language
=====================================================

cef2go is an open source project founded by [Czarek Tomczak]
(http://www.linkedin.com/in/czarektomczak) in 2014
to provide Go bindings for the [Chromium Embedded Framework]
(https://code.google.com/p/chromiumembedded/) (CEF).
cef2go can act as a GUI toolkit, allowing you to create an HTML 5
based GUI in your application. Or you can just provide browser
capabilities to your application.

Supported platforms: Windows (Linux should appear soon).
For OS X see the [go-cef](https://github.com/adieu/go-cef) project.

Currently the cef2go example creates just a simple window with
the Chromium browser embedded. More advanced bindings are in
plans, and that includes javascript bindings and callbacks, so
that you can have bidirectional communication between Go and
Javascript.

cef2go is licensed under the BSD 3-clause license, see the LICENSE
file.


Binary demo
-----------
See the cef2go binary demo for Windows, works out of the box:
[](...)


Getting started on Windows
--------------------------
1. Install mingw and add C:\MinGW\bin to PATH. You can install mingw
   using [mingw-get-setup.exe]
   (http://sourceforge.net/projects/mingw/files/Installer/).
   Select packages to install: "mingw-developer-toolkit",
   "mingw32-base", "msys-base". cef2go was tested and works fine
   with GCC 4.8.2. You can check gcc version with "gcc --version".

2. Download CEF 3 branch 1750 revision 1590 binaries:
   [cef_binary_3.1750.1590_windows32.7z]
   (cef_binary_3.1750.1590_windows32.7z)

   Copy Release/* to cef2go/Release  
   Copy Resources/* to cef2go/Release  

3. Run build_win.bat

