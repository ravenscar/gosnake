package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Snake struct {
	Body      []Point
	direction Direction
	Color     color.Color
}

func (s Snake) canSetDirection(d Direction) bool {
	if d.h != 0 && s.direction.h != 0 {
		return false
	}
	if d.v != 0 && s.direction.v != 0 {
		return false
	}
	return true
}

func (s *Snake) SoftSetDirection(d Direction) {
	if s.canSetDirection(d) {
		s.direction = d
	}
}

func (s Snake) Draw(screen *ebiten.Image) {
	for _, p := range s.Body {
		vector.DrawFilledRect(screen,
			p.x*blockSize+1, p.y*blockSize,
			blockSize-2, blockSize,
			s.Color, true)
		vector.DrawFilledRect(screen,
			p.x*blockSize, p.y*blockSize+1,
			blockSize, blockSize-2,
			s.Color, true)
	}
	headColor := color.RGBA{0, 0, 0, 63}

	x := s.Body[0].x
	y := s.Body[0].y

	vector.DrawFilledRect(screen,
		x*blockSize+2, y*blockSize+2,
		blockSize-4, blockSize-4,
		headColor, true)

	var eye1 Point
	var eye2 Point

	if s.direction.isSameDirection(Up()) {
		eye1 = Point{x*blockSize + 2, y*blockSize + 2}
		eye2 = Point{x*blockSize + blockSize - 3, y*blockSize + 2}
	}

	if s.direction.isSameDirection(left()) {
		eye1 = Point{x*blockSize + 2, y*blockSize + 2}
		eye2 = Point{x*blockSize + 2, y*blockSize + blockSize - 3}
	}

	if s.direction.isSameDirection(Down()) {
		eye1 = Point{x*blockSize + 2, y*blockSize + blockSize - 3}
		eye2 = Point{x*blockSize + blockSize - 3, y*blockSize + blockSize - 3}
	}

	if s.direction.isSameDirection(Right()) {
		eye1 = Point{x*blockSize + blockSize - 3, y*blockSize + 2}
		eye2 = Point{x*blockSize + blockSize - 3, y*blockSize + blockSize - 3}
	}

	vector.DrawFilledRect(screen, eye1.x, eye1.y, 1, 1, color.Black, true)
	vector.DrawFilledRect(screen, eye2.x, eye2.y, 1, 1, color.Black, true)
}

func NewSnake() Snake {
	dir := RandomDirection()
	// start off central and build up from random dir above
	x := float32(math.Floor(float64(arenaWidth / 2)))
	y := float32(math.Floor(float64(arenaHeight / 2)))
	body := []Point{{x, y}}
	for i := 0; i < 3; i++ {
		next := body[0].Translate(dir)
		body = append([]Point{next}, body...)
	}
	return Snake{
		Body:      body,
		direction: dir,
		Color:     color.RGBA{194, 249, 112, 255},
	}
}

func (s Snake) NextPoint() Point {
	return s.Body[0].Translate(s.direction)
}

func (s Snake) CollidesWith(p Point) bool {
	return s.CollidesWithMulti([]Point{p})
}

func (s Snake) CollidesWithMulti(points []Point) bool {
	for _, p := range s.Body {
		if p.CollidesWithMulti(points) {
			return true
		}
	}
	return false
}

func (s *Snake) move(foodTester CollisionChecker, wallTester CollisionChecker) {
	next := s.NextPoint()

	if wallTester.CollidesWith(next) {
		// we collide with a wall so we need to get a new direction
		head := s.Body[0]
		dir := s.direction
		for wallTester.CollidesWith(head.Translate(dir)) {
			newDir := RandomDirection()
			// we can't just set it here as it could collide with another wall
			if s.canSetDirection(newDir) {
				dir = newDir
			}
		}
		s.SoftSetDirection(dir)
	}

	next = s.NextPoint()

	if wallTester.CollidesWith(next) {
		panic("could not avoid collision")
	}

	if foodTester.CollidesWith(next) {
		s.Body = append([]Point{next}, s.Body[:len(s.Body)]...)
	} else {
		s.Body = append([]Point{next}, s.Body[:len(s.Body)-1]...)
	}
}
