package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	winTitle   = "Go Snake"
	boardColor = 0xffCC98
	snakeColor = 0xAE2D68
	winWidth,
	winHeight,
	maxWidth,
	maxHeight,
	cellWidth,
	cellHeight int32 = 1400, 800, 1080, 600, 20, 20
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
	board := newBoard(surface, boardWidth, boardHeight)
	var startingPosX, startingPosY int32
	s := newSnake(board, surface, startingPosX, startingPosY)
	s.paintBody(board, surface)

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
