package food

import (
	"math/rand"
	"time"

	"github.com/charlesonunze/go-snake/pkg/board"
	"github.com/charlesonunze/go-snake/pkg/common"
	"github.com/charlesonunze/go-snake/pkg/utils"
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
	key := utils.GetPositionKey(f.X, f.Y)
	b.OccupiedSquares[key] = common.Cell{
		X: f.X,
		Y: f.Y,
	}
	utils.PaintCell(b, f.color, f.X, f.Y)
}

// Respawn handles the regeneration of the food when it's been consumed by a snake
func (f *Food) Respawn(b *board.Board) {
	key := utils.GetPositionKey(f.X, f.Y)

	// free up space
	delete(b.OccupiedSquares, key)
	utils.PaintCell(b, b.Color, f.X, f.Y)

	// get new random position
	xCellCount, yCellCount := utils.GetCellsCount(*b)
	f.X = int32(rand.Intn(xCellCount))
	f.Y = int32(rand.Intn(yCellCount))

	// if the generated position exists in an occupied square, recurse
	key = utils.GetPositionKey(f.X, f.Y)
	_, ok := b.OccupiedSquares[key]
	if ok {
		f.Respawn(b)
	}

	b.OccupiedSquares[key] = common.Cell{
		X: f.X,
		Y: f.Y,
	}
	utils.PaintCell(b, f.color, f.X, f.Y)
}
