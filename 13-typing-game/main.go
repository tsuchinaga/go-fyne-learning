package main

import (
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"gitlab.com/tsuchinaga/go-fyne-learning/13-typing-game/words"

	"fyne.io/fyne/canvas"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"gitlab.com/tsuchinaga/go-fyne-learning/theme"
)

var (
	white = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	gray  = color.RGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xff}
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.KoruriTheme{})

	w := a.NewWindow("typing")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	word := words.Words[r.Intn(len(words.Words))]

	enText := canvas.NewText(word.En, white)
	enText.TextSize = enText.TextSize * 2

	inputtedEnText := canvas.NewText("", gray)
	inputtedEnText.TextSize = inputtedEnText.TextSize * 2

	jaText := canvas.NewText(word.Ja, white)

	okText := canvas.NewText("0", white)
	ngText := canvas.NewText("0", white)

	w.SetContent(widget.NewVBox(
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(layout.NewCenterLayout(), widget.NewHBox(inputtedEnText, enText)),
		fyne.NewContainerWithLayout(layout.NewCenterLayout(), jaText),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(layout.NewCenterLayout(), widget.NewHBox(
			widget.NewLabel("OK: "), okText,
			widget.NewLabel("NG: "), ngText))))

	var cnt, ok, ng int
	w.Canvas().SetOnTypedRune(func(rn rune) {
		if word.En[cnt] != uint8(rn) {
			ng++
			ngText.Text = strconv.Itoa(ng)
			defer ngText.Refresh()
		} else {
			cnt++
			ok++
			if cnt == len(word.En) {
				word = words.Words[r.Intn(len(words.Words))]
				cnt = 0
				jaText.Text = word.Ja
				defer jaText.Refresh()
			}

			okText.Text = strconv.Itoa(ok)
			defer okText.Refresh()

			enText.Text = word.En[cnt:]
			inputtedEnText.Text = word.En[:cnt]
			defer enText.Refresh()
			defer inputtedEnText.Refresh()
		}
	})

	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
