package main

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type food cell

func newFood(startingPosX, startingPosY int32) food {
	rand.Seed(time.Now().UTC().UnixNano())

	return food{
		x: startingPosX,
		y: startingPosY,
	}
}

func (f *food) paintBody(board [][]*sdl.Rect, surface *sdl.Surface) error {
	paintCell(board, surface, foodColor, f.x, f.y)

	return nil
}
