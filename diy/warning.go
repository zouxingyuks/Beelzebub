package diy

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func EmptyWarning(windows fyne.Window) {
	dialog.ShowCustom("警告", "懂了懂了，马上就改", container.NewGridWithRows(1, widget.NewLabel("输个空串你想表示啥")), windows)

}
