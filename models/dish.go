package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/storage"
)

type Dish struct {
	posX, posY float32
	status     bool
	img        *canvas.Image
	generator  *BurgerPartsGenerator
	window     fyne.Window
}

func NewDish(generator *BurgerPartsGenerator, window fyne.Window) *Dish {
	return &Dish{
		window:    window,
		generator: generator,
		posX:      410,
		posY:      520,
		status:    true,
		img:       canvas.NewImageFromURI(storage.NewFileURI("./assets/dish.png")),
	}
}

func (d *Dish) GetImage() *canvas.Image {
	dishImage := d.img
	dishImage.Resize(fyne.NewSize(180, 30))
	dishImage.Move(fyne.NewPos(410, 520))

	return dishImage
}

func (d *Dish) Run() {
	d.window.Canvas().(desktop.Canvas).SetOnKeyDown(func(event *fyne.KeyEvent) {
		step := float32(50)

		switch event.Name {
		case fyne.KeyLeft:
			if d.posX > 20 {
				d.posX -= step
			}
		case fyne.KeyRight:
			if d.posX < 800 {
				d.posX += step
			}
		}
		d.img.Move(fyne.NewPos(d.posX, d.posY))
	})
	//d.status = true
	//for d.status {
	//	d.posX += 20
	//	if d.posX > 820 {
	//		d.posX = -180
	//		time.Sleep(400 * time.Millisecond)
	//		d.generator.MoveItems()
	//	}
	//	d.img.Move(fyne.NewPos(d.posX, d.posY))
	//	time.Sleep(100 * time.Millisecond)
	//}
}

func (d *Dish) Stop() {
	d.status = false
}
