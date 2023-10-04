package models

import (
	"fyne.io/fyne/v2/storage"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Dish struct {
	posX, posY float32
	status     bool
	img        *canvas.Image
	generator  *BurgerPartsGenerator
}

func NewDish(generator *BurgerPartsGenerator) *Dish {
	return &Dish{
		generator: generator,
		posX:      -183,
		posY:      520,
		status:    true,
		img:       canvas.NewImageFromURI(storage.NewFileURI("./assets/dish.png")),
	}
}

func (d *Dish) GetImage() *canvas.Image {
	dishImage := d.img
	dishImage.Resize(fyne.NewSize(180, 30))
	dishImage.Move(fyne.NewPos(-183, 520))

	return dishImage
}

func (d *Dish) Run() {
	d.status = true
	for d.status {
		d.posX += 20
		if d.posX > 820 {
			d.posX = -180
			time.Sleep(400 * time.Millisecond)
			d.generator.MoveItems()
		}
		d.img.Move(fyne.NewPos(d.posX, d.posY))
		time.Sleep(100 * time.Millisecond)
	}
}

func (d *Dish) Stop() {
	d.status = false
}
