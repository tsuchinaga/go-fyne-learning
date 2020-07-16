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

	// label and button
	box1 := widget.NewVBox(
		widget.NewLabel("Hello, World!"),
		widget.NewLabel("こんにちわーるど！"),
		widget.NewButton("終了", func() {
			a.Quit()
		}))

	// entry
	box2 := widget.NewVBox(
		widget.NewForm(
			widget.NewFormItem("ID", widget.NewEntry()),
			widget.NewFormItem("PASS", widget.NewPasswordEntry()),
			widget.NewFormItem("メモ", widget.NewMultiLineEntry())),
		widget.NewButton("ログイン", func() {
			println("pressed login button")
		}))

	// radio, check and select
	radio := widget.NewRadio([]string{"OK", "NG"}, func(s string) { println(s) })
	radio.Horizontal = true
	box3 := widget.NewVBox(
		radio,
		widget.NewHBox(
			widget.NewCheck("りんご", func(b bool) { println("りんご", b) }),
			widget.NewCheck("みかん", func(b bool) { println("みかん", b) }),
			widget.NewCheck("いちご", func(b bool) { println("いちご", b) }),
		),
		widget.NewSelect([]string{"東京", "大阪", "名古屋", "福岡", "札幌"}, func(s string) { println(s) }))

	w.SetContent(widget.NewVBox(box1, box2, box3))

	w.Resize(fyne.NewSize(512, 512))
	w.ShowAndRun()
}
