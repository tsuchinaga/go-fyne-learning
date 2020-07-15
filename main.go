package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Fyne Learn")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello, World!")))

	w.Resize(fyne.NewSize(256, 256))
	w.ShowAndRun()
}
