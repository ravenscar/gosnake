package main

import (
	"image/color"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	blockSize   = 8.0
	arenaWidth  = 80.0
	arenaHeight = 45.0
	sWidth      = blockSize * arenaWidth
	sHeight     = blockSize * arenaHeight
	speedMillis = 100
)

type CollisionChecker interface {
	CollidesWith(p Point) bool
	CollidesWithMulti(p []Point) bool
}

type Walls struct{}

type Point struct {
	x float32
	y float32
}

type Direction struct {
	h float32
	v float32
}

func Up() Direction {
	return Direction{0, -1}
}

func Down() Direction {
	return Direction{0, 1}
}

func left() Direction {
	return Direction{-1, 0}
}

func Right() Direction {
	return Direction{1, 0}
}

func RandomDirection() Direction {
	r := rand.IntN(4)
	switch r {
	case 0:
		return Up()
	case 1:
		return left()
	case 2:
		return Down()
	case 3:
		return Right()
	}
	panic("Could not get random direction")
}

func (d1 Direction) isSameDirection(d2 Direction) bool {
	return d1.h == d2.h && d1.v == d2.v
}

func (p1 Point) CollidesWith(p2 Point) bool {
	return math.Floor(float64(p1.x-p2.x)) == 0 && math.Floor(float64(p1.y-p2.y)) == 0
}

func (p1 Point) CollidesWithMulti(points []Point) bool {
	for _, p2 := range points {
		if p1.CollidesWith(p2) {
			return true
		}
	}
	return false
}

func (w Walls) Draw(screen *ebiten.Image) {
	wallColor := color.RGBA{150, 40, 27, 255}
	bgColor := color.RGBA{131, 101, 57, 255}

	vector.DrawFilledRect(screen,
		0, 0,
		sWidth, sHeight,
		wallColor, true)
	vector.DrawFilledRect(screen,
		blockSize, blockSize,
		sWidth-blockSize*2, sHeight-blockSize*2,
		bgColor, true)
}

func (w Walls) CollidesWith(p Point) bool {
	if math.Round(float64(p.y)) <= 0 {
		return true
	}
	if math.Round(float64(p.x)) <= 0 {
		return true
	}
	if math.Round(float64(p.y)) >= arenaHeight-1 {
		return true
	}
	if math.Round(float64(p.x)) >= arenaWidth-1 {
		return true
	}
	return false
}

func (w Walls) CollidesWithMulti(points []Point) bool {
	for _, p := range points {
		if w.CollidesWith(p) {
			return true
		}
	}
	return false
}

func (p Point) Translate(d Direction) Point {
	return Point{p.x + d.h, p.y + d.v}
}
