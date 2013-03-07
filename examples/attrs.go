package main

import (
	"fmt"
	"github.com/salviati/go-qt5/qt5"
)

var exit = make(chan bool)

func main() {
	fmt.Println(qt5.Version())
	qt5.Main(func() {
		go main_ui()
		qt5.Run()
		exit <- true
	})
}

func main_ui() {
	qt5.OnInsertObject(func(item interface{}) {
		fmt.Println("add item", item)
	})
	qt5.OnRemoveObject(func(item interface{}) {
		fmt.Println("remove item", item)
	})

	w := qt5.NewWidget()
	defer w.Close()

	w.SetWindowTitle("This is a test")
	fmt.Println(w.WindowTitle())

	vbox := qt5.NewVBoxLayout()
	fmt.Println(vbox)
	w.SetLayout(vbox)

	lbl := qt5.NewLabel()
	lbl.SetText("<h2><i>Hello</i> <font color=blue><a href=\"ui\">UI</a></font></h2>")
	lbl.OnLinkActivated(fnTEST)
	vbox.AddWidget(lbl)
	vbox.AddStretch(0)

	//runtime.GC()

	btn := qt5.NewButton()
	btn.SetText("WbcdefgwqABCDEFQW")
	font := btn.Font()
	defer font.Close()
	font.SetPointSize(16)
	btn.SetFont(font)
	fmt.Println("f3->", btn.Font())

	btn2 := qt5.NewButton()
	font.SetPointSize(18)
	btn2.SetAttr("text", "WbcdefgwqABCDEFQW")
	btn2.SetAttr("font", font)

	btn.OnClicked(func() {
		fmt.Println(btn)
		btn.Close()
	})
	btn.OnCloseEvent(func(e *qt5.CloseEvent) {
		fmt.Println("Close", e)
	})
	btn3 := qt5.NewButton()
	btn3.SetText("Exit")
	btn3.OnClicked(func() {
		qt5.Exit(0)
	})

	l := w.Layout()
	fmt.Println("ll", l)
	l.AddWidget(btn)
	l.AddWidget(btn2)
	l.AddWidget(btn3)
	//vbox.AddWidget(btn)
	f := btn2.Attr("parent")
	fmt.Println("parent->", f, f == nil)

	fmt.Println(btn.Font())

	w.OnResizeEvent(func(e *qt5.ResizeEvent) {
		fmt.Println(e)
	})

	w.OnPaintEvent(func(ev *qt5.PaintEvent) {
		fnPaint(ev, w)
	})

	//w.Show()
	w.SetVisible(true)
	<-exit
}

func fnPaint(ev *qt5.PaintEvent, w *qt5.Widget) {
	p := qt5.NewPainter()
	defer p.Close()
	p.Begin(w)
	p.DrawPoint(qt5.Pt(10, 10))
	p.DrawLine(qt5.Pt(10, 10), qt5.Pt(100, 100))
	p.End()
}

func fnTEST(link string) {
	fmt.Println("link:", link)
}
