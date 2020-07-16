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

	left := widget.NewVBox()
	for i := 0; i < 100; i++ {
		left.Append(widget.NewLabel("foo"))
	}
	right := widget.NewVBox()
	for i := 0; i < 100; i++ {
		right.Append(widget.NewLabel("bar"))
	}

	w := a.NewWindow("Fyne Learn")
	w.SetContent(widget.NewHSplitContainer(
		widget.NewVScrollContainer(left),
		widget.NewVScrollContainer(right)))

	w.Resize(fyne.NewSize(512, 512))
	w.ShowAndRun()
}
