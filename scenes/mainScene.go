package scenes

import (
	"cook_burgers/models"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type MainScene struct {
	window fyne.Window
	status bool
	quit   chan bool
	wg     sync.WaitGroup
}

var conveyorModel *models.Conveyor
var dishModel *models.Dish
var burgerPartsGeneratorModel *models.BurgerPartsGenerator
var pointsCounterScene *PointsCounter
var gameOverLabel *canvas.Image

func NewMainScene(window fyne.Window) *MainScene {
	return &MainScene{
		window: window,
		status: false,
		quit:   make(chan bool),
		wg:     sync.WaitGroup{},
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
	burgerImage.Resize(fyne.NewSize(100, 90))
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

	restartGameButton := widget.NewButton("Restart Game", scene.RestartGame)
	restartGameButton.Resize(fyne.NewSize(150, 30))
	restartGameButton.Move(fyne.NewPos(425, 50))

	pointsCounter := container.NewVBox(
		pointsCounterScene.GetCounterLabel(),
	)
	pointsCounter.Resize(fyne.NewSize(150, 30))
	pointsCounter.Move(fyne.NewPos(825, 20))

	gameOverLabel = canvas.NewImageFromURI(storage.NewFileURI("./assets/game_over.png"))
	gameOverLabel.Resize(fyne.NewSize(300, 300))
	gameOverLabel.Move(fyne.NewPos(325, 140))
	gameOverLabel.Hide()

	scene.window.SetContent(container.NewWithoutLayout(burgerImage, conveyorImage, bottomBreadImage, ketchupImage, lettuceImage, beefImage, topBreadImage, dishImage, startGameButton, restartGameButton, pointsCounter, gameOverLabel))
}

func (scene *MainScene) GetStatus() bool {
	return scene.status
}

func (scene *MainScene) SetStatus(status bool) {
	scene.status = status
}

func (scene *MainScene) StartGame() {
	if !scene.GetStatus() {
		scene.quit = make(chan bool)
		burgerPartsGeneratorModel.SetGameStatus(true)

		go dishModel.Run()
		scene.wg.Add(1)
		go burgerPartsGeneratorModel.Run(&scene.wg, scene.quit)
		go pointsCounterScene.Run()
		go scene.GameOver()
	}
	scene.SetStatus(true)
}

func (scene *MainScene) GameOver() {
	for scene.GetStatus() {
		if !burgerPartsGeneratorModel.GetGameStatus() {
			dishModel.SetStatus(false)
			burgerPartsGeneratorModel.Stop()
			gameOverLabel.Show()
			close(scene.quit)
			scene.wg.Wait()
			println("Game Over")
			time.Sleep(time.Second * 2)
			gameOverLabel.Hide()
			pointsCounterScene.SetPoints(0)
			scene.SetStatus(false)
		}
	}
}

func (scene *MainScene) RestartGame() {
	if scene.GetStatus() {
		dishModel.SetStatus(false)
		burgerPartsGeneratorModel.Stop()
		close(scene.quit)
		scene.wg.Wait()
		pointsCounterScene.SetPoints(0)
	}
	scene.SetStatus(false)
	scene.StartGame()
}
