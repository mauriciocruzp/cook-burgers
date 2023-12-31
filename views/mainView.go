package views

import (
	"cook_burgers/scenes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type MainView struct{}

func NewMainView() *MainView {
	return &MainView{}
}

func (v *MainView) Run() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Cook Burgers")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(1000, 650))

	mainMenuScene := scenes.NewMainScene(myWindow)
	mainMenuScene.Show()
	myWindow.ShowAndRun()
}
