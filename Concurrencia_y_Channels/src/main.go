package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	// Simulate work
	for i := 0; i < 5; i++ {
		fmt.Printf("Worker %d working on task %d\n", id, i)
	}
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")
}
