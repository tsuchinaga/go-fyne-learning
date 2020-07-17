package theme

import (
	"image/color"

	"gitlab.com/tsuchinaga/go-fyne-learning/bundle"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

type KoruriTheme struct{}

// return bundled font resource
func (KoruriTheme) TextFont() fyne.Resource     { return bundle.ResorceKoruriRegularTtf }
func (KoruriTheme) TextBoldFont() fyne.Resource { return bundle.ResorceKoruriRegularTtf }

func (KoruriTheme) BackgroundColor() color.Color      { return theme.DarkTheme().BackgroundColor() }
func (KoruriTheme) ButtonColor() color.Color          { return theme.DarkTheme().ButtonColor() }
func (KoruriTheme) DisabledButtonColor() color.Color  { return theme.DarkTheme().DisabledButtonColor() }
func (KoruriTheme) IconColor() color.Color            { return theme.DarkTheme().IconColor() }
func (KoruriTheme) DisabledIconColor() color.Color    { return theme.DarkTheme().DisabledIconColor() }
func (KoruriTheme) HyperlinkColor() color.Color       { return theme.DarkTheme().HyperlinkColor() }
func (KoruriTheme) TextColor() color.Color            { return theme.DarkTheme().TextColor() }
func (KoruriTheme) TextSize() int                     { return theme.DarkTheme().TextSize() }
func (KoruriTheme) DisabledTextColor() color.Color    { return theme.DarkTheme().DisabledTextColor() }
func (KoruriTheme) HoverColor() color.Color           { return theme.DarkTheme().HoverColor() }
func (KoruriTheme) PlaceHolderColor() color.Color     { return theme.DarkTheme().PlaceHolderColor() }
func (KoruriTheme) PrimaryColor() color.Color         { return theme.DarkTheme().PrimaryColor() }
func (KoruriTheme) FocusColor() color.Color           { return theme.DarkTheme().FocusColor() }
func (KoruriTheme) ScrollBarColor() color.Color       { return theme.DarkTheme().ScrollBarColor() }
func (KoruriTheme) ShadowColor() color.Color          { return theme.DarkTheme().ShadowColor() }
func (KoruriTheme) TextItalicFont() fyne.Resource     { return theme.DarkTheme().TextItalicFont() }
func (KoruriTheme) TextBoldItalicFont() fyne.Resource { return theme.DarkTheme().TextBoldItalicFont() }
func (KoruriTheme) TextMonospaceFont() fyne.Resource  { return theme.DarkTheme().TextMonospaceFont() }
func (KoruriTheme) Padding() int                      { return theme.DarkTheme().Padding() }
func (KoruriTheme) IconInlineSize() int               { return theme.DarkTheme().IconInlineSize() }
func (KoruriTheme) ScrollBarSize() int                { return theme.DarkTheme().ScrollBarSize() }
func (KoruriTheme) ScrollBarSmallSize() int           { return theme.DarkTheme().ScrollBarSmallSize() }
