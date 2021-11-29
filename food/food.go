package food

import (
	"math/rand"
	"time"

	"github.com/charlesonunze/go-snake/board"
	"github.com/charlesonunze/go-snake/common"
	"github.com/charlesonunze/go-snake/utils"
	"github.com/veandco/go-sdl2/sdl"
)

// Food defines the properties of a food
type Food struct {
	X       int32
	Y       int32
	color   uint32
	surface *sdl.Surface
}

// New creates a pointer to the food struct
func New(surface *sdl.Surface, startingPosX, startingPosY int32) *Food {
	rand.Seed(time.Now().UTC().UnixNano())

	return &Food{
		X:       startingPosX,
		Y:       startingPosY,
		surface: surface,
		color:   0x4B4B4B,
	}
}

// Draw handles the painting of the food on the board canvas
func (f *Food) Draw(b *board.Board) {
	utils.PaintCell(b, f.color, f.X, f.Y)
}

// Respawn handles the regeneration of the food when it's been consumed by a snake
func (f *Food) Respawn(b *board.Board, snakeBody []*common.Cell) {
	foodPosX := f.X
	foodPosY := f.Y
	utils.PaintCell(b, b.Color, foodPosX, foodPosY)

	xCellCount, yCellCount := utils.GetCellsCount(*b)

	f.X = int32(rand.Intn(xCellCount))
	f.Y = int32(rand.Intn(yCellCount))

	for _, cell := range snakeBody {
		if f.X == cell.X && f.Y == cell.Y {
			// what happens if it repspawns again on the snake body you ask?
			// well, stop asking and mind your business
			f.X = int32(rand.Intn(xCellCount))
			f.Y = int32(rand.Intn(yCellCount))
		}
	}

	utils.PaintCell(b, f.color, f.X, f.Y)
}
