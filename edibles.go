package main

import (
	"image/color"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Edible struct {
	Point
	Name string
}

type Edibles struct {
	Available []Edible
}

func (es *Edibles) spawn() {
	if len(es.Available) > 0 {
		return
	}

	x := float32(math.Floor(rand.Float64()*(arenaWidth-2) + 1))
	y := float32(math.Floor(rand.Float64()*(arenaHeight-2) + 1))
	e := Edible{Point{x, y}, "APPLE"}

	es.Available = append(es.Available, e)
}

func (es *Edibles) Draw(screen *ebiten.Image) {
	for _, e := range es.Available {
		switch e.Name {
		case "APPLE":
			{
				vector.DrawFilledRect(screen,
					e.x*blockSize+1, e.y*blockSize,
					blockSize-2, blockSize,
					color.RGBA{255, 0, 0, 255}, true)
				vector.DrawFilledRect(screen,
					e.x*blockSize, e.y*blockSize+1,
					blockSize, blockSize-2,
					color.RGBA{255, 0, 0, 255}, true)

			}
		}
	}
}

func (es *Edibles) ClearAt(p Point) {
	for i, e := range es.Available {
		if e.CollidesWith(p) {
			es.Available = append((es.Available)[:i], (es.Available)[i+1:]...)
			return
		}
	}
}

func (es Edibles) CollidesWith(p Point) bool {
	for _, e := range es.Available {
		if e.CollidesWith(p) {
			return true
		}
	}
	return false
}

func (es Edibles) CollidesWithMulti(points []Point) bool {
	for _, p := range points {
		if es.CollidesWith(p) {
			return true
		}
	}
	return false
}
