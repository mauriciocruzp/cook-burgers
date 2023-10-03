package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type Conveyor struct {
	image *canvas.Image
}

func NewConveyor() *Conveyor {
	return &Conveyor{
		image: canvas.NewImageFromURI(storage.NewFileURI("./assets/conveyor.png")),
	}
}

func (c *Conveyor) GetImage() *canvas.Image {
	conveyorImage := c.image
	conveyorImage.Resize(fyne.NewSize(1052, 100))
	conveyorImage.Move(fyne.NewPos(-30, 550))

	return conveyorImage
}
