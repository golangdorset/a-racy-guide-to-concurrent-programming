package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	// A list of urls we want to get stats on.
	urls = []string{
		"https://www.bbc.co.uk",
		"https://www.google.com",
		"https://www.reddit.com",
		"https://www.twitch.tv",
	}

	// Make a channel of stats to send results to.
	statsChan = make(chan *stats)
)

// stats represents the response statistics of a url.
type stats struct {
	url     string
	code    int
	latency float64
}

// String returns a string representation of the url statistics.
func (u stats) String() string {
	return fmt.Sprintf("%q code: %d latency: %.3fs", u.url, u.code, u.latency)
}

func main() {
	// Make a 'done' channel to indicate we're done working.
	doneChan := make(chan bool)

	// Range over our list of urls in a goroutine. This makes it much easier to
	// determine when we're done getting the urls while still doing it concurrently.
	go func() {
		var wg sync.WaitGroup

		// Range over the urls dispatching a work goroutine for each one.
		for _, u := range urls {
			wg.Add(1)
			go work(u, &wg)
		}

		// Once the waitgroup routines are complete send a signal to the done channel.
		wg.Wait()
		doneChan <- true
	}()

	// Print out the results as they come in.
	// Note that when the done channel receives a message the select block breaks
	// out to the loop marker.
loop:
	for {
		select {
		case s := <-statsChan:
			fmt.Println(s)
		case <-doneChan:
			break loop
		}
	}
}

// work makes a call to getUrl for the given url, assuming that a waitgroup is
// controlling the calling goroutine of work.
func work(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	stat, err := getUrl(url)
	if err != nil {
		log.Println(err)
		return
	}

	statsChan <- stat
}

// getUrl makes a GET request to the given url and returns the statistics.
func getUrl(url string) (*stats, error) {
	now := time.Now()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	return &stats{
		url:     url,
		code:    resp.StatusCode,
		latency: time.Since(now).Seconds(),
	}, nil
}
