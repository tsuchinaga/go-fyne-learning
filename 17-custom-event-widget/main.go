package main

import (
	"image/color"

	"fyne.io/fyne/driver/desktop"

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

	w := a.NewWindow("custom event widget")
	w.SetContent(widget.NewVBox(newCustomWidget("カスタムイベントウィジェット")))

	w.ShowAndRun()
}

func newCustomWidget(text string) fyne.CanvasObject {
	wid := &customWidget{text: text}
	wid.ExtendBaseWidget(wid)
	return wid
}

// customWidget - Widget interfaceを実装する独自ウィジェット
// 他のwidgetと同じく、BaseWidgetを埋め込んでる
type customWidget struct {
	widget.BaseWidget
	text  string
	hover bool
	posX  int
	posY  int
}

func (w *customWidget) Dragged(event *fyne.DragEvent) {
	w.posY += event.DraggedY
	w.Refresh()
}

func (w *customWidget) DragEnd() {
	w.Refresh()
}

func (w *customWidget) MouseIn(*desktop.MouseEvent) {
	w.hover = true
	w.Refresh()
}

func (w *customWidget) MouseOut() {
	w.hover = false
	w.Refresh()
}

func (w *customWidget) MouseMoved(*desktop.MouseEvent) {}

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

type customWidgetRenderer struct {
	customWidget *customWidget
	text         *canvas.Text
	objects      []fyne.CanvasObject
}

// Layout - widgetをどのように描画するか
func (r *customWidgetRenderer) Layout(fyne.Size) {
	r.customWidget.Move(fyne.NewPos(r.customWidget.posX, r.customWidget.posY))
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

func (r *customWidgetRenderer) BackgroundColor() color.Color {
	if r.customWidget.hover {
		return color.White
	} else {
		return color.Black
	}
}

func (r *customWidgetRenderer) Objects() []fyne.CanvasObject { return r.objects }
func (r *customWidgetRenderer) Destroy()                     {}
