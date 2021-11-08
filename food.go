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

func (f *food) respawn(board [][]*sdl.Rect, surface *sdl.Surface, snake *snake) {
	foodPosX := f.x
	foodPosY := f.y
	paintCell(board, surface, boardColor, foodPosX, foodPosY)

	xCellCount, yCellCount := getCellsCount()

	f.x = int32(rand.Intn(xCellCount))
	f.y = int32(rand.Intn(yCellCount))

	for _, cell := range snake.body {
		if f.x == cell.x && f.y == cell.y {
			f.x = int32(rand.Intn(xCellCount))
			f.y = int32(rand.Intn(yCellCount))
		}
	}

	paintCell(board, surface, foodColor, f.x, f.y)
}
