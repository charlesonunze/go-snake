package sprite

import "github.com/charlesonunze/go-snake/pkg/board"

type Sprite interface {
	DetectCollision(b *board.Board)
	HandleCollision(b *board.Board)
}
