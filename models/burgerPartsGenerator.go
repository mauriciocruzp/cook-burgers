package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"math/rand"
)

type BurgerPartsGenerator struct {
	topBread    *canvas.Image
	ketchup     *canvas.Image
	lettuce     *canvas.Image
	beef        *canvas.Image
	bottomBread *canvas.Image
}

func NewBurgerPartsGenerator() *BurgerPartsGenerator {
	return &BurgerPartsGenerator{
		topBread:    canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/top_bread.png")),
		ketchup:     canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/ketchup.png")),
		lettuce:     canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/lettuce.png")),
		beef:        canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/beef.png")),
		bottomBread: canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/bottom_bread.png")),
	}
}

func (b *BurgerPartsGenerator) Run() {
	b.bottomBread.Show()
	b.ketchup.Show()
	b.lettuce.Show()
	b.beef.Show()
	b.topBread.Show()
}

func randPosition() fyne.Position {
	randXPos := 0 + rand.Intn(820-0)
	randYPos := 200 + rand.Intn(480-200)
	return fyne.NewPos(float32(randXPos), float32(randYPos))
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
	topBreadImage.Move(fyne.NewPos(340, 300))
	topBreadImage.Hide()

	return topBreadImage
}

func (b *BurgerPartsGenerator) GetKetchup() *canvas.Image {
	ketchupImage := b.ketchup
	ketchupImage.Resize(fyne.NewSize(180, 30))
	ketchupImage.Move(fyne.NewPos(40, 380))
	ketchupImage.Hide()

	return ketchupImage
}

func (b *BurgerPartsGenerator) GetLettuce() *canvas.Image {
	lettuceImage := b.lettuce
	lettuceImage.Resize(fyne.NewSize(180, 30))
	lettuceImage.Move(fyne.NewPos(480, 440))
	lettuceImage.Hide()

	return lettuceImage
}

func (b *BurgerPartsGenerator) GetBeef() *canvas.Image {
	beefImage := b.beef
	beefImage.Resize(fyne.NewSize(180, 30))
	beefImage.Move(fyne.NewPos(640, 260))
	beefImage.Hide()

	return beefImage
}

func (b *BurgerPartsGenerator) GetBottomBread() *canvas.Image {
	bottomBreadImage := b.bottomBread
	bottomBreadImage.Resize(fyne.NewSize(180, 30))
	bottomBreadImage.Move(fyne.NewPos(200, 200))
	bottomBreadImage.Hide()

	return bottomBreadImage
}
