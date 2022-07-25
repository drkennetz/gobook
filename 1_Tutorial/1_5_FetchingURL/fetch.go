package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// performs a concurrent fetch
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel
	}
	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())
}

// Fetch returns the content found at a specified url
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf(err.Error()) // send to channel
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
