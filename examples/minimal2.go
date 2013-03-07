package main

import (
	"github.com/salviati/go-qt5/qt5"
)

var exit = make(chan bool)

func main() {
	qt5.Main(func() {
		go ui_main()
		qt5.Run()
		exit <- true
	})
}

func ui_main() {
	w := qt5.NewWidget()
	w.SetWindowTitle(qt5.Version())
	w.SetSizev(300, 200)
	defer w.Close()
	w.Show()
	<-exit
}
