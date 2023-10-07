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
	dishModel = models.NewDish(scene.window)
	burgerPartsGeneratorModel = models.NewBurgerPartsGenerator(dishModel)

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

	scene.window.SetContent(container.NewWithoutLayout(burgerImage, conveyorImage, bottomBreadImage, ketchupImage, lettuceImage, beefImage, topBreadImage, dishImage, startGameButton))
}

func (scene *MainScene) StartGame() {
	go dishModel.Run()
	go burgerPartsGeneratorModel.Run()
}
