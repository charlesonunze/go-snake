package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/charlesonunze/go-snake/board"
	"github.com/charlesonunze/go-snake/food"
	"github.com/charlesonunze/go-snake/snake"
	"github.com/charlesonunze/go-snake/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	winTitle        = "Go Snake"
	winWidth  int32 = 1400
	winHeight int32 = 800
	// @TODO handle max w and h
	maxWidth  int32 = 1080
	maxHeight int32 = 600
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
	board := board.New(surface, boardWidth, boardHeight, winWidth, winHeight)

	startingPosX, startingPosY := utils.GetStartingPosition(*board)
	s := snake.New(surface, startingPosX, startingPosY)
	s.Draw(board)

	startingPosX, startingPosY = utils.GetStartingPosition(*board)
	f := food.New(surface, startingPosX, startingPosY)
	f.Draw(board)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false

			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN && t.State == sdl.PRESSED {
					s.HandleDirection(t.Keysym.Sym)
				}
			}
		}

		time.Sleep(100 * time.Millisecond)

		s.Move(board)
		s.EatFood(board, f)

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
