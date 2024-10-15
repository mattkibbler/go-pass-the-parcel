# go-pass-the-parcel

A little program written to help get my head around goroutines, channels and context in Go.

Simulates a game of "pass the parcel" where each player passes a "parcel" around until the music, i,e. a timer, stops and the player holding the parcel wins.

Each player is a goroutine, channels are used to pass the parcel and a context with a timeout is used to end the game.

## Usage

Nothing complicated here. Simply call `go run main.go` to kick things off and relive your childhood disappointment.
