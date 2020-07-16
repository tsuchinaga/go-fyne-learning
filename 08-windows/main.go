package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"gitlab.com/tsuchinaga/go-fyne-learning/theme"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&theme.MisakiTheme{})

	subWindows := make([]fyne.Window, 0)
	w := a.NewWindow("Fyne Learn")
	w.SetContent(widget.NewVBox(
		layout.NewSpacer(),
		widget.NewButton("新しいウィンドウ", func() {
			sw := a.NewWindow("Sub Window")
			sw.SetContent(widget.NewVBox(
				layout.NewSpacer(),
				widget.NewButton("とじる", func() {
					sw.Close()
				})))
			sw.Show()
			subWindows = append(subWindows, sw)
		}),
		layout.NewSpacer(),
		widget.NewButton("とじる", func() {
			a.Quit()
		}),
	))

	w.Resize(fyne.NewSize(512, 512))
	w.ShowAndRun()
}
