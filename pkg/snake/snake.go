package snake

import (
	"fmt"

	"github.com/charlesonunze/go-snake/pkg/board"
	"github.com/charlesonunze/go-snake/pkg/common"
	"github.com/charlesonunze/go-snake/pkg/food"
	"github.com/charlesonunze/go-snake/pkg/utils"
	"github.com/veandco/go-sdl2/sdl"
)

// Snake defines the properties of a snake
type Snake struct {
	direction  string
	body       []*common.Cell
	color      uint32
	surface    *sdl.Surface
	cellWidth  int32
	cellHeight int32
	Round      int32
}

// New creates a pointer to the snake struct
func New(surface *sdl.Surface, startingPosX, startingPosY int32) *Snake {
	return &Snake{
		surface:    surface,
		color:      0xAE2D68,
		direction:  "right",
		cellWidth:  20,
		cellHeight: 20,
		body: []*common.Cell{
			{
				X: startingPosX,
				Y: startingPosY,
			},
		},
	}
}

// Draw handles the painting of the snake on the board canvas
func (s *Snake) Draw(b *board.Board) {
	for _, cell := range s.body {
		key := utils.GetPositionKey(cell.X, cell.Y)
		b.OccupiedSquares[key] = common.Cell{
			X: cell.X,
			Y: cell.Y,
		}
		utils.PaintCell(b, s.color, cell.X, cell.Y)
	}
}

// EatFood handles sustenance consumption
func (s *Snake) EatFood(b *board.Board, f *food.Food) {
	if s.body[0].X == f.X && s.body[0].Y == f.Y {
		tailPosX := s.body[len(s.body)-1].X
		tailPosY := s.body[len(s.body)-1].Y
		tail := &common.Cell{
			X: tailPosX,
			Y: tailPosY,
		}
		s.body = append(s.body, tail)

		b.Score++

		f.Respawn(b)
	}
}

// HandleDirection changes the direction of the snake on keypress
func (s *Snake) HandleDirection(keyCode sdl.Keycode) error {
	switch keyCode {
	case sdl.K_UP:
		if s.direction == "down" {
			return fmt.Errorf("cc")
		}
		s.direction = "up"

	case sdl.K_DOWN:
		if s.direction == "up" {
			return fmt.Errorf("cc")
		}
		s.direction = "down"

	case sdl.K_LEFT:
		if s.direction == "right" {
			return fmt.Errorf("cc")
		}
		s.direction = "left"

	case sdl.K_RIGHT:
		if s.direction == "left" {
			return fmt.Errorf("cc")
		}
		s.direction = "right"
	}

	return nil
}

func (s *Snake) GetSnakeLength() int32 {
	return int32(len(s.body)) * s.cellWidth
}

func (s *Snake) GetHeadPosition() *common.Cell {
	return s.body[0]
}

// Move handles the movement of the snake
func (s *Snake) Move(b *board.Board, r *bool) {
	switch s.direction {
	case "right":
		s.moveRight(b, r)
	case "down":
		s.moveDown(b, r)
	case "left":
		s.moveLeft(b, r)
	case "up":
		s.moveUp(b, r)
	}
}

func (s *Snake) moveRight(b *board.Board, r *bool) {
	tailPosX := s.body[len(s.body)-1].X
	tailPosY := s.body[len(s.body)-1].Y

	head := []*common.Cell{
		{
			X: s.body[0].X + 1,
			Y: s.body[0].Y,
		},
	}

	s.body = append(head, s.body...)
	s.body = s.body[:len(s.body)-1]

	screenEdge := b.Width / b.CellWidth

	if s.body[0].X+1 <= screenEdge {
		s.Draw(b)
		s.Round++
	} else {
		*r = false
	}

	key := utils.GetPositionKey(tailPosX, tailPosY)
	delete(b.OccupiedSquares, key)
	utils.PaintCell(b, b.Color, tailPosX, tailPosY)
}

func (s *Snake) moveDown(b *board.Board, r *bool) {
	tailPosX := s.body[len(s.body)-1].X
	tailPosY := s.body[len(s.body)-1].Y

	head := []*common.Cell{
		{
			X: s.body[0].X,
			Y: s.body[0].Y + 1,
		},
	}

	s.body = append(head, s.body...)
	s.body = s.body[:len(s.body)-1]

	screenEdge := b.Height / b.CellHeight

	if s.body[0].Y+1 <= screenEdge {
		s.Draw(b)
		s.Round++
	} else {
		*r = false
	}

	key := utils.GetPositionKey(tailPosX, tailPosY)
	delete(b.OccupiedSquares, key)
	utils.PaintCell(b, b.Color, tailPosX, tailPosY)
}

func (s *Snake) moveUp(b *board.Board, r *bool) {
	tailPosX := s.body[len(s.body)-1].X
	tailPosY := s.body[len(s.body)-1].Y

	head := []*common.Cell{
		{
			X: s.body[0].X,
			Y: s.body[0].Y - 1,
		},
	}

	s.body = append(head, s.body...)
	s.body = s.body[:len(s.body)-1]

	screenEdge := int32(-2)

	if s.body[0].Y-1 == screenEdge {
		*r = false
	} else {
		s.Draw(b)
		s.Round++
	}

	key := utils.GetPositionKey(tailPosX, tailPosY)
	delete(b.OccupiedSquares, key)
	utils.PaintCell(b, b.Color, tailPosX, tailPosY)
}

func (s *Snake) moveLeft(b *board.Board, r *bool) {
	tailPosX := s.body[len(s.body)-1].X
	tailPosY := s.body[len(s.body)-1].Y

	head := []*common.Cell{
		{
			X: s.body[0].X - 1,
			Y: s.body[0].Y,
		},
	}

	s.body = append(head, s.body...)
	s.body = s.body[:len(s.body)-1]

	screenEdge := int32(-2)

	if s.body[0].X-1 == screenEdge {
		*r = false
	} else {
		s.Draw(b)
		s.Round++
	}

	key := utils.GetPositionKey(tailPosX, tailPosY)
	delete(b.OccupiedSquares, key)
	utils.PaintCell(b, b.Color, tailPosX, tailPosY)
}
