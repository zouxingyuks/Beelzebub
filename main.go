package main

import "fyne.io/fyne/v2/app"

func main() {
	a := app.New()
	a.Settings().SetTheme(&myTheme{})
	a.SetIcon(resourceIconIco)
	s := newSelector()
	s.loadUI(a)
	a.Run()
}
