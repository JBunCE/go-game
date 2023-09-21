package main

import (
	"fyne.io/fyne/v2"
	"jueguito/views"
)

func main() {
	window := views.NewMainWindow("RPG CHINGON", fyne.Size{Width: 1366, Height: 768})
	window.Start()
}
