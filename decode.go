package main

import "fyne.io/fyne/v2"

type decode struct {
	output  string //输出部分
	windows fyne.Window
}

func newDecode() *decode {
	return &decode{
		output:  "",
		windows: nil,
	}
}
func (d *decode) loadUI(app fyne.App) interface{} {
	d.windows = app.NewWindow(decodeTitle)
	d.windows.Show()
	return d
}
