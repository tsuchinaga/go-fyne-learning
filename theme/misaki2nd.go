package theme

import (
	"image/color"

	"gitlab.com/tsuchinaga/go-fyne-learning/bundle"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

type MisakiTheme struct{}

// return bundled font resource
func (MisakiTheme) TextFont() fyne.Resource     { return bundle.ResourceMisakigothic2ndTtf }
func (MisakiTheme) TextBoldFont() fyne.Resource { return bundle.ResourceMisakigothic2ndTtf }

func (MisakiTheme) BackgroundColor() color.Color      { return theme.DarkTheme().BackgroundColor() }
func (MisakiTheme) ButtonColor() color.Color          { return theme.DarkTheme().ButtonColor() }
func (MisakiTheme) DisabledButtonColor() color.Color  { return theme.DarkTheme().DisabledButtonColor() }
func (MisakiTheme) IconColor() color.Color            { return theme.DarkTheme().IconColor() }
func (MisakiTheme) DisabledIconColor() color.Color    { return theme.DarkTheme().DisabledIconColor() }
func (MisakiTheme) HyperlinkColor() color.Color       { return theme.DarkTheme().HyperlinkColor() }
func (MisakiTheme) TextColor() color.Color            { return theme.DarkTheme().TextColor() }
func (MisakiTheme) DisabledTextColor() color.Color    { return theme.DarkTheme().DisabledTextColor() }
func (MisakiTheme) HoverColor() color.Color           { return theme.DarkTheme().HoverColor() }
func (MisakiTheme) PlaceHolderColor() color.Color     { return theme.DarkTheme().PlaceHolderColor() }
func (MisakiTheme) PrimaryColor() color.Color         { return theme.DarkTheme().PrimaryColor() }
func (MisakiTheme) FocusColor() color.Color           { return theme.DarkTheme().FocusColor() }
func (MisakiTheme) ScrollBarColor() color.Color       { return theme.DarkTheme().ScrollBarColor() }
func (MisakiTheme) ShadowColor() color.Color          { return theme.DarkTheme().ShadowColor() }
func (MisakiTheme) TextSize() int                     { return theme.DarkTheme().TextSize() }
func (MisakiTheme) TextItalicFont() fyne.Resource     { return theme.DarkTheme().TextItalicFont() }
func (MisakiTheme) TextBoldItalicFont() fyne.Resource { return theme.DarkTheme().TextBoldItalicFont() }
func (MisakiTheme) TextMonospaceFont() fyne.Resource  { return theme.DarkTheme().TextMonospaceFont() }
func (MisakiTheme) Padding() int                      { return theme.DarkTheme().Padding() }
func (MisakiTheme) IconInlineSize() int               { return theme.DarkTheme().IconInlineSize() }
func (MisakiTheme) ScrollBarSize() int                { return theme.DarkTheme().ScrollBarSize() }
func (MisakiTheme) ScrollBarSmallSize() int           { return theme.DarkTheme().ScrollBarSmallSize() }
