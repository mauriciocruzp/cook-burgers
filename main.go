package main

import (
	"cook_burgers/scenes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Cocina Hamburguesas")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(800, 600))

	mainMenuScene := scenes.NewMainScene(myWindow)
	mainMenuScene.Show()
	myWindow.ShowAndRun()
}
