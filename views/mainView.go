package views

import (
	"cook_burgers/scenes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type MainView struct {
	title string
}

func NewMainView(title string) *MainView {
	return &MainView{
		title: title,
	}
}

func (v *MainView) Run() {
	myApp := app.New()
	myWindow := myApp.NewWindow(v.title)
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(1000, 650))

	mainMenuScene := scenes.NewMainScene(myWindow)
	mainMenuScene.Show()
	myWindow.ShowAndRun()
}
