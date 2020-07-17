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
	box := widget.NewVBox()
	w.SetContent(widget.NewVScrollContainer(box))
	w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		box.Prepend(widget.NewLabel(string(event.Name)))
		for 100 < len(box.Children) {
			box.Children = box.Children[:100]
		}
	})

	w.Resize(fyne.NewSize(512, 512))
	w.ShowAndRun()
}
