package scenes

import (
	"cook_burgers/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type MainScene struct {
	window fyne.Window
	status bool
}

var conveyorModel *models.Conveyor
var dishModel *models.Dish
var burgerPartsGeneratorModel *models.BurgerPartsGenerator
var pointsCounterScene *PointsCounter

func NewMainScene(window fyne.Window) *MainScene {
	return &MainScene{
		window: window,
		status: false,
	}
}

func (scene *MainScene) Show() {
	conveyorModel = models.NewConveyor()
	dishModel = models.NewDish(scene.window)
	burgerPartsGeneratorModel = models.NewBurgerPartsGenerator(dishModel)
	pointsCounterScene = pointsCounterScene.NewPointsCounter(scene.window, dishModel)

	conveyorImage := conveyorModel.GetImage()
	dishImage := dishModel.GetImage()

	burgerImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/burger.png"))
	burgerImage.Resize(fyne.NewSize(100, 100))
	burgerImage.Move(fyne.NewPos(10, 10))

	burgerParts := burgerPartsGeneratorModel.GetBurgerParts()

	bottomBreadImage := burgerParts[0]
	lettuceImage := burgerParts[1]
	beefImage := burgerParts[2]
	ketchupImage := burgerParts[3]
	topBreadImage := burgerParts[4]

	startGameButton := widget.NewButton("Start Game", scene.StartGame)
	startGameButton.Resize(fyne.NewSize(150, 30))
	startGameButton.Move(fyne.NewPos(425, 10))

	pauseGameButton := widget.NewButton("Pause Game", nil)
	pauseGameButton.Resize(fyne.NewSize(150, 30))
	pauseGameButton.Move(fyne.NewPos(425, 50))

	pointsCounter := container.NewVBox(
		pointsCounterScene.GetCounterLabel(),
	)
	pointsCounter.Move(fyne.NewPos(825, 20))
	pointsCounter.Resize(fyne.NewSize(150, 30))

	scene.window.SetContent(container.NewWithoutLayout(burgerImage, conveyorImage, bottomBreadImage, ketchupImage, lettuceImage, beefImage, topBreadImage, dishImage, startGameButton, pauseGameButton, pointsCounter))
}

func (scene *MainScene) GetWindow() fyne.Window {
	return scene.window
}

func (scene *MainScene) StartGame() {
	if !scene.status {
		go dishModel.Run()
		go burgerPartsGeneratorModel.Run()
		go pointsCounterScene.Run()
	}
	scene.status = true
}
