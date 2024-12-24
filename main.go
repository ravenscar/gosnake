package main

import (
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player       Snake
	Food         Edibles
	LastMoveTime int
	Walls        Walls
}

func (g *Game) move() {
	now := int(time.Now().UnixMilli())
	if now < g.LastMoveTime+speedMillis {
		return
	}
	g.LastMoveTime = now
	g.Player.move(g.Food, g.Walls)
	g.Food.ClearAt(g.Player.Body[0])
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Player.SoftSetDirection(Up())
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Player.SoftSetDirection(left())
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Player.SoftSetDirection(Down())
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Player.SoftSetDirection(Right())
	}
	g.Food.spawn()
	g.move()
	g.Food.spawn()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Walls.Draw(screen)
	g.Food.Draw(screen)
	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return sWidth, sHeight
}

func main() {
	game := &Game{
		Player:       NewSnake(),
		LastMoveTime: int(time.Now().UnixMilli()),
	}

	ebiten.SetWindowSize(sWidth, sHeight)
	ebiten.SetWindowTitle("Snake")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
