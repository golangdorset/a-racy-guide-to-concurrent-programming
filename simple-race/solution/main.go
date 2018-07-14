package main

import (
	"fmt"
	"math/rand"
	"sync"
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

	// When it comes to concurrent programming these two are your best friends!
	var (
		// The mutex is a mutual exclusion lock. It locks access to memory for
		// a given goroutine.
		mu sync.Mutex

		// A WaitGroup waits for a collection of goroutines to finish.
		wg sync.WaitGroup
	)

	// Run the competition 10 times to get an average winner.
	for i := 0; i <= 10; i++ {
		// Tell the waitgroup we have 10 routines to run.
		wg.Add(1)

		go func() {
			// Tell the waitgroup that this routine is done, the defer keyword
			// ensures this runs when the function completes.
			defer wg.Done()

			// Lock prevents any other routines from accessing 'results',
			// 'src', 'contestants' and the underlying 'r' Rand var.
			mu.Lock()

			winner := randomWinner(src, contestants)
			results[winner]++

			// Unlock opens up the 'results', 'src', 'contestants' and the
			// underlying 'r' Rand var to be accessed again.
			mu.Unlock()
		}()
	}

	// Wait blocks until all the routines we told the waitgroup about complete.
	wg.Wait()

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

	fmt.Println(results)
	fmt.Printf("%q is the winner!\n", overallWinner)
}

// randomWinner returns a contestant at a random index in the input slice.
func randomWinner(src rand.Source, input []string) string {
	r := rand.New(src)
	return input[r.Intn(len(input))]
}
