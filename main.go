package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"gitlab.com/tsuchinaga/go-fyne-learning/theme"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&theme.MisakiTheme{})

	w := a.NewWindow("Fyne Learn")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello, World!"),
		widget.NewLabel("こんにちわーるど！"),
		widget.NewButton("終了", func() {
			a.Quit()
		})))

	w.Resize(fyne.NewSize(256, 256))
	w.ShowAndRun()
}
