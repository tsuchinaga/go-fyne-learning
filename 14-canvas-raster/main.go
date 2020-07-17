package main

import (
	"image/color"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Raster")

	raster := canvas.NewRasterWithPixels(
		func(x, y, _, _ int) color.Color {
			if (x/25+y/25+int(time.Now().Unix()%2))%2 == 0 {
				return color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}
			} else {
				return color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}
			}
		})

	go func() {
		for {
			time.Sleep(1 * time.Second)
			raster.Refresh()
		}
	}()

	w.SetContent(raster)
	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
