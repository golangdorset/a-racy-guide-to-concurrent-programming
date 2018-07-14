package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// A list of urls we want to get stats on.
var urls = []string{
	"https://www.bbc.co.uk",
	"https://www.google.com",
	"https://www.reddit.com",
	"https://www.twitch.tv",
}

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
	// Make a channel of stats to send results to.
	statsChan := make(chan *stats)

	// Loop over the urls and hit each one concurrently.
	for _, u := range urls {
		go func() {
			stat, err := getUrl(u)
			if err != nil {
				log.Fatal(err)
			}

			// Send the result to the stats channel.
			statsChan <- stat
		}()
	}

	// Print out the results as they come in.
	for {
		select {
		case s := <-statsChan:
			fmt.Println(s)
		}
	}
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
