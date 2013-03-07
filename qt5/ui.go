// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qt5

const (
	name = "go-qt5"
	version = "0.1.1"
)

func About() string {
	return name + " " + version
}

func Version() string {
	return version
}

func Main(fn func()) int {
	fnAppMain = fn
	return theApp.AppMain()
}

func Run() int {
	return theApp.Run()
}

func Exit(code int) {
	theApp.Exit(code)
}

func CloseAllWindows() {
	theApp.CloseAllWindows()
}

func App() *app {
	return &theApp
}
