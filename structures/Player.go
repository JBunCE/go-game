package structures

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type IPlayer interface {
	PrintName()
	SetLife(life int)
	GetLife() int
	Print(container *fyne.Container)
}

type player struct {
	name       string
	Life       int
	sprites    [7]*canvas.Image
	spriteSize fyne.Size
}

func NewPlayer(name string) IPlayer {
	return player{
		Life: 100,
		name: name,
		sprites: [7]*canvas.Image{
			canvas.NewImageFromFile("src/sprites/player/main/sprite_idle0.png"),
			canvas.NewImageFromFile("src/sprites/player/main/sprite_idle1.png"),
			canvas.NewImageFromFile("src/sprites/player/main/sprite_idle2.png"),
			canvas.NewImageFromFile("src/sprites/player/main/sprite_idle3.png"),
			canvas.NewImageFromFile("src/sprites/player/main/sprite_idle4.png"),
			canvas.NewImageFromFile("src/sprites/player/main/sprite_idle5.png"),
			canvas.NewImageFromFile("src/sprites/player/main/sprite_idle6.png"),
		},
		spriteSize: fyne.NewSize(512, 512),
	}
}

func (p player) Print(sceneContainer *fyne.Container) {
	p.sprites[0].Resize(p.spriteSize)
	p.sprites[0].Move(fyne.NewPos(50, 150))

	p.sprites[1].Resize(p.spriteSize)
	p.sprites[1].Move(fyne.NewPos(50, 150))

	p.sprites[2].Resize(p.spriteSize)
	p.sprites[2].Move(fyne.NewPos(50, 150))

	p.sprites[3].Resize(p.spriteSize)
	p.sprites[3].Move(fyne.NewPos(50, 150))

	p.sprites[4].Resize(p.spriteSize)
	p.sprites[4].Move(fyne.NewPos(50, 150))

	p.sprites[5].Resize(p.spriteSize)
	p.sprites[5].Move(fyne.NewPos(50, 150))

	p.sprites[6].Resize(p.spriteSize)
	p.sprites[6].Move(fyne.NewPos(50, 150))

	playerOnScreen := p.sprites[0]

	go func() {
		number := 0
		for {
			time.Sleep(150 * time.Millisecond)
			sceneContainer.Remove(playerOnScreen)
			playerOnScreen = p.sprites[number]
			sceneContainer.Add(playerOnScreen)
			sceneContainer.Refresh()
			number++
			if number == 7 {
				number = 0
			}
		}
	}()

	sceneContainer.Add(playerOnScreen)
}

func (p player) PrintName() {
	fmt.Print(p.name)
}

func (p player) SetLife(life int) {
	p.Life = life
}

func (p player) GetLife() int {
	return p.Life
}
