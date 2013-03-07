#go-qt5 0.1.1
=====

##Introduction
go-qt5 is a cross-platform golang ui tool kit, based on qt5.

Lua code that generates the wrappers (`uiobjs.go` and `cdrv.cpp`) can be found under `make`.

The wrapper code is by far incomplete. Adding new functionality usually consists of editing or adding files under `make/ui`, and updating `make/make.lua` script, and making relevant changes in `ui` and `qtdrv`.

##System
Windows / Linux / MacOS X

##License
    go-qt5 lib BSD
    qtdrv lib LGPL

##Build go-qt5 and examples

###1.get go-qt5
    $ go get github.com/salviati/go-qt5
###2.build qtdrv, need install QtSDK
    $ cd go-qt5/qtdrv
    $ qmake "CONFIG+=release"
    $ make
###3.build go-qt5
    $ cd go-qt5/ui
    $ go install
###4.build examples
    $ cd go-qt5/examples
    $ go build -ldflags '-r ../lib' minimal.go
    $ ./minimal

##Examples

    package main

    import (
	    "github.com/visualfc/go-qt5/ui"
    )
    
    func main() {
	    ui.Main(func() {
		    w := ui.NewWidget()
		    w.SetWindowTitle(ui.Version())
		    w.SetSizev(300, 200)
		    defer w.Close()
		    w.Show()
		    ui.Run()
	    })
    }


