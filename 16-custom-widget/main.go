package main

import (
	"image/color"

	"fyne.io/fyne/theme"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	myTheme "gitlab.com/tsuchinaga/go-fyne-learning/theme"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(myTheme.KoruriTheme{})

	w := a.NewWindow("custom widget")
	w.SetContent(widget.NewVBox(newCustomWidget("カスタムウィジェット")))

	w.ShowAndRun()
}

func newCustomWidget(text string) fyne.CanvasObject {
	return &customWidget{text: text}
}

// customWidget - Widget interfaceを実装する独自ウィジェット
// 他のwidgetと同じく、BaseWidgetを埋め込んでる
type customWidget struct {
	widget.BaseWidget
	text string
}

// CreateRenderer - ウィジェットのレンダラーを作る
// ウィジェットの描画をする構造体を返すやつ
func (w *customWidget) CreateRenderer() fyne.WidgetRenderer {
	text := canvas.NewText(w.text, theme.TextColor())
	return &customWidgetRenderer{
		customWidget: w,
		text:         text,
		objects:      []fyne.CanvasObject{text},
	}
}

// MinSize - 自身の最小サイズを返す
func (w *customWidget) MinSize() fyne.Size {
	w.ExtendBaseWidget(w)
	return w.BaseWidget.MinSize()
}

type customWidgetRenderer struct {
	customWidget *customWidget
	text         *canvas.Text
	objects      []fyne.CanvasObject
}

// Layout - widget内に何をどのようにレイアウトするか
func (r *customWidgetRenderer) Layout(fyne.Size) {
	r.text.Move(fyne.NewPos(theme.Padding()*2, theme.Padding()*1))
	r.text.Resize(r.text.MinSize())
}

// MinSize - ウィジェットの最小サイズ
func (r *customWidgetRenderer) MinSize() fyne.Size {
	size := r.text.MinSize()
	size = size.Add(fyne.NewSize(theme.Padding()*4, theme.Padding()*2))
	return size
}

// Refresh - 再描画
func (r *customWidgetRenderer) Refresh() {
	r.Layout(r.customWidget.MinSize())
	canvas.Refresh(r.customWidget)
}

func (r *customWidgetRenderer) BackgroundColor() color.Color { return theme.BackgroundColor() }
func (r *customWidgetRenderer) Objects() []fyne.CanvasObject { return r.objects }
func (r *customWidgetRenderer) Destroy()                     {}
