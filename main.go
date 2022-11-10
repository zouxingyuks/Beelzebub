package main

import (
	"Beelzebub/controller"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&MyTheme{})
	a.SetIcon(resourceIconIco)
	s := controller.NewSelector()
	s.LoadUI(a)
	a.Run()
}
