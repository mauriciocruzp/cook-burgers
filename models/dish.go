package models

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Dish struct {
	posX, posY float32
	status     bool
	img        *canvas.Image
}

func NewDish(img *canvas.Image, posx float32, posy float32) *Dish {
	return &Dish{
		posX:   posx,
		posY:   posy,
		status: true,
		img:    img,
	}
}

func (d *Dish) Run() {
	d.status = true
	for d.status {
		d.posX += 20
		if d.posX > 620 {
			d.posX = -180
			time.Sleep(500 * time.Millisecond)
		}
		d.img.Move(fyne.NewPos(d.posX, d.posY))
		time.Sleep(100 * time.Millisecond)

	}
}

func (d *Dish) Stop() {
	d.status = false
}
