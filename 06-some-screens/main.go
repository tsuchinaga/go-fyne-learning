package main

import (
	"fmt"

	"fyne.io/fyne/layout"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"gitlab.com/tsuchinaga/go-fyne-learning/theme"
)

var (
	w           fyne.Window
	currentPage int
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&theme.MisakiTheme{})

	w = a.NewWindow("Fyne Learn")
	defaultPage()

	w.Resize(fyne.NewSize(512, 512))
	w.ShowAndRun()
}

func NewPage(i int) fyne.CanvasObject {
	moveButtonsBox := widget.NewHBox(
		widget.NewButton("前のページ", prevPage),
		layout.NewSpacer(),
		widget.NewLabel(fmt.Sprintf("ページ%d", i+1)),
		layout.NewSpacer(),
		widget.NewButton("次のページ", nextPage),
	)
	moveButtonsBox.CreateRenderer()

	page := widget.NewVBox()
	page.Append(moveButtonsBox)
	page.Append(widget.NewButton("最初のページ", defaultPage))
	return page
}

func defaultPage() {
	currentPage = 0
	w.SetContent(NewPage(currentPage))
}

func prevPage() {
	if currentPage < 1 {
		return
	}
	currentPage--
	w.SetContent(NewPage(currentPage))
}

func nextPage() {
	currentPage++
	w.SetContent(NewPage(currentPage))
}
