package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	winTitle          = "Go Snake"
	boardColor        = 0xffCC98
	snakeColor        = 0xAE2D68
	foodColor         = 0x4B4B4B
	winWidth    int32 = 1400
	winHeight   int32 = 800
	maxWidth    int32 = 1080
	maxHeight   int32 = 600
	boardWidth  int32 = 1080
	boardHeight int32 = 600
	cellWidth   int32 = 20
	cellHeight  int32 = 20
)

func run() error {
	var w, h int

	flag.IntVar(&w, "w", 1080, "width of the game board")
	flag.IntVar(&h, "h", 600, "height of the game board")
	flag.Parse()

	boardWidth, boardHeight := int32(w), int32(h)
	if boardWidth == 0 || boardHeight == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := sdl.Init(sdl.INIT_EVERYTHING)
	handleErr("could not initialize SDL:", err)
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		winTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		winWidth,
		winHeight,
		sdl.WINDOW_SHOWN)
	handleErr("failed to create window:", err)
	defer window.Destroy()

	surface, err := window.GetSurface()
	handleErr("could not get surface:", err)

	// Init game board
	board := newBoard(surface)
	var startingPosX, startingPosY int32
	s := newSnake(startingPosX, startingPosY)
	s.paintBody(board, surface)

	f := newFood(startingPosX, startingPosY+1)
	f.paintBody(board, surface)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false

			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN && t.State == sdl.PRESSED {
					switch t.Keysym.Sym {
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
				}
			}
		}

		time.Sleep(100 * time.Millisecond)
		s.move(board, surface)

		if s.body[0].x == f.x && s.body[0].y == f.y {
			f.respawn(board, surface)
		}

		window.UpdateSurface()
	}

	return nil
}

func handleErr(msg string, err error) error {
	if err != nil {
		return fmt.Errorf(msg+" %v", err)
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}
