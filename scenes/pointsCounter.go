package scenes

import (
	"cook_burgers/models"
	"fmt"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type PointsCounter struct {
	counterLabel *widget.Label
	points       int
	window       fyne.Window
	dish         *models.Dish
}

func (pc *PointsCounter) NewPointsCounter(window fyne.Window, dish *models.Dish) *PointsCounter {
	return &PointsCounter{
		counterLabel: widget.NewLabel("Burgers: 0"),
		points:       0,
		window:       window,
		dish:         dish,
	}
}

func (pc *PointsCounter) Run(wg *sync.WaitGroup, quit chan bool) {
	defer wg.Done()

	for {
		select {
		case <-quit:
			return
		default:
			if len(pc.dish.GetItemsOnDish()) == 5 {
				pc.UpdateCounter()
				time.Sleep(time.Second * 4)
			}
		}
	}
}

func (pc *PointsCounter) UpdateCounter() {
	pc.points++

	pc.counterLabel.SetText("Burgers: " + fmt.Sprint(pc.points))

	pc.window.Canvas().Refresh(pc.counterLabel)
}

func (pc *PointsCounter) GetCounterLabel() *widget.Label {
	return pc.counterLabel
}

func (pc *PointsCounter) GetPoints() int {
	return pc.points
}

func (pc *PointsCounter) SetPoints(points int) {
	pc.points = points
	pc.counterLabel.SetText("Burgers: " + fmt.Sprint(pc.points))

	pc.window.Canvas().Refresh(pc.counterLabel)
}
