package main

import (
	"fmt"
	"github.com/salviati/go-qt5/qt5"
)

func main() {
	qt5.Main(ui_main)
}

func ui_main() {
	w := qt5.NewWidget()
	defer w.Close()
	list := qt5.NewListWidget()
	vbox := qt5.NewVBoxLayout()
	vbox.AddWidget(list)
	w.SetLayout(vbox)
	go func() {
		list.OnCurrentItemChanged(func(item, old *qt5.ListWidgetItem) {
			go func() {
				fmt.Println(item.Attr("text"))
			}()
		})

		item := qt5.NewListWidgetItem()
		item.SetText("Item1")
		list.AddItem(item)
		list.AddItem(qt5.NewListWidgetItemWithText("Item2"))
	}()

	w.Show()

	qt5.Run()
}
