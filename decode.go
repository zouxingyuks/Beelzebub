package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type decode struct {
	input   binding.String
	output  binding.String //输出部分
	button  widget.Button
	windows fyne.Window
	uiW     float32
	uiH     float32
}

func newDecode() (d *decode) {
	d = &decode{
		windows: nil,
		uiH:     100,
		uiW:     400,
	}
	d.input = binding.NewString()
	d.output = binding.NewString()
	return d
}

// Input 输入框
func (d *decode) Input(Holder string) (entry *widget.Entry) {
	entry = widget.NewEntryWithData(d.input)
	entry.PlaceHolder = Holder
	return entry
}

// Output 输出
func (d *decode) Output() (label *widget.Label) {
	label = widget.NewLabelWithData(d.output)
	return label
}
func (d *decode) Decode() (b *widget.Button) {
	b = widget.NewButton(decodeTitle, func() {
		//todo 解码按钮
		input, _ := d.input.Get()
		output := Translator(input)
		d.output.Set(output)
	})
	return b
}

func (d *decode) loadUI(app fyne.App) interface{} {
	d.windows = app.NewWindow(decodeTitle)
	d.windows.SetContent(container.NewGridWithColumns(2,
		d.Decode(),
		container.NewGridWithRows(2,
			d.Input("请输入合法密文"),
			d.Output(),
		),
	))
	d.windows.Resize(fyne.NewSize(d.uiW, d.uiH))
	d.windows.Show()
	return d
}
