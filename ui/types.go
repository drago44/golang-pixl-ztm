package ui

import (
	"fyne.io/fyne/v2"
	"pixl-ztm/apptype"
	"pixl-ztm/pxcanvas"
	"pixl-ztm/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
