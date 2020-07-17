package main

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/canvas"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"gitlab.com/tsuchinaga/go-fyne-learning/theme"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&theme.MisakiTheme{})

	randomColorText := canvas.NewText("random", color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
	go func() {
		for {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			time.Sleep(500 * time.Millisecond)
			randomColorText.Color = color.RGBA{
				R: uint8(r.Float64() * 256),
				G: uint8(r.Float64() * 256),
				B: uint8(r.Float64() * 256),
				A: 0xff,
			}
			randomColorText.Refresh()
		}
	}()

	textL := canvas.NewText("", color.RGBA{R: 0x50, G: 0x50, B: 0x50, A: 1})
	textR := canvas.NewText("", color.RGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 1})
	lrBox := widget.NewHBox(textL, textR)
	go func() {
		text := "Hello, World!"
		n := 0
		for {
			time.Sleep(200 * time.Millisecond)
			textL.Text = text[:n]
			textR.Text = text[n:]
			lrBox.Refresh()

			n = (n + 1) % (len(text) + 1)
		}
	}()

	w := a.NewWindow("Fyne Learn")
	w.SetContent(widget.NewVBox(
		canvas.NewText("red", color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}),
		canvas.NewText("green", color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}),
		canvas.NewText("blue", color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}),
		randomColorText,
		lrBox))

	w.Resize(fyne.NewSize(512, 512))
	w.ShowAndRun()
}
