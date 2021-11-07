package main

import "github.com/veandco/go-sdl2/sdl"

func newBoard(surface *sdl.Surface, boardWidth, boardHeight int32) [][]*sdl.Rect {
	middlePositionX := int32((winWidth - boardWidth) / 2)
	middlePositionY := int32((winHeight - boardHeight) / 2)

	posX, posY := middlePositionX, middlePositionY

	xCells := int(boardWidth / cellWidth)
	yCells := int(boardHeight / cellHeight)

	board := make([][]*sdl.Rect, yCells)

	for i := range board {
		for j := 0; j < xCells; j++ {
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
