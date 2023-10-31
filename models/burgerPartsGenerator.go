package models

import (
	"math/rand"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type BurgerPartsGenerator struct {
	burgerParts    []*canvas.Image
	status         bool
	dish           *Dish
	timeToCollapse time.Duration
	gameStatus     bool
}

func NewBurgerPartsGenerator(dish *Dish) *BurgerPartsGenerator {
	bottomBread := canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/bottom_bread.png"))
	lettuce := canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/lettuce.png"))
	beef := canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/beef.png"))
	ketchup := canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/ketchup.png"))
	topBread := canvas.NewImageFromURI(storage.NewFileURI("./assets/burger/top_bread.png"))

	return &BurgerPartsGenerator{
		burgerParts:    []*canvas.Image{bottomBread, lettuce, beef, ketchup, topBread},
		status:         true,
		dish:           dish,
		timeToCollapse: 40,
		gameStatus:     true,
	}
}

func (b *BurgerPartsGenerator) Run(wg *sync.WaitGroup, quit chan bool) {
	defer wg.Done()
	b.SetStatus(true)

	for b.status {
		for _, item := range b.burgerParts {
			select {
			case <-quit:
				return
			default:
				b.ResetItem(item)
				item.Show()
				go b.CollapseItem(item)

				time.Sleep(time.Second * 4)
			}
		}

		b.HideAllImages()
		b.dish.itemsCounter = 30
		b.dish.DeleteItemsOnDish()
		if b.timeToCollapse <= 10 && b.timeToCollapse > 3 {
			b.timeToCollapse = b.timeToCollapse - 1
		} else if b.timeToCollapse > 10 {
			b.timeToCollapse = b.timeToCollapse - 10
		}
	}
}

func (b *BurgerPartsGenerator) Stop() {
	b.SetStatus(false)
	b.dish.DeleteItemsOnDish()
	b.HideAllImages()
}

func (b *BurgerPartsGenerator) GetBurgerParts() []*canvas.Image {
	burgerParts := b.burgerParts

	for _, item := range burgerParts {
		item.Resize(fyne.NewSize(180, 30))
		item.Move(randPosition())
		item.Hide()
	}

	return burgerParts
}

func (b *BurgerPartsGenerator) GetGameStatus() bool {
	return b.gameStatus
}

func (b *BurgerPartsGenerator) SetGameStatus(status bool) {
	b.gameStatus = status
}

func (b *BurgerPartsGenerator) GetStatus() bool {
	return b.status
}

func (b *BurgerPartsGenerator) SetStatus(status bool) {
	b.status = status
}

func (b *BurgerPartsGenerator) HideAllImages() {
	for _, item := range b.burgerParts {
		item.Hide()
	}
}

func randPosition() fyne.Position {
	randXPos := 0 + rand.Intn(820-0)
	return fyne.NewPos(float32(randXPos), 0)
}

func (b *BurgerPartsGenerator) CollapseItem(image *canvas.Image) {
	for b.status {
		image.Move(fyne.NewPos(image.Position().X, image.Position().Y+5))
		time.Sleep(time.Millisecond * b.timeToCollapse)
		if image.Position().X-100 < b.dish.img.Position().X && image.Position().X+280 > b.dish.img.Position().X+180 && image.Position().Y == 520-b.dish.itemsCounter {
			b.dish.SetItemOnDish(image)
			break
		}

		if image.Position().Y >= 520 {
			b.HideAllImages()
			b.timeToCollapse = 40
			b.SetGameStatus(false)
			break
		}
	}
}

func (b *BurgerPartsGenerator) ResetItem(image *canvas.Image) {
	image.Move(randPosition())
}
