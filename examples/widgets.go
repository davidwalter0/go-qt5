package main

import (
	"fmt"
	"github.com/salviati/go-qt5/qt5"
	"io/ioutil"
	"runtime"
)

type MainWindow struct {
	qt5.Widget
	tab  *qt5.TabWidget
	sbar *qt5.StatusBar
}

func (p *MainWindow) Init() *MainWindow {
	if p.Widget.Init() == nil {
		return nil
	}
	p.SetWindowTitle("MainWindow")

	p.tab = qt5.NewTabWidget()

	p.tab.AddTab(p.createStdTab(), "Standard", nil)
	p.tab.AddTab(p.createMyTab(), "Custom", nil)
	p.tab.AddTab(p.createToolBox(), "ToolBox", nil)

	p.sbar = qt5.NewStatusBar()

	menubar := qt5.NewMenuBar()
	menu := qt5.NewMenuWithTitle("&File")
	//menu.SetTitle("&File")
	menubar.AddMenu(menu)

	act := qt5.NewAction()
	act.SetText("&Quit")
	act.OnTriggered(func(bool) {
		p.Close()
	})
	ic := qt5.NewIconWithFile("images/close.png")
	//defer ic.Close()
	act.SetIcon(ic)
	menu.AddAction(act)

	toolBar := qt5.NewToolBar()
	toolBar.AddAction(act)
	toolBar.AddSeparator()
	cmb := qt5.NewComboBox()
	cmb.AddItem("test1")
	cmb.AddItem("test2")
	cmb.SetToolTip("ComboBox")
	cmbAct := toolBar.AddWidget(cmb)
	fmt.Println(cmbAct)

	vbox := qt5.NewVBoxLayout()
	vbox.SetMargin(0)
	vbox.SetSpacing(0)
	vbox.SetMenuBar(menubar)
	vbox.AddWidget(toolBar)
	vbox.AddWidget(p.tab)
	vbox.AddWidget(p.sbar)

	p.SetLayout(vbox)

	p.tab.OnCurrentChanged(func(index int) {
		p.sbar.ShowMessage("current: "+p.tab.TabText(index), 0)
	})

	systray := qt5.NewSystemTray()
	systray.SetContextMenu(menu)
	systray.SetIcon(ic)
	systray.SetVisible(true)
	systray.ShowMessage("hello", "this is a test", qt5.Information, 1000)
	ic2 := systray.Icon()
	fmt.Println(ic2)

	p.SetWindowIcon(ic2)

	return p
}

func (p *MainWindow) createStdTab() *qt5.Widget {
	w := qt5.NewWidget()
	vbox := qt5.NewVBoxLayout()
	w.SetLayout(vbox)

	ed := qt5.NewLineEdit()
	ed.SetInputMask("0000-00-00")
	ed.SetText("2012-01-12")

	lbl := qt5.NewLabel()
	lbl.SetText("Label")
	btn := qt5.NewButton()
	btn.SetText("Button")
	chk := qt5.NewCheckBox()
	chk.SetText("CheckBox")
	radio := qt5.NewRadio()
	radio.SetText("Radio")
	cmb := qt5.NewComboBox()
	cmb.AddItem("001")
	cmb.AddItem("002")
	cmb.AddItem("003")
	cmb.SetCurrentIndex(2)
	fmt.Println(cmb.CurrentIndex())
	cmb.OnCurrentIndexChanged(func(v int) {
		fmt.Println(cmb.ItemText(v))
	})

	slider := qt5.NewSlider()
	slider.SetTickInterval(50)
	slider.SetTickPosition(qt5.TicksBothSides)
	slider.SetSingleStep(1)

	scl := qt5.NewScrollBar()
	fmt.Println(slider.Range())

	dial := qt5.NewDial()

	dial.SetNotchesVisible(true)
	dial.SetNotchTarget(10)
	fmt.Println(dial.NotchSize())

	vbox.AddWidget(ed)
	vbox.AddWidget(lbl)
	vbox.AddWidget(btn)
	vbox.AddWidget(chk)
	vbox.AddWidget(radio)
	vbox.AddWidget(cmb)
	vbox.AddWidget(slider)
	vbox.AddWidget(scl)
	vbox.AddWidget(dial)
	vbox.AddStretch(0)
	return w
}

