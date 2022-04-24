# go-snake

The snake game

## System Requirements

Follow this [link](https://github.com/veandco/go-sdl2#requirements) to install tools needed to run this application properly.

## Assumptions

- Only the snake and food exist in the game board
- The food position does not until consumed by the snake.
- The speed of the snake does not change.
- The game should end when the snake goes past the screen edge.

## Improvements

Both the food and snake can implement a "Sprite" interface, that has methods like DetectCollision and HandleCollision.
This allows us add more characters to the game, each of them implementing their own case of what will happen if they collide with other characters in the game board.

## Task duration

Approximately 100hours in actual development.

## Running the application

```go run ./cmd/snake.go -w 1000 -h 800```

where w is the width of the board and h is the height of the board
