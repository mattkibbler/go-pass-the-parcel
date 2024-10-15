package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func player(ctx context.Context, id int, ch chan int, next chan int) {
	// Infinite loop to always wait for the parcel until the ctx is cancelled
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Player %d: Music stopped\n", id)
		case parcel := <-ch:
			fmt.Printf("Player %d: Got the parcel!\n", id)
			// Pause, holding on to the parcel
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			// Pass the parcel
			next <- parcel
		}
	}
}

func main() {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	numPlayers := 5
	channels := make([]chan int, numPlayers)
	for i := 0; i < numPlayers; i++ {
		channels[i] = make(chan int)
	}

	for i := 0; i < numPlayers; i++ {
		next := channels[(i+1)%numPlayers]
		go player(ctx, i+1, channels[i], next)
	}

	// Start the game by giving the parcel to the first player
	fmt.Println("Starting the game...")
	channels[0] <- 1

	<-ctx.Done()
	fmt.Println("Game over...")
}
