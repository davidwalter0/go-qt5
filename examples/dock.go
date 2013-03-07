package main

import (
	"fmt"
	"github.com/salviati/go-qt5/qt5"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	qt5.Main(main_ui)
}

func main_ui() {
	qt5.OnInsertObject(func(v interface{}) {
		fmt.Println("add item", v)
	})
	qt5.OnRemoveObject(func(v interface{}) {
		fmt.Println("remove item", v)
	})
	w := qt5.NewMainWindow()
	defer w.Close()
	go func() {
		dock := qt5.NewDockWidgetWithTitle("Dock")
		dock.SetDock(qt5.NewButtonWithText("Hello"))
		w.AddDockWidget(qt5.LeftDockWidgetArea, dock)
		btn := qt5.NewButtonWithText("HideDock")
		w.SetCentralWidget(btn)
		w.SetSize(qt5.Sz(200, 200))

		tb := qt5.NewToolBarWithTitle("Standard")
		tb.AddWidget(qt5.NewButtonWithText("ok"))
		w.AddToolBar(tb)

		tb.OnCloseEvent(func(e *qt5.CloseEvent) {
			fmt.Println("tb close", e)
		})
		sb := qt5.NewStatusBar()
		w.SetStatusBar(sb)
		sb.OnCloseEvent(func(e *qt5.CloseEvent) {
			fmt.Println("sb close", e)
		})

		btn.OnClicked(func() {
			dock.Hide()
			runtime.GC()
			btn.SetText(btn.Text())
		})
		dock.OnCloseEvent(func(e *qt5.CloseEvent) {
			fmt.Println(e)
		})

		go func() {
			for {
				timer := time.NewTimer(1)
				select {
				case <-timer.C:
					btn.SetText(btn.Text())
					btn.SetText(btn.Text())
					btn.SetText(btn.Text())
					fmt.Println(">", btn.Text())
					if btn.Text() != "HideDock" {
						panic("close")
					}
				}
			}
		}()

		dock.OnVisibilityChanged(func(b bool) {
			fmt.Println(b)
			if !b {
				time.AfterFunc(1e9, func() {
					if dock.IsValid() {
						dock.Show()
					}
				})
			}
		})

		w.Show()
	}()

	qt5.Run()
}
