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

func NewMainScene(window fyne.Window) *MainScene {
	return &MainScene{
		window: window,
	}
}

func (scene *MainScene) Show() {
	burgerImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/burger.png"))
	burgerImage.Resize(fyne.NewSize(100, 100))
	burgerImage.Move(fyne.NewPos(10, 10))

	conveyorImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/conveyor.png"))
	conveyorImage.Resize(fyne.NewSize(800, 100))
	conveyorImage.Move(fyne.NewPos(-3, 500))

	dishImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/dish.png"))
	dishImage.Resize(fyne.NewSize(180, 30))
	dishImage.Move(fyne.NewPos(-183, 470))

	dishModel = models.NewDish(dishImage, -183, 470)
	conveyorModel = models.NewConveyor(conveyorImage)

	startGameButton := widget.NewButton("Start Game", scene.StartGame)
	startGameButton.Resize(fyne.NewSize(150, 30))
	startGameButton.Move(fyne.NewPos(300, 10))

	pauseGameButton := widget.NewButton("Pause Game", scene.StopGame)
	pauseGameButton.Resize(fyne.NewSize(150, 30))
	pauseGameButton.Move(fyne.NewPos(300, 50))

	scene.window.SetContent(container.NewWithoutLayout(burgerImage, conveyorImage, dishImage, startGameButton, pauseGameButton))
}

func (scene *MainScene) StartGame() {
	go dishModel.Run()
}

func (scene *MainScene) StopGame() {
	dishModel.Stop()
}
