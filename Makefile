.PHONY: detect_os Linux Darwin clean

export GOPATH=$(PWD)
UNAME_S = $(shell uname -s)

ifeq ($(UNAME_S), Linux)
	INC=-I. \
		-I/usr/include/gtk-2.0 \
		-I/usr/include/glib-2.0 \
		-I/usr/include/cairo \
		-I/usr/include/pango-1.0 \
		-I/usr/include/gdk-pixbuf-2.0 \
		-I/usr/include/atk-1.0 \
		-I/usr/lib/x86_64-linux-gnu/glib-2.0/include \
		-I/usr/lib/x86_64-linux-gnu/gtk-2.0/include \
		-I/usr/lib/i386-linux-gnu/gtk-2.0/include \
		-I/usr/lib/i386-linux-gnu/glib-2.0/include \
		-I/usr/lib64/glib-2.0/include \
		-I/usr/lib64/gtk-2.0/include
	export CC=gcc $(INC)
	export CGO_LDFLAGS=-L $(PWD)/Release -lcef
else ifeq ($(UNAME_S), Darwin)
	INC=-I.
	export CGO_ENABLED=1
	export CC=clang $(INC)
	export CGO_LDFLAGS=-F$(PWD)/Release/tmp -framework Cocoa -framework cef
endif

detect_os:
	make $(UNAME_S)

Linux:
	clear
	go install gtk
	go install cef
	go test -ldflags "-r $(PWD)/Release" src/tests/cef_test.go
	go build -ldflags "-r ." -o Release/cef2go src/main_linux.go
	cd Release && ./cef2go && cd ../

Darwin:
	clear
	
	@# Required for linking. Go doesn't allow framework name
	@# to contain spaces, so we're making a copy of the framework
	@# without spaces.
	@if [ ! -d Release/tmp ]; then \
		echo Copying CEF framework directory to Release/tmp ;\
		mkdir -p Release/tmp ;\
		cp -rf Release/cef2go.app/Contents/Frameworks/Chromium\ Embedded\ Framework.framework Release/tmp/cef.framework ;\
		mv Release/tmp/cef.framework/Chromium\ Embedded\ Framework Release/tmp/cef.framework/cef ;\
	fi
	go install -x cef
	@# CEF requires specific app bundle / directory structure
	@# on OSX, but Go doesn't allow for such thing when 
	@# running test. So turning off test.
	@# go test -ldflags "-r $(PWD)/Release" src/tests/cef_test.go
	rm -f Release/cef2go.app/Contents/MacOS/cef2go
	go build -x -ldflags "-r ." -o Release/cef2go.app/Contents/MacOS/cef2go src/main_darwin.go
	install_name_tool -change @executable_path/Chromium\ Embedded\ Framework @executable_path/../Frameworks/Chromium\ Embedded\ Framework.framework/Chromium\ Embedded\ Framework Release/cef2go.app/Contents/MacOS/cef2go
	cp -f Release/example.html Release/cef2go.app/Contents/MacOS/example.html
	cd Release/cef2go.app/Contents/MacOS && ./cef2go && cd ../../../../

clean:
	clear
	go clean -i cef
