package main

import (
	"cook_burgers/views"
	"os"
)

func main() {
	os.Setenv("FYNE_THEME", "light")
	mainView := views.NewMainView()

	mainView.Run()
}
