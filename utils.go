package main

import "github.com/veandco/go-sdl2/sdl"

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
