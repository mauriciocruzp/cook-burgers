package models

import (
	"fyne.io/fyne/v2/canvas"
)

type Conveyor struct {
	status    bool
	canvasImg *canvas.Image
}

func NewConveyor(img *canvas.Image) *Conveyor {
	return &Conveyor{
		status:    true,
		canvasImg: img,
	}
}
