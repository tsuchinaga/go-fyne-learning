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
	w.SetContent(widget.NewVBox(widget.NewVBox(
		widget.NewLabel("Hello, World!"),
		widget.NewLabel("こんにちわーるど！"),
		widget.NewButton("終了", func() {
			a.Quit()
		})),
		widget.NewVBox(
			widget.NewForm(
				widget.NewFormItem("ID", widget.NewEntry()),
				widget.NewFormItem("PASS", widget.NewPasswordEntry()),
				widget.NewFormItem("メモ", textWrapBreak(widget.NewMultiLineEntry()))),
			widget.NewButton("ログイン", func() {
				println("pressed login button")
			}))))

	w.Resize(fyne.NewSize(512, 512))
	w.ShowAndRun()
}

func textWrapBreak(entry *widget.Entry) fyne.CanvasObject {
	entry.Wrapping = fyne.TextWrapBreak
	return entry
}
