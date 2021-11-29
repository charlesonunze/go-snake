package utils

import (
	"math/rand"
	"time"

	"github.com/charlesonunze/go-snake/board"
	"github.com/veandco/go-sdl2/sdl"
)

// PaintCell handles the painting of a pixel on the board
func PaintCell(b *board.Board, color uint32, posX, posY int32) {
	rect := &sdl.Rect{
		X: b.Canvas[posY][posX].X,
		Y: b.Canvas[posY][posX].Y,
		W: b.CellHeight,
		H: b.CellHeight,
	}

	b.Surface.FillRect(rect, color)
}

// GetCellsCount returns the number of pixels in the X and Y axis of the board
func GetCellsCount(b board.Board) (int, int) {
	xCellCount := int(b.Width / b.CellWidth)
	yCellCount := int(b.Height / b.CellHeight)

	return xCellCount, yCellCount
}

// GetStartingPosition returns a random position on the board
func GetStartingPosition(b board.Board) (int32, int32) {
	rand.Seed(time.Now().UTC().UnixNano())
	xCellCount, yCellCount := GetCellsCount(b)

	x := int32(rand.Intn(xCellCount))
	y := int32(rand.Intn(yCellCount))

	return x, y
}
