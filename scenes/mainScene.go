package scenes

import (
	"cook_burgers/models"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type MainScene struct {
	window     fyne.Window
	imageIndex int
	imageURIs  []string
}

var c *models.Conveyor

func NewMainScene(window fyne.Window) *MainScene {
	imageURIs := []string{
		"./assets/conveyor-images/14.png",
		"./assets/conveyor-images/15.png",
		"./assets/conveyor-images/16.png",
		"./assets/conveyor-images/17.png",
		"./assets/conveyor-images/18.png",
		"./assets/conveyor-images/19.png",
		"./assets/conveyor-images/20.png",
		"./assets/conveyor-images/21.png",
		"./assets/conveyor-images/23.png",
		"./assets/conveyor-images/24.png",
		"./assets/conveyor-images/25.png",
		"./assets/conveyor-images/26.png",
		"./assets/conveyor-images/27.png",
	}
	return &MainScene{
		window:     window,
		imageIndex: 0,
		imageURIs:  imageURIs,
	}
}

func (scene *MainScene) Show() {
	conveyor := canvas.NewImageFromURI(storage.NewFileURI("./assets/conveyor-images/14.png"))
	conveyor.Resize(fyne.NewSize(980, 420))
	conveyor.Move(fyne.NewPos(-40, 292))

	c = models.NewConveyor(conveyor)

	scene.window.SetContent(container.NewWithoutLayout(conveyor))

	go scene.startImageChangeTimer()
}

func (scene *MainScene) startImageChangeTimer() {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			scene.imageIndex = (scene.imageIndex + 1) % len(scene.imageURIs)
			newURI := storage.NewFileURI(scene.imageURIs[scene.imageIndex])
			newCanvasImage := canvas.NewImageFromURI(newURI)
			newCanvasImage.Resize(fyne.NewSize(980, 420))
			newCanvasImage.Move(fyne.NewPos(-40, 292))

			scene.window.SetContent(container.NewWithoutLayout(newCanvasImage))
		}
	}
}
