package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainMenu struct {
	window fyne.Window
	app    fyne.App
}

func NewMainMenu(window fyne.Window, app fyne.App) *MainMenu {
	return &MainMenu{window: window, app: app}
}

func (m *MainMenu) PrincipalMenu() {
	titleLabel := widget.NewLabel("JUEGO CHINGON")
	titleLabel.Move(fyne.NewPos(680, 50))

	start := widget.NewButton("Iniciar juego chingon", m.startGame)
	start.Resize(fyne.NewSize(200, 50))
	start.Move(fyne.NewPos(650, 100))

	credits := widget.NewButton("Creditos chingones", m.credits)
	credits.Resize(fyne.NewSize(200, 50))
	credits.Move(fyne.NewPos(650, 180))

	exit := widget.NewButton("Salir del juego chingon", m.exitGame)
	exit.Resize(fyne.NewSize(200, 50))
	exit.Move(fyne.NewPos(650, 260))

	enemyImag := canvas.NewImageFromFile("assets/menu/demon.png")
	enemyImag.Resize(fyne.NewSize(512, 512))
	enemyImag.Move(fyne.NewPos(950, 150))

	playerImag := canvas.NewImageFromFile("assets/menu/Azrael.png")
	playerImag.Resize(fyne.NewSize(512, 512))
	playerImag.Move(fyne.NewPos(50, 150))

	m.window.SetContent(container.NewWithoutLayout(start, exit, titleLabel, credits, enemyImag, playerImag))
}

func (m *MainMenu) startGame() {
	firstScene := NewFirstScene(m.window, m.app)
	firstScene.StartScene()
}

func (m *MainMenu) exitGame() {
	m.window.Close()
}

func (m *MainMenu) credits() {
	author := widget.NewLabel("Jorge B. >> CONEJO <<")
	author.Move(fyne.NewPos(670, 50))

	image := canvas.NewImageFromFile("assets/menu/iwi.png")
	image.Resize(fyne.NewSize(380, 520))
	image.Move(fyne.NewPos(550, 100))

	back := widget.NewButton("volver", func() {
		m.PrincipalMenu()
	})
	back.Resize(fyne.NewSize(200, 50))
	back.Move(fyne.NewPos(650, 700))

	m.window.SetContent(container.NewWithoutLayout(author, image, back))
}
