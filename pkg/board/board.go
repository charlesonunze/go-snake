package board

import (
	"github.com/charlesonunze/go-snake/pkg/common"
	"github.com/veandco/go-sdl2/sdl"
)

// Board defines the properties of a board
type Board struct {
	Width           int32
	Height          int32
	CellWidth       int32
	CellHeight      int32
	Score           int32
	Color           uint32
	Canvas          [][]*sdl.Rect
	Surface         *sdl.Surface
	OccupiedSquares map[string]common.Cell
}

// New creates a pointer to the board struct
func New(surface *sdl.Surface, boardWidth, boardHeight, winWidth, winHeight int32) *Board {
	var (
		cellWidth  int32  = 20
		cellHeight int32  = 20
		boardColor uint32 = 0xffCC98
	)

	middlePositionX := (winWidth - boardWidth) / 2
	middlePositionY := (winHeight - boardHeight) / 2

	posX, posY := middlePositionX, middlePositionY

	xCellCount := int(boardWidth / cellWidth)
	yCellCount := int(boardHeight / cellHeight)

	boardCanvas := make([][]*sdl.Rect, yCellCount)

	for i := range boardCanvas {
		for j := 0; j < xCellCount; j++ {
			rect := &sdl.Rect{
				X: posX,
				Y: posY,
				W: cellWidth,
				H: cellHeight,
			}

			// @TODO - handle error
			surface.FillRect(rect, boardColor)

			boardCanvas[i] = append(boardCanvas[i], rect)
			posX += cellWidth
		}

		posX = middlePositionX
		posY += cellHeight
	}

	return &Board{
		Width:           boardWidth,
		Height:          boardHeight,
		CellWidth:       cellWidth,
		CellHeight:      cellHeight,
		Color:           boardColor,
		Canvas:          boardCanvas,
		Surface:         surface,
		OccupiedSquares: map[string]common.Cell{},
	}
}
