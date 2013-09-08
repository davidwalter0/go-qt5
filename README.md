#go-qt5
=====

##Before you start
This is a fork of visualfc's qt4 bindings, and several critical bugs are inherited along the way. Until these bugs are fixed, **this package is not recommended for any real use**.
I don't have any time to actively work on this project, but I'll keep reviewing and merging pull requests.

##Introduction
go-qt5 provides with qt5 bindings for Go programming language, based on visualfc's go-ui library.

Lua code that generates the wrappers (`uiobjs.go` and `cdrv.cpp`) can be found under `make`.

The wrapper code is by far incomplete, so pull requests are more than welcome. Adding new functionality usually consists of editing or adding files under `make/qt5`, and updating `make/make.lua` script, and making relevant changes in `qt5` and `qtdrv`.


##License
	go-qt5 lib BSD
	qtdrv lib LGPL

##Using go-qt5

###1. get go-qt5
	$ go get github.com/salviati/go-qt5
###2. generate bindings
	$ cd $GOPATH/src/github.com/salviati/go-qt5/make
	$ lua make.lua
	$ lua makelib.lua
###3. build & install C layer
	$ cd $GOPATH/src/github.com/salviati/go-qt5/goqtdrv5
	$ qmake "CONFIG+=release"
	$ make
	# make install
###4.build go-qt5
	$ cd $GOPATH/src/github.com/salviati/go-qt5/qt5
	$ go install
###5.build examples
	$ cd $GOPATH/src/github.com/salviati/go-qt5/examples
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
