package snake

import (
	"fmt"

	"github.com/charlesonunze/go-snake/board"
	"github.com/charlesonunze/go-snake/common"
	"github.com/charlesonunze/go-snake/food"
	"github.com/charlesonunze/go-snake/utils"
	"github.com/veandco/go-sdl2/sdl"
)

// Snake defines the properties of a snake
type Snake struct {
	// speed     float32
	direction string
	body      []*common.Cell
	color     uint32
	surface   *sdl.Surface
}

// New creates a pointer to the snake struct
func New(surface *sdl.Surface, startingPosX, startingPosY int32) *Snake {
	return &Snake{
		// speed:     1000,
		surface:   surface,
		color:     0xAE2D68,
		direction: "right",
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
func (s *Snake) Move(b *board.Board) {
	switch s.direction {
	case "right":
		s.moveRight(b)
	case "down":
		s.moveDown(b)
	case "left":
		s.moveLeft(b)
	case "up":
		s.moveUp(b)
	}
}

func (s *Snake) moveRight(b *board.Board) {
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

	s.Draw(b)

	key := utils.GetPositionKey(tailPosX, tailPosY)
	delete(b.OccupiedSquares, key)
	utils.PaintCell(b, b.Color, tailPosX, tailPosY)
}

func (s *Snake) moveDown(b *board.Board) {
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

	s.Draw(b)

	key := utils.GetPositionKey(tailPosX, tailPosY)
	delete(b.OccupiedSquares, key)
	utils.PaintCell(b, b.Color, tailPosX, tailPosY)
}

func (s *Snake) moveUp(b *board.Board) {
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

	s.Draw(b)

	key := utils.GetPositionKey(tailPosX, tailPosY)
	delete(b.OccupiedSquares, key)
	utils.PaintCell(b, b.Color, tailPosX, tailPosY)
}

func (s *Snake) moveLeft(b *board.Board) {
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

	s.Draw(b)

	key := utils.GetPositionKey(tailPosX, tailPosY)
	delete(b.OccupiedSquares, key)
	utils.PaintCell(b, b.Color, tailPosX, tailPosY)
}
