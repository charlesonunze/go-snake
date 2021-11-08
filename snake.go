package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type cell struct {
	x int32
	y int32
}

type snake struct {
	speed     float32
	direction string
	body      []*cell
}

func newSnake(startingPosX, startingPosY int32) *snake {
	return &snake{
		speed:     1000,
		direction: "right",
		body: []*cell{
			{
				x: startingPosX,
				y: startingPosY,
			},
		},
	}
}

func (s *snake) paintBody(board [][]*sdl.Rect, surface *sdl.Surface) error {
	for _, cell := range s.body {
		paintCell(board, surface, snakeColor, cell.x, cell.y)
	}

	return nil
}

func (s *snake) move(board [][]*sdl.Rect, surface *sdl.Surface) {
	switch s.direction {
	case "right":
		s.moveRight(board, surface)
	case "down":
		s.moveDown(board, surface)
	case "left":
		s.moveLeft(board, surface)
	case "up":
		s.moveUp(board, surface)
	}
}

func (s *snake) moveRight(board [][]*sdl.Rect, surface *sdl.Surface) {
	tailPosX := s.body[len(s.body)-1].x
	tailPosY := s.body[len(s.body)-1].y

	head := []*cell{
		{
			x: s.body[0].x + 1,
			y: s.body[0].y,
		},
	}

	s.body = append(head, s.body...)
	s.body = s.body[:len(s.body)-1]

	s.paintBody(board, surface)
	paintCell(board, surface, boardColor, tailPosX, tailPosY)
}

func (s *snake) moveDown(board [][]*sdl.Rect, surface *sdl.Surface) {
	tailPosX := s.body[len(s.body)-1].x
	tailPosY := s.body[len(s.body)-1].y

	head := []*cell{
		{
			x: s.body[0].x,
			y: s.body[0].y + 1,
		},
	}

	s.body = append(head, s.body...)
	s.body = s.body[:len(s.body)-1]

	s.paintBody(board, surface)
	paintCell(board, surface, boardColor, tailPosX, tailPosY)
}

func (s *snake) moveUp(board [][]*sdl.Rect, surface *sdl.Surface) {
	tailPosX := s.body[len(s.body)-1].x
	tailPosY := s.body[len(s.body)-1].y

	head := []*cell{
		{
			x: s.body[0].x,
			y: s.body[0].y - 1,
		},
	}

	s.body = append(head, s.body...)
	s.body = s.body[:len(s.body)-1]

	s.paintBody(board, surface)
	paintCell(board, surface, boardColor, tailPosX, tailPosY)
}

func (s *snake) moveLeft(board [][]*sdl.Rect, surface *sdl.Surface) {
	tailPosX := s.body[len(s.body)-1].x
	tailPosY := s.body[len(s.body)-1].y

	head := []*cell{
		{
			x: s.body[0].x - 1,
			y: s.body[0].y,
		},
	}

	s.body = append(head, s.body...)
	s.body = s.body[:len(s.body)-1]

	s.paintBody(board, surface)
	paintCell(board, surface, boardColor, tailPosX, tailPosY)
}
