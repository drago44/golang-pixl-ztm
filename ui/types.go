package ui

import (
	"fyne.io/fyne/v2"
	"pixl-ztm/apptype"
	"pixl-ztm/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
