package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	winTitle   = "Go Snake"
	boardColor = 0xffCC98
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
	newBoard(surface, boardWidth, boardHeight)

	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
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