func (p *MainWindow) createToolBox() qt5.IWidget {
	tb := qt5.NewToolBox()
	tb.AddItem(qt5.NewButtonWithText("button"), "btn", nil)
	tb.AddItem(qt5.NewLabelWithText("Label\nInfo"), "Label", nil)
	pixmap := qt5.NewPixmapWithFile("images/liteide128.png")
	//defer pixmap.Close()
	lbl := qt5.NewLabel()
	lbl.SetPixmap(pixmap)
	tb.AddItem(lbl, "Lalel Pixmap", nil)
	buf, err := ioutil.ReadFile("images/liteide128.png")
	if err == nil {
		pixmap2 := qt5.NewPixmapWithData(buf)
		tb.AddItem(qt5.NewLabelWithPixmap(pixmap2), "Lalel Pixmap2", nil)
	}
	return tb
}

func (p *MainWindow) createMyTab() *qt5.Widget {
	w := qt5.NewWidget()
	vbox := qt5.NewVBoxLayout()
	hbox := qt5.NewHBoxLayout()
	my := new(MyWidget).Init()
	lbl := qt5.NewLabel()
	lbl.SetText("this is custome widget - draw lines")
	btn := qt5.NewButton()
	btn.SetText("Clear")
	btn.OnClicked(func() {
		my.Clear()
	})
	hbox.AddWidget(lbl)
	hbox.AddWidgetWith(btn, 0, qt5.AlignRight)
	vbox.AddLayout(hbox)
	vbox.AddWidgetWith(my, 1, 0)
	w.SetLayout(vbox)
	return w
}

func main() {
	qt5.Main(ui_main)
}

func ui_main() {
	exit := make(chan bool)
	go func() {
		fmt.Println("vfc/ui")
		qt5.OnInsertObject(func(v interface{}) {
			fmt.Println("add item", v)
		})
		qt5.OnRemoveObject(func(v interface{}) {
			fmt.Println("remove item", v)
		})
		w := new(MainWindow).Init()
		defer w.Close()

		w.SetSizev(400, 300)
		w.OnCloseEvent(func(e *qt5.CloseEvent) {
			fmt.Println("close", e)
		})
		w.Show()
		<- exit
	}()
	qt5.Run()
	exit <- true
}

type MyWidget struct {
	qt5.Widget
	lines [][]qt5.Point
	line  []qt5.Point
	font  *qt5.Font
}

func (p *MyWidget) Name() string {
	return "MyWidget"
}

func (p *MyWidget) String() string {
	return qt5.DumpObject(p)
}

func (p *MyWidget) Init() *MyWidget {
	if p.Widget.Init() == nil {
		return nil
	}
	p.font = qt5.NewFontWith("Timer", 16, 87)
	p.font.SetItalic(true)
	p.Widget.OnPaintEvent(func(e *qt5.PaintEvent) {
		p.paintEvent(e)
	})
	p.Widget.OnMousePressEvent(func(e *qt5.MouseEvent) {
		p.mousePressEvent(e)
	})
	p.Widget.OnMouseMoveEvent(func(e *qt5.MouseEvent) {
		p.mouseMoveEvent(e)
	})
	p.Widget.OnMouseReleaseEvent(func(e *qt5.MouseEvent) {
		p.mouseReleaseEvent(e)
	})
	qt5.InsertObject(p)
	return p
}

func (p *MyWidget) Clear() {
	p.lines = [][]qt5.Point{}
	p.Update()
}

func (p *MyWidget) paintEvent(e *qt5.PaintEvent) {
	paint := qt5.NewPainter()
	defer paint.Close()

	paint.Begin(p)
	paint.SetFont(p.font)
	paint.DrawLines(p.line)
	paint.SetFont(p.font)
	paint.DrawText(qt5.Pt(100, 100), "draw test")
	for _, v := range p.lines {
		//paint.DrawLines(v)
		paint.DrawPolyline(v)
	}
	paint.End()
	runtime.GC()
}

func (p *MyWidget) mousePressEvent(e *qt5.MouseEvent) {
	p.line = append(p.line, e.Pos())
	p.Update()
}

func (p *MyWidget) mouseMoveEvent(e *qt5.MouseEvent) {
	p.line = append(p.line, e.Pos())
	p.Update()
}

func (p *MyWidget) mouseReleaseEvent(e *qt5.MouseEvent) {
	p.line = append(p.line, e.Pos())
	p.lines = append(p.lines, p.line)
	p.line = []qt5.Point{}
	p.Update()
}
