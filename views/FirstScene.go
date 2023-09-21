package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"jueguito/structures"
	"strconv"
)

type FirstScene struct {
	window    fyne.Window
	app       fyne.App
	container *fyne.Container
}

func NewFirstScene(window fyne.Window, app fyne.App) *FirstScene {
	return &FirstScene{window: window, app: app, container: container.NewWithoutLayout()}
}

func (f *FirstScene) StartScene() {
	player := structures.NewPlayer("mondongo")
	playerLife := widget.NewLabel("Player life" + strconv.Itoa(player.GetLife()))
	playerLife.Move(fyne.NewPos(250, 100))
	f.container.Add(playerLife)
	player.Print(f.container)

	enemy := structures.NewEnemy("papoi")
	enemy.Print(f.container)
	enemyLife := widget.NewLabel("Enemy life:" + strconv.Itoa(enemy.GetLife()))
	enemyLife.Move(fyne.NewPos(1150, 100))
	f.container.Add(enemyLife)

	f.window.SetContent(f.container)
}
