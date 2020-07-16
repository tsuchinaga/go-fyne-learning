package main

import (
	"bufio"
	"errors"
	"time"

	"fyne.io/fyne/layout"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"gitlab.com/tsuchinaga/go-fyne-learning/theme"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&theme.MisakiTheme{})

	w := a.NewWindow("Fyne Learn")
	w.SetContent(widget.NewVBox(
		layout.NewSpacer(),
		widget.NewButton("NewConfirm", func() {
			cnf := dialog.NewConfirm("確認", "fyne どうや？", func(b bool) { println("callback", b) }, w)
			cnf.SetDismissText("微妙")
			cnf.SetConfirmText("ええでこれ")
			cnf.Show()
		}),
		widget.NewButton("NewFileOpen", func() {
			fileOpen := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
				if err != nil {
					println(err)
				}
				if closer != nil {
					defer closer.Close()
					println(closer.Name())
					bc := bufio.NewScanner(closer)
					for bc.Scan() {
						println(bc.Text())
					}
				}
			}, w)
			fileOpen.SetDismissText("ほないくで")
			fileOpen.Show()
		}),
		widget.NewButton("NewError", func() {
			er := dialog.NewError(errors.New("エラーでたわ"), w)
			er.SetDismissText("まじかーい")
			er.Show()
		}),
		widget.NewButton("NewInformation", func() {
			info := dialog.NewInformation("おしらせ", "ちょっときいて", w)
			info.SetDismissText("おけまる")
			info.Show()
		}),
		widget.NewButton("NewProgressInfinite", func() {
			pr := dialog.NewProgressInfinite("処理中", "がんばってまーす", w)
			pr.Show()
			go func() {
				time.Sleep(3 * time.Second)
				pr.Hide()
			}()
		})))

	w.Resize(fyne.NewSize(512, 512))
	w.ShowAndRun()
}
