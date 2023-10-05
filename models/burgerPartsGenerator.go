package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"math/rand"
	"time"
)

type BurgerPartsGenerator struct {
	topBread    *canvas.Image
	ketchup     *canvas.Image
	lettuce     *canvas.Image
	beef        *canvas.Image
	bottomBread *canvas.Image
	status      bool
}

func NewBurgerPartsGenerator() *BurgerPartsGenerator {
	return &BurgerPartsGenerator{
		topBread:    canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/top_bread.png")),
		ketchup:     canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/ketchup.png")),
		lettuce:     canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/lettuce.png")),
		beef:        canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/beef.png")),
		bottomBread: canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/bottom_bread.png")),
		status:      true,
	}
}

func (b *BurgerPartsGenerator) Run() {
	for b.status {
		b.bottomBread.Show()
		go b.CollapseItem(b.bottomBread)

		time.Sleep(time.Second * 3)

		b.ketchup.Show()
		go b.CollapseItem(b.ketchup)

		time.Sleep(time.Second * 3)

		b.lettuce.Show()
		go b.CollapseItem(b.lettuce)
		time.Sleep(time.Second * 3)

		b.beef.Show()
		go b.CollapseItem(b.beef)
		time.Sleep(time.Second * 3)

		b.topBread.Show()
		go b.CollapseItem(b.topBread)
		time.Sleep(time.Second * 3)
	}
}

func randPosition() fyne.Position {
	randXPos := 0 + rand.Intn(820-0)
	return fyne.NewPos(float32(randXPos), 0)
}

func (b *BurgerPartsGenerator) CollapseItem(image *canvas.Image) {
	for b.status {
		image.Move(fyne.NewPos(image.Position().X, image.Position().Y+5))
		time.Sleep(time.Millisecond * 20)
	}
}

func (b *BurgerPartsGenerator) MoveItems() {
	b.topBread.Move(randPosition())
	b.ketchup.Move(randPosition())
	b.lettuce.Move(randPosition())
	b.beef.Move(randPosition())
	b.bottomBread.Move(randPosition())
}

func (b *BurgerPartsGenerator) GetTopBread() *canvas.Image {
	topBreadImage := b.topBread
	topBreadImage.Resize(fyne.NewSize(180, 30))
	topBreadImage.Move(fyne.NewPos(340, 0))
	topBreadImage.Hide()

	return topBreadImage
}

func (b *BurgerPartsGenerator) GetKetchup() *canvas.Image {
	ketchupImage := b.ketchup
	ketchupImage.Resize(fyne.NewSize(180, 30))
	ketchupImage.Move(fyne.NewPos(40, 0))
	ketchupImage.Hide()

	return ketchupImage
}

func (b *BurgerPartsGenerator) GetLettuce() *canvas.Image {
	lettuceImage := b.lettuce
	lettuceImage.Resize(fyne.NewSize(180, 30))
	lettuceImage.Move(fyne.NewPos(480, 0))
	lettuceImage.Hide()

	return lettuceImage
}

func (b *BurgerPartsGenerator) GetBeef() *canvas.Image {
	beefImage := b.beef
	beefImage.Resize(fyne.NewSize(180, 30))
	beefImage.Move(fyne.NewPos(640, 0))
	beefImage.Hide()

	return beefImage
}

func (b *BurgerPartsGenerator) GetBottomBread() *canvas.Image {
	bottomBreadImage := b.bottomBread
	bottomBreadImage.Resize(fyne.NewSize(180, 30))
	bottomBreadImage.Move(fyne.NewPos(200, 0))
	bottomBreadImage.Hide()

	return bottomBreadImage
}

func (b *BurgerPartsGenerator) Stop() {
	b.status = false
}
