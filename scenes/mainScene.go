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
}

var conveyorModel *models.Conveyor
var dishModel *models.Dish
var burgerPartsGeneratorModel *models.BurgerPartsGenerator

func NewMainScene(window fyne.Window) *MainScene {
	return &MainScene{
		window: window,
	}
}

func (scene *MainScene) Show() {
	conveyorModel = models.NewConveyor()
	burgerPartsGeneratorModel = models.NewBurgerPartsGenerator()
	dishModel = models.NewDish(burgerPartsGeneratorModel, scene.window)

	conveyorImage := conveyorModel.GetImage()
	dishImage := dishModel.GetImage()

	burgerImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/burger.png"))
	burgerImage.Resize(fyne.NewSize(100, 100))
	burgerImage.Move(fyne.NewPos(10, 10))

	topBreadImage := burgerPartsGeneratorModel.GetTopBread()
	ketchupImage := burgerPartsGeneratorModel.GetKetchup()
	lettuceImage := burgerPartsGeneratorModel.GetLettuce()
	beefImage := burgerPartsGeneratorModel.GetBeef()
	bottomBreadImage := burgerPartsGeneratorModel.GetBottomBread()

	startGameButton := widget.NewButton("Start Game", scene.StartGame)
	startGameButton.Resize(fyne.NewSize(150, 30))
	startGameButton.Move(fyne.NewPos(425, 10))

	pauseGameButton := widget.NewButton("Pause Game", scene.StopGame)
	pauseGameButton.Resize(fyne.NewSize(150, 30))
	pauseGameButton.Move(fyne.NewPos(425, 50))

	scene.window.SetContent(container.NewWithoutLayout(burgerImage, conveyorImage, bottomBreadImage, ketchupImage, lettuceImage, beefImage, topBreadImage, dishImage, startGameButton, pauseGameButton))
}

func (scene *MainScene) StartGame() {
	go dishModel.Run()
	go burgerPartsGeneratorModel.Run()
	go burgerPartsGeneratorModel.MoveItems()
}

func (scene *MainScene) StopGame() {
	dishModel.Stop()
	burgerPartsGeneratorModel.Stop()
}
