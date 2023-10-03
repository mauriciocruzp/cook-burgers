package main

import "cook_burgers/views"

func main() {
	mainView := views.NewMainView("Cocina Hamburguesas")

	mainView.Run()
}
