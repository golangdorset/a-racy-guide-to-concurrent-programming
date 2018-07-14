package main

import (
	"fmt"
	"math/rand"
	"time"
)

var contestants = []string{"john", "jane", "bill", "ted"}

func main() {
	// Create a new pseudo-random source of entropy.
	src := rand.NewSource(time.Now().Unix())

	// Build a map to count contestant points.
	results := make(map[string]int, len(contestants))
	for _, c := range contestants {
		results[c] = 0
	}

	// Run the competition 10 times to get an average winner.
	for i := 0; i <= 10; i++ {
		go func() {
			winner := randomWinner(src, contestants)

			results[winner]++
		}()
	}

	// Work out who has the most points.
	var (
		overallWinner string
		lastPoints    int
	)
	for name, points := range results {
		if points >= lastPoints {
			overallWinner = name
			lastPoints = points
		}
	}

	// Wtf? Sometimes the winner is wrong. Sometimes it crashes?!
	fmt.Println(results)
	fmt.Printf("%q is the winner!\n", overallWinner)
}

// randomWinner returns a contestant at a random index in the input slice.
func randomWinner(src rand.Source, input []string) string {
	r := rand.New(src)
	return input[r.Intn(len(input))]
}
