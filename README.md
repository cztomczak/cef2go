CEF2go - HTML 5 based GUI toolkit for the Go language
=====================================================

CEF2go is an open source project founded by [Czarek Tomczak]
(http://www.linkedin.com/in/czarektomczak) in 2014
to provide Go bindings for the [Chromium Embedded Framework]
(https://code.google.com/p/chromiumembedded/) (CEF).
CEF2go can act as a GUI toolkit, allowing you to create an HTML 5
based GUI in your application. Or you can provide browser
capabilities to your application.

Supported platforms: Windows, Linux (OS X not yet ready).

Currently the CEF2go example creates just a simple window with
the Chromium browser embedded. You can set a few options for
your application like the cache directory. More advanced bindings
are in plans, and that includes javascript bindings and callbacks, so
that you can have bidirectional communication between Go and
Javascript.

CEF2go is licensed under the BSD 3-clause license, see the LICENSE
file.

Help
----
Ask questions on the [CEF2go Forum](http://groups.google.com/group/cef2go).


Binary examples
---------------
Windows example: [releases/tag/v0.10]
(https://github.com/CzarekTomczak/cef2go/releases/tag/v0.10)  

Linux example: [releases/tag/v0.11]
(https://github.com/CzarekTomczak/cef2go/releases/tag/v0.11)  


Support development
-------------------

Both code contributions and Paypal donations are welcome.
[![Donate through Paypal]
(https://raw2.github.com/CzarekTomczak/cef2go/master/donate.gif)]
(https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=9CAMRSA48UVT8)


Getting started on Windows
--------------------------
1. Install mingw and add C:\MinGW\bin to PATH. You can install mingw
   using [mingw-get-setup.exe]
   (http://sourceforge.net/projects/mingw/files/Installer/).
   Select packages to install: "mingw-developer-toolkit",
   "mingw32-base", "msys-base". CEF2go was tested and works fine
   with GCC 4.8.2. You can check gcc version with "gcc --version".

2. Download CEF 3 branch 1750 revision 1590 binaries:
   [cef_binary_3.1750.1590_windows32.7z]
   (https://github.com/CzarekTomczak/cef2go/releases/download/cef3-b1750-r1590/cef_binary_3.1750.1590_windows32.7z)  
   Copy Release/* to cef2go/Release  
   Copy Resources/* to cef2go/Release  

3. Run build.bat


Getting started on Linux
------------------------
1. These instructions work fine with Ubuntu 12.04 64-bit. 
   May also work with other versions, but were not tested.

2. Install CEF dependencies:  
   `sudo apt-get install build-essential libgtk2.0-dev libgtkglext1-dev`

3. Download CEF 3 branch 1750 revision 1604 binaries:
   [cef_binary_notcmalloc_3.1750.1604_linux64.zip]
   (https://github.com/CzarekTomczak/cef2go/releases/download/cef3-b1750-r1604/cef_binary_notcmalloc_3.1750.1604_linux64.zip)  
   Copy Release/* to cef2go/Release

4. Run "make" command.
