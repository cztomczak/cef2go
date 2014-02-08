.PHONY: all test clean

export GOPATH=$(PWD)
INC = -I/usr/include/gtk-2.0 \
	-I/usr/include/glib-2.0 \
	-I/usr/include/cairo \
	-I/usr/include/pango-1.0 \
	-I/usr/include/gdk-pixbuf-2.0 \
	-I/usr/include/atk-1.0 \
	-I/usr/lib/x86_64-linux-gnu/glib-2.0/include \
	-I/usr/lib/x86_64-linux-gnu/gtk-2.0/include \
	-I/usr/lib/i386-linux-gnu/gtk-2.0/include \
	-I/usr/lib/i386-linux-gnu/glib-2.0/include
export CC=gcc $(INC)

all:
	go install cef
	go test -ldflags "-Wl,-rpath,$(PWD)/Release" src/tests/cef_test.go
	go build -o Release/cef2go src/main_linux.go

clean:
	go clean
