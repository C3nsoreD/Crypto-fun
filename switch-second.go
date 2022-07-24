package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Started.")
	// go spinner(100 * time.Millisecond)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan bool)
	// spinner
	go func() {
		for {
			for _, i := range `-\|/` {
				fmt.Printf("\r%c", i)
				time.Sleep(45 * time.Millisecond)
			}
		}
	}()

	check(done, *ticker)
}

// func body defined in main()
func spinner(delay time.Duration) {}

func check(done chan bool, ticker time.Ticker) {
	for {
		select {
		case <-done:
			fmt.Println("Done")
			// os.Exit(1)
		case t := <-ticker.C:
			fmt.Println("Waiting... ", t)
		}
	}
}
