package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"gitlab.com/tsuchinaga/go-fyne-learning/theme"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.KoruriTheme{})

	w := a.NewWindow("custom layout")
	w.SetContent(fyne.NewContainerWithLayout(new(diagonalLayout),
		canvas.NewText("こんにち", color.White),
		canvas.NewText("わーるど", color.White),
		canvas.NewText("そしてみなさん", color.White)))

	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}

type diagonalLayout struct{}

func (d diagonalLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	x, y := size.Width, size.Height
	for _, o := range objects {
		size := o.MinSize()
		x -= size.Width
		y -= size.Height

		o.Resize(size)
		o.Move(fyne.NewPos(x, y))
	}
}

func (d diagonalLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	var w, h int
	for _, o := range objects {
		childSize := o.MinSize()
		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}
