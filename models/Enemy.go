package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type Enemy struct {
	life                int
	name                string
	idleSprites         [7]*canvas.Image
	spriteSize          fyne.Size
	spritePos           fyne.Position
	spriteAnimationTime time.Duration
}

func NewEnemy(name string) *Enemy {
	return &Enemy{
		life: 100,
		name: name,
		idleSprites: [7]*canvas.Image{
			canvas.NewImageFromFile("assets/sprites/enemy/idle/enemy_idle0.png"),
			canvas.NewImageFromFile("assets/sprites/enemy/idle/enemy_idle1.png"),
			canvas.NewImageFromFile("assets/sprites/enemy/idle/enemy_idle2.png"),
			canvas.NewImageFromFile("assets/sprites/enemy/idle/enemy_idle3.png"),
			canvas.NewImageFromFile("assets/sprites/enemy/idle/enemy_idle4.png"),
			canvas.NewImageFromFile("assets/sprites/enemy/idle/enemy_idle5.png"),
			canvas.NewImageFromFile("assets/sprites/enemy/idle/enemy_idle6.png"),
		},
		spriteSize:          fyne.NewSize(512, 512),
		spritePos:           fyne.NewPos(910, 150),
		spriteAnimationTime: 120 * time.Millisecond,
	}
}

func (e Enemy) Print(sceneContainer *fyne.Container) {
	e.idleSprites[0].Resize(e.spriteSize)
	e.idleSprites[0].Move(e.spritePos)

	e.idleSprites[1].Resize(e.spriteSize)
	e.idleSprites[1].Move(e.spritePos)

	e.idleSprites[2].Resize(e.spriteSize)
	e.idleSprites[2].Move(e.spritePos)

	e.idleSprites[3].Resize(e.spriteSize)
	e.idleSprites[3].Move(e.spritePos)

	e.idleSprites[4].Resize(e.spriteSize)
	e.idleSprites[4].Move(e.spritePos)

	e.idleSprites[5].Resize(e.spriteSize)
	e.idleSprites[5].Move(e.spritePos)

	e.idleSprites[6].Resize(e.spriteSize)
	e.idleSprites[6].Move(e.spritePos)

	enemyOnScreen := e.idleSprites[0]

	go func() {
		number := 0
		for {
			time.Sleep(e.spriteAnimationTime)
			sceneContainer.Remove(enemyOnScreen)
			enemyOnScreen = e.idleSprites[number]
			sceneContainer.Add(enemyOnScreen)
			sceneContainer.Refresh()
			number++
			if number == 7 {
				number = 0
			}
		}
	}()

	sceneContainer.Add(enemyOnScreen)
}

func (e Enemy) GetLife() int {
	return e.life
}
