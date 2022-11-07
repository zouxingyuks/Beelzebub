package main

import (
	"Beelzebub/diy"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
)

type selector struct {
	window fyne.Window
	e      *encode
	d      *decode
	button diy.Button
}

func newSelector() (s *selector) {
	s = &selector{
		window: nil,
	}
	s.button.New(3)
	s.e = newEncode()
	s.d = newDecode()

	return s
}

func (s *selector) loadUI(app fyne.App) {

	//处理选择窗口UI
	s.window = app.NewWindow("模式选择")
	s.window.SetContent(
		container.NewGridWithRows(3,
			widget.NewLabel(optionsTitle),

			s.button.AddButton(encodeTitle, func() {
				s.e.loadUI(app)
			}),
			s.button.AddButton(decodeTitle, func() {
				if s.e == nil || len(ruler) == 0 {
					dialog.ShowCustomConfirm("警告", "现在就去设置", "我知道了,仍要继续", widget.NewLabel("尚未设置解码规则"), func(b bool) {
						if b {
							s.e.loadUI(app)

						} else {
							s.d.loadUI(app)

						}
					}, s.window)
					log.Println("尚未设置解码规则")
					return
				}
				s.d.loadUI(app)
			}),
		),
	)
	s.window.Resize(fyne.NewSize(400, 400))
	s.window.Show()
}
