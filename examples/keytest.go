package main

import (
	"fmt"
	"github.com/salviati/go-qt5/qt5"
)

func printInfo() {
	info := qt5.Value("info")
	fmt.Println(info)
}

func setInfo() {
	qt5.SetValue("info", "new info")
}

func main() {
	qt5.Main(ui_main)
}

func ui_main() {

	w := qt5.NewWidget()

	lbox := qt5.NewVBoxLayout()
	lbl1 := qt5.NewLabel()
	lbl1.SetText("This is a info1")
	lbl2 := qt5.NewLabel()
	lbl2.SetText("This is a info2")

	ed1 := qt5.NewLineEdit()

	printInfo := func() {
		info := qt5.Value("info")
		ed1.SetText(fmt.Sprint(info))
	}

	lbox.AddWidget(lbl1)
	lbox.AddWidget(lbl2)
	lbox.AddWidget(ed1)

	rbox := qt5.NewVBoxLayout()

	btn1 := qt5.NewButton()
	btn1.SetText("Change")

	btn2 := qt5.NewButton()
	btn2.SetText("Value")

	btn3 := qt5.NewButton()
	btn3.SetText("SetValue")

	rbox.AddWidget(btn1)
	rbox.AddWidget(btn2)
	rbox.AddWidget(btn3)

	b := true
	btn1.OnClicked(func() {
		var text string
		if b {
			qt5.SetKey(lbl1, "info", "text")
			text = "info1"
		} else {
			qt5.SetKey(lbl2, "info", "text")
			text = "info2"
		}
		b = !b
		btn1.SetText(text)
	})

	btn2.OnClicked(printInfo)
	btn3.OnClicked(setInfo)

	hbox := qt5.NewHBoxLayout()
	hbox.AddLayout(lbox)
	hbox.AddLayout(rbox)

	w.SetLayout(hbox)

	w.Show()

	qt5.Run()
}
