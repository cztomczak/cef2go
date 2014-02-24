CEF2go - HTML 5 based GUI toolkit for the Go language
=====================================================

CEF2go is an open source project founded by [Czarek Tomczak]
(http://www.linkedin.com/in/czarektomczak) in 2014
to provide Go bindings for the [Chromium Embedded Framework]
(https://code.google.com/p/chromiumembedded/) (CEF).
CEF2go can act as a GUI toolkit, allowing you to create an HTML 5
based GUI in your application. Or you can provide browser
capabilities to your application.

Supported platforms: Windows, Linux, Mac OSX.

Currently the CEF2go example creates just a simple window with
the Chromium browser embedded. You can set a few options for
your application like the cache directory. More advanced bindings
are in plans, and that includes javascript bindings and callbacks, so
that you can have bidirectional communication between Go and
Javascript in a native way. Though, it is already possible to
communicate with Go from Javascript, see the next section for
an example.

CEF2go is licensed under the BSD 3-clause license, see the LICENSE
file.

Communication between Go and Javascript
---------------------------------------
For now to make communication between Go and javascript possible
you have to run an internal http server and communicate using 
XMLHttpRequests in javascript. See the [http_server_windows.go]
(https://github.com/CzarekTomczak/cef2go/blob/master/src/http_server_windows.go)
example that embeds both a http server and a Chromium browser
in a standalone application. To run it type "build.bat http_server". 
The http server is listening at 127.0.0.1:54007, thus it is not 
accessible from the outside, it can be accessed only from the 
machine it is running on.


Help
----
Ask questions on the [CEF2go Forum](http://groups.google.com/group/cef2go).  
See the auto generated docs for the following packages:
 * [cef](https://godoc.org/github.com/CzarekTomczak/cef2go/src/cef)
 * [cocoa](https://godoc.org/github.com/CzarekTomczak/cef2go/src/cocoa)
 * [gtk](https://godoc.org/github.com/CzarekTomczak/cef2go/src/gtk)
 * [wingui](https://godoc.org/github.com/CzarekTomczak/cef2go/src/wingui)


Binary examples
---------------
The binary examples provided here use branch 1750 (Chrome 33), 
which is the Beta channel of Google Chrome. A stable release of branch
1750 should appear at the beginning of March. Then these examples
should be updated.

Windows example: [releases/tag/v0.10]
(https://github.com/CzarekTomczak/cef2go/releases/tag/v0.10)  

Linux example: [releases/tag/v0.11]
(https://github.com/CzarekTomczak/cef2go/releases/tag/v0.11)  

Mac OSX example: [releases/tag/v0.12]
(https://github.com/CzarekTomczak/cef2go/releases/tag/v0.12)


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

3. Run build.bat (or "build.bat noconsole" to get rid of the console
    window when running the final executable)


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


Getting started on Mac OS X
---------------------------
1. These instructions work fine with OS X 10.8 Mountain Lion.
   May also work with other versions, but were not tested.

2. Install Go 32-bit. Tested with Go 1.2-386 for OSX 10.8.
   CEF binaries for OSX 64-bit are still experimental, that's
   why we're using 32-bit. Though you can try building with
   CEF 64-bit, download binaries from [cefbuilds.com]
   (http://cefbuilds.com).

3. Install command line tools (make is required) from:  
   https://developer.apple.com/downloads/  
   (In my case command line tools for Mountain Lion from September 2013)

4. Install XCode (gcc that comes with XCode is required). 
   Use the link above. In my case it was XCode 4.6.3 from June 2013.

5. Download CEF 3 branch 1750 revision 1625 binaries for 32-bit:
   [releases/tag/v0.12]
   (https://github.com/CzarekTomczak/cef2go/releases/tag/v0.12)  
   Copy the cef2go.app directory to cef2go/Release.

6. Run "make" command.
