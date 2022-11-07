package diy

import "fyne.io/fyne/v2/widget"

type Button struct {
	ButtonNum int
	buttons   map[string]*widget.Button
}

func (b *Button) AddButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	b.buttons[text] = button
	return button
}

// New 生成button
func (b *Button) New(n int) { //n为组成的按钮数
	b.buttons = make(map[string]*widget.Button, n)
}
