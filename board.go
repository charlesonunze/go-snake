package main

import "github.com/veandco/go-sdl2/sdl"

func newBoard(surface *sdl.Surface) [][]*sdl.Rect {
	middlePositionX := int32((winWidth - boardWidth) / 2)
	middlePositionY := int32((winHeight - boardHeight) / 2)

	posX, posY := middlePositionX, middlePositionY

	xCellCount, yCellCount := getCellsCount()

	board := make([][]*sdl.Rect, yCellCount)

	for i := range board {
		for j := 0; j < xCellCount; j++ {
			rect := &sdl.Rect{
				X: posX,
				Y: posY,
				W: cellWidth,
				H: cellHeight,
			}

			surface.FillRect(rect, boardColor)

			board[i] = append(board[i], rect)
			posX += cellWidth
		}

		posX = middlePositionX
		posY += cellHeight
	}

	return board
}
