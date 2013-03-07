#go-qt5
=====

##Introduction
go-qt5 is provides with qt5 bindings for Go programming language, based on visualfc's go-ui library.

Lua code that generates the wrappers (`uiobjs.go` and `cdrv.cpp`) can be found under `make`.

The wrapper code is by far incomplete, so pull requests are more than welcome. Adding new functionality usually consists of editing or adding files under `make/qt5`, and updating `make/make.lua` script, and making relevant changes in `qt5` and `qtdrv`.


##License
	go-qt5 lib BSD
	qtdrv lib LGPL

##Using go-qt5

###1. get & install go-qt5
	$ go get github.com/salviati/go-qt5
###2. build & install C layer
	$ cd $GOPATH/src/salviati/go-qt5/qtdrv
	$ qmake "CONFIG+=release"
	$ make
	# make install
###3.build go-qt5
	$ cd $GOPATH/src/salviati/go-qt5/qt5
	$ go install
###4.build examples
	$ cd $GOPATH/src/salviati/go-qt5/examples
	$ go run minimal.go

##A minimal example

	package main

	import (
	    "github.com/salviati/go-qt5/qt5"
    )
    
    func main() {
	    qt5.Main(func() {
		    w := qt5.NewWidget()
		    w.SetWindowTitle(qt5.Version())
		    w.SetSizev(300, 200)
		    defer w.Close()
		    w.Show()
		    qt5.Run()
	    })
    }
