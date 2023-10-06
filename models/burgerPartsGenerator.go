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
	dish        *Dish
}

func NewBurgerPartsGenerator(dish *Dish) *BurgerPartsGenerator {
	return &BurgerPartsGenerator{
		topBread:    canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/top_bread.png")),
		ketchup:     canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/ketchup.png")),
		lettuce:     canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/lettuce.png")),
		beef:        canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/beef.png")),
		bottomBread: canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/bottom_bread.png")),
		status:      true,
		dish:        dish,
	}
}

func (b *BurgerPartsGenerator) Run() {
	for b.status {
		b.ResetItem(b.bottomBread)
		b.bottomBread.Show()
		go b.CollapseItem(b.bottomBread)

		time.Sleep(time.Second * 4)

		b.ResetItem(b.ketchup)
		b.ketchup.Show()
		go b.CollapseItem(b.ketchup)

		time.Sleep(time.Second * 4)

		b.ResetItem(b.lettuce)
		b.lettuce.Show()
		go b.CollapseItem(b.lettuce)
		time.Sleep(time.Second * 4)

		b.ResetItem(b.beef)
		b.beef.Show()
		go b.CollapseItem(b.beef)
		time.Sleep(time.Second * 4)

		b.ResetItem(b.topBread)
		b.topBread.Show()
		go b.CollapseItem(b.topBread)
		time.Sleep(time.Second * 6)
		b.hideAllImages()
	}
}

func (b *BurgerPartsGenerator) hideAllImages() {
	b.bottomBread.Hide()
	b.ketchup.Hide()
	b.lettuce.Hide()
	b.beef.Hide()
	b.topBread.Hide()
}

func randPosition() fyne.Position {
	randXPos := 0 + rand.Intn(820-0)
	return fyne.NewPos(float32(randXPos), 0)
}

func (b *BurgerPartsGenerator) CollapseItem(image *canvas.Image) {
	for b.status {
		image.Move(fyne.NewPos(image.Position().X, image.Position().Y+5))
		time.Sleep(time.Millisecond * 40)
		if image.Position().X-100 < b.dish.img.Position().X && image.Position().X+280 > b.dish.img.Position().X+180 && image.Position().Y == 490 {
			b.dish.SetItemOnDish(image)
			break
		}
	}
}

func (b *BurgerPartsGenerator) ResetItem(image *canvas.Image) {
	image.Move(randPosition())
}

func (b *BurgerPartsGenerator) GetTopBread() *canvas.Image {
	topBreadImage := b.topBread
	topBreadImage.Resize(fyne.NewSize(180, 30))
	topBreadImage.Move(randPosition())
	topBreadImage.Hide()

	return topBreadImage
}

func (b *BurgerPartsGenerator) GetKetchup() *canvas.Image {
	ketchupImage := b.ketchup
	ketchupImage.Resize(fyne.NewSize(180, 30))
	ketchupImage.Move(randPosition())
	ketchupImage.Hide()

	return ketchupImage
}

func (b *BurgerPartsGenerator) GetLettuce() *canvas.Image {
	lettuceImage := b.lettuce
	lettuceImage.Resize(fyne.NewSize(180, 30))
	lettuceImage.Move(randPosition())
	lettuceImage.Hide()

	return lettuceImage
}

func (b *BurgerPartsGenerator) GetBeef() *canvas.Image {
	beefImage := b.beef
	beefImage.Resize(fyne.NewSize(180, 30))
	beefImage.Move(randPosition())
	beefImage.Hide()

	return beefImage
}

func (b *BurgerPartsGenerator) GetBottomBread() *canvas.Image {
	bottomBreadImage := b.bottomBread
	bottomBreadImage.Resize(fyne.NewSize(180, 30))
	bottomBreadImage.Move(randPosition())
	bottomBreadImage.Hide()

	return bottomBreadImage
}

func (b *BurgerPartsGenerator) Stop() {
	b.status = false
}
