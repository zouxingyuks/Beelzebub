package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type encode struct {
	output string //输出部分

	buttons map[string]*widget.Button
	windows fyne.Window

	key     binding.String
	content binding.String
	ruler   map[string]string
}

func newEncode() (e *encode) {
	e = &encode{
		output:  "",
		windows: nil,
		buttons: make(map[string]*widget.Button, 20),
		ruler:   make(map[string]string, 100), //暂定最多能有100条规则
	}
	e.key = binding.NewString()
	e.content = binding.NewString()
	return e
}
func (e *encode) addButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	e.buttons[text] = button
	return button
}
func (e *encode) addEntry(text binding.String) *widget.Entry {
	entry := widget.NewEntryWithData(text)
	return entry
}
func (e *encode) loadUI(app fyne.App) interface{} {
	e.windows = app.NewWindow(encodeTitle)
	e.windows.SetContent(container.NewGridWithColumns(2,
		container.NewGridWithRows(2,
			e.addButton(s_save, func() {
				//todo 保存文件
			}),
			e.addButton(s_load, func() {
				//todo 加载文件
			}),
		),
		container.NewGridWithRows(3,
			container.NewGridWithColumns(3,
				e.addButton(s_add, func() {
					//todo  添加规则
				}),
				e.addEntry(e.key),
				e.addEntry(e.content),
			),
		),
	))
	e.windows.Show()
	return e
}
