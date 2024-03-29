package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/charlesonunze/go-snake/pkg/board"
	"github.com/charlesonunze/go-snake/pkg/food"
	"github.com/charlesonunze/go-snake/pkg/snake"
	"github.com/charlesonunze/go-snake/pkg/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	winTitle        = "Go Snake"
	winWidth  int32 = 1200
	winHeight int32 = 800

	minBoardWidth  int32 = 400
	maxBoardWidth  int32 = 1200
	minBoardHeight int32 = 400
	maxBoardHeight int32 = 1200
)

func run() error {
	var w, h int
	flag.IntVar(&w, "w", 1080, "width of the game board")
	flag.IntVar(&h, "h", 600, "height of the game board")
	flag.Parse()

	boardWidth, boardHeight := int32(w), int32(h)
	if boardWidth < minBoardWidth ||
		boardWidth > maxBoardWidth ||
		boardHeight < minBoardHeight ||
		boardHeight > maxBoardHeight {
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
	b := board.New(surface, boardWidth, boardHeight, winWidth, winHeight)

	startingPosX, startingPosY := utils.GetStartingPosition(*b)
	s := snake.New(surface, startingPosX, startingPosY)
	s.Draw(b)

	startingPosX, startingPosY = utils.GetStartingPosition(*b)
	f := food.New(surface, startingPosX, startingPosY)
	f.Draw(b)

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
		printGameState(*b, *s)

		s.Move(b, &running)
		s.EatFood(b, f)

		window.UpdateSurface()
	}

	fmt.Println("GAME OVER!!!!!!!")

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

func printGameState(b board.Board, s snake.Snake) {
	fmt.Println("")

	fmt.Printf("SCORE: %d \n", b.Score)

	fmt.Printf("SNAKE LENGTH: %d  \n", s.GetSnakeLength())
	fmt.Printf("SNAKE ROUND: %d \n", s.Round)

	head := s.GetHeadPosition()
	fmt.Printf("SNAKE HEAD POSITION: %d, %d \n", head.X, head.Y)
	fmt.Printf("SNAKE HEAD POSITION: %v \n", b.OccupiedSquares)
}
