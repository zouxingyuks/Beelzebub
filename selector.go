package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type selector struct {
	window  fyne.Window
	e       *encode
	d       *decode
	buttons map[string]*widget.Button
}

func (s *selector) addButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	s.buttons[text] = button
	return button
}
func newSelector() (s *selector) {
	s = &selector{
		window:  nil,
		buttons: make(map[string]*widget.Button, 3),
	}
	s.e = newEncode()
	s.d = newDecode()

	return s
}

func (s *selector) loadUI(app fyne.App) {

	//处理选择窗口UI
	s.window = app.NewWindow("模式选择")
	s.window.SetContent(
		container.NewGridWithRows(2,
			widget.NewLabel(optionsTitle),
			container.NewGridWithColumns(2,
				s.addButton(encodeTitle, func() {
					s.e.loadUI(app)
				}),
				s.addButton(decodeTitle, func() {
					s.d.loadUI(app)
				}),
			),
		))
	s.window.Resize(fyne.NewSize(200, 100))
	s.window.Show()
}
