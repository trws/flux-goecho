# CGO_CFLAGS:=$(shell pkg-config --cflags)
#-I/g/g12/scogland/projects/flux/flux-core/src/include/ -I/g/g12/scogland/projects/flux/flux-core -I/usr/include/json -I/usr/include/czmq
# CGO_LDFLAGS:=$(shell pkg-config --libs)
#-L/g/g12/scogland/projects/flux/build/src/common/.libs -lflux-core

all: gotest.so

gotest.so gotest.h: test.go name.c Makefile
	rm -f gotest.h #clear dead header if it exists
	go build -buildmode c-shared -ldflags '-extldflags "-z nodelete"' -o gotest.so .

clean:
	rm -f *.so *.h

test: test_src/testdlsym.c all
	gcc test_src/testdlsym.c -o test -ldl -g -std=c99 -pedantic
	LD_LIBRARY_PATH=~/projects/flux/flux-core/src/common/.libs:./ ./test
