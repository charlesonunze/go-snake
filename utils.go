package main

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func paintCell(board [][]*sdl.Rect, surface *sdl.Surface, color uint32, posX, posY int32) {
	rect := &sdl.Rect{
		X: board[posY][posX].X,
		Y: board[posY][posX].Y,
		W: cellWidth,
		H: cellHeight,
	}
	surface.FillRect(rect, color)
}

func getCellsCount() (int, int) {
	xCellCount := int(boardWidth / cellWidth)
	yCellCount := int(boardHeight / cellHeight)

	return xCellCount, yCellCount
}

func getStartingPosition() (int32, int32) {
	rand.Seed(time.Now().UTC().UnixNano())
	xCellCount, yCellCount := getCellsCount()

	x := int32(rand.Intn(xCellCount))
	y := int32(rand.Intn(yCellCount))

	return x, y
}
