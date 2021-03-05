package main

import (
	"fmt"
	"sync"
	"time"
)

// function to call and wait to end
func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// initialize a waitgroup
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// wait for em all to finish
	wg.Wait()
}
