package main

import (
	"log"
	"runtime"

	"golang.org/x/text/transform"

	"golang.org/x/text/encoding/japanese"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"gitlab.com/tsuchinaga/go-fyne-learning/theme"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {
	a := app.New()
	a.Settings().SetTheme(&theme.MisakiTheme{})

	w := a.NewWindow("Fyne Learn")
	w.SetContent(fyne.NewContainerWithLayout(layout.NewCenterLayout(),
		widget.NewVBox(widget.NewButton("ですくとっぷつうち", func() {
			title := "Fyneの練習"
			content := "通知のテスト"
			if runtime.GOOS == "windows" {
				enc := japanese.ShiftJIS.NewEncoder()
				var err error
				if title, _, err = transform.String(enc, title); err != nil {
					log.Println(err)
					return
				}
				if content, _, err = transform.String(enc, content); err != nil {
					log.Println(err)
					return
				}
			}

			fyne.CurrentApp().SendNotification(fyne.NewNotification(title, content))
		}))))

	w.Resize(fyne.NewSize(512, 512))
	w.ShowAndRun()
}
