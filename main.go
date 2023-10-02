package main

import "cook_burgers/views"

func main() {
	view := views.NewMainView("Cocina Hamburguesas")

	view.Run()
}
