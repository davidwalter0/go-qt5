// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qt5

/*
#include <stdlib.h>
extern int drv(int id, int action, void *exp, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6);

static void init_callback(int classid, int drvid)
{
	extern int drv_callback(void*,void*,void*,void*,void*);
	drv(classid,drvid,&drv_callback,0,0,0,0,0,0,0);
}

static void init_appmain(int classid, int drvid)
{
	extern void drv_appmain();
	drv(classid,drvid,&drv_appmain,0,0,0,0,0,0,0);
}

static void init_result(int classid, int drvid)
{
	extern int drv_result(void*,int);
	drv(classid,drvid,&drv_result,0,0,0,0,0,0,0);
}

static void init_utf8_info(int classid, int drvid)
{
	extern void utf8_info_copy(void*,void*,int);
	drv(classid,drvid,&utf8_info_copy,0,0,0,0,0,0,0);
}

*/
// #cgo LDFLAGS: -lgoqt5drv
import "C"
import "unsafe"
import "fmt"
import "os"
import "reflect"
import "image/color"

type string_info struct {
	Data uintptr
	Len  int
}

type rgba uint32

func (c rgba) RGBA() (r, g, b, a uint32) {
	return uint32((c >> 16) & 0xff), uint32((c >> 8) & 0xff), uint32(c & 0xff), uint32(c >> 24)
}

func make_rgba(c color.Color) rgba {
	if c == nil {
		return 0
	}
	r, g, b, a := c.RGBA()
	return rgba(((a & 0xff) << 24) | ((r & 0xff) << 16) | ((g & 0xff) << 8) | (b & 0xff))
}

type slice_info struct {
	Data uintptr
	Len  int
	Cap  int
}

type utf8_info struct {
	data []byte
}

func (d utf8_info) String() string {
	return string(d.data)
}

//export utf8_info_copy
func utf8_info_copy(p unsafe.Pointer, data unsafe.Pointer, size C.int) {
	((*utf8_info)(p)).data = C.GoBytes(data, size)
}

func init() {
	C.init_callback(-1, 0)
	C.init_result(-2, 0)
	C.init_utf8_info(-3, 0)
	C.init_appmain(-4, 0)
}

var func_map = make(map[unsafe.Pointer]interface{})

func _drv(id int, act int, a0, a1, a2, a3, a4, a5, a6, a7, a8, a9 unsafe.Pointer) int {
	return int(C.drv(C.int(id), C.int(act), nil, a0, a1, a2, a3, a4, a5, a6))
}

func _drv_ch(id int, act int, a0, a1, a2, a3, a4, a5, a6, a7, a8, a9 unsafe.Pointer) int {
	ch := make(chan int)
	C.drv(C.int(id), C.int(act), unsafe.Pointer(&ch), a0, a1, a2, a3, a4, a5, a6)
	<-ch
	return 0
}

func _drv_event(id int, event int, obj iobj, fn interface{}) {
	var pfn unsafe.Pointer = unsafe.Pointer(&fn)
	func_map[pfn] = fn
	_drv(id, event, unsafe.Pointer(obj.info()), pfn, nil, nil, nil, nil, nil, nil, nil, nil)
}

func _drv_event_ch(id int, event int, obj iobj, fn interface{}) {
	var pfn unsafe.Pointer = unsafe.Pointer(&fn)
	func_map[pfn] = fn
	_drv_ch(id, event, unsafe.Pointer(obj.info()), pfn, nil, nil, nil, nil, nil, nil, nil, nil)
}

func nativeToObject(native uintptr, classid int) interface{} {
	if native == 0 {
		return nil
	}
	obj := FindObject(native)
	if obj == nil {
		obj = NewObjectWithNative(classid, native)
	}
	return obj
}

//export drv_result
func drv_result(ch unsafe.Pointer, r int32) int32 {
	go func() {
		*(*chan int)(ch) <- int(r)
	}()
	return 0
}

var theApp app
var fnAppMain func()

//export drv_appmain
func drv_appmain() {
	theApp.onRemoveObject(drvRemoveObject)
	registerAllClass()
	defer clearAllObject()
	if fnAppMain != nil {
		fnAppMain()
	}
}

//export drv_callback
func drv_callback(pfn unsafe.Pointer, a1, a2, a3, a4 unsafe.Pointer) int32 {
	fn, ok := func_map[pfn]
	if !ok {
		return 0
	}
	switch v := (fn).(type) {
	case func():
		v()
	case func(int):
		v(*(*int)(a1))
	case func(uint):
		v(*(*uint)(a1))
	case func(bool):
		v(*(*bool)(a1))
	case func(uintptr):
		v(uintptr(a1))
	case func(string):
		v(((*utf8_info)(a1)).String())
	case func(*Action):
		obj := nativeToObject(uintptr(a1), CLASSID_ACTION)
		var act *Action
		if v1, ok := obj.(*Action); ok {
			act = v1
		}
		v(act)
	case func(*ListWidgetItem):
		obj := nativeToObject(uintptr(a1), CLASSID_LISTWIDGETITEM)
		var item *ListWidgetItem
		if v1, ok := obj.(*ListWidgetItem); ok {
			item = v1
		}
		v(item)
	case func(*ListWidgetItem, *ListWidgetItem):
		obj1 := nativeToObject(uintptr(a1), CLASSID_LISTWIDGETITEM)
		obj2 := nativeToObject(uintptr(a2), CLASSID_LISTWIDGETITEM)
		var item *ListWidgetItem
		var oldItem *ListWidgetItem
		if v1, ok := obj1.(*ListWidgetItem); ok {
			item = v1
		}
		if v2, ok := obj2.(*ListWidgetItem); ok {
			oldItem = v2
		}
		v(item, oldItem)
	case func(*ShowEvent):
		v((*ShowEvent)(a1))
	case func(*HideEvent):
		v((*HideEvent)(a1))
	case func(*CloseEvent):
		v((*CloseEvent)(a1))
	case func(*KeyEvent):
		v((*KeyEvent)(a1))
	case func(*MouseEvent):
		v((*MouseEvent)(a1))
	case func(*WheelEvent):
		v((*WheelEvent)(a1))
	case func(*MoveEvent):
		v((*MoveEvent)(a1))
	case func(*ResizeEvent):
		v((*ResizeEvent)(a1))
	case func(*EnterEvent):
		v((*EnterEvent)(a1))
	case func(*LeaveEvent):
		v((*LeaveEvent)(a1))
	case func(*FocusEvent):
		v((*FocusEvent)(a1))
	case func(*PaintEvent):
		v((*PaintEvent)(a1))
	case func(*TimerEvent):
		v((*TimerEvent)(a1))
	case func(int, int):
		v(*(*int)(a1), *(*int)(a2))

	default:
		warning("Warning drv_callback, func type \"%s\" not match!", reflect.TypeOf(v))
	}
	return 1
}

func warning(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
