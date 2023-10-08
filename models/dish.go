package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/storage"
)

type Dish struct {
	posX, posY   float32
	status       bool
	img          *canvas.Image
	window       fyne.Window
	itemsCounter float32
	itemsOnDish  []*canvas.Image
}

func NewDish(window fyne.Window) *Dish {
	return &Dish{
		window:       window,
		posX:         410,
		posY:         520,
		status:       true,
		img:          canvas.NewImageFromURI(storage.NewFileURI("./assets/dish.png")),
		itemsCounter: 30,
		itemsOnDish:  []*canvas.Image{},
	}
}

func (d *Dish) GetImage() *canvas.Image {
	dishImage := d.img
	dishImage.Resize(fyne.NewSize(180, 30))
	dishImage.Move(fyne.NewPos(410, 520))

	return dishImage
}

func (d *Dish) GetItemsOnDish() []*canvas.Image {
	return d.itemsOnDish
}

func (d *Dish) SetItemOnDish(item *canvas.Image) {
	d.itemsOnDish = append(d.itemsOnDish, item)
	item.Move(fyne.NewPos(d.posX, d.posY-d.itemsCounter))
	d.itemsCounter = d.itemsCounter + 30
}

func (d *Dish) Run() {
	d.SetStatus(true)
	d.window.Canvas().(desktop.Canvas).SetOnKeyDown(func(event *fyne.KeyEvent) {
		if d.GetStatus() {
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
			if d.itemsCounter > 30 {
				for _, item := range d.itemsOnDish {
					item.Move(fyne.NewPos(d.posX, item.Position().Y))
				}
			}
		}
	})
}

func (d *Dish) GetStatus() bool {
	return d.status
}

func (d *Dish) SetStatus(status bool) {
	d.status = status
}
