/*
Channels : similar to JS event emitters or pipes

A typed conduit for sending and receiving values between goroutines.
Declared using: chan Type
Think of it like a pipe — one goroutine sends data in, another receives it.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func chann() {
	ch := make(chan string)

	go func() {
		ch <- "hello from goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)
}

// producer-consumer

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}
/*
The main goroutine sends jobs into a channel.
Multiple worker goroutines process the jobs concurrently.
We use a WaitGroup to wait for all work to finish.
*/
func prodConsumer() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // no more jobs

	// Wait for all workers to finish
	wg.Wait()
	close(results)

	// Collect results
	fmt.Println("\nResults:")
	for res := range results {
		fmt.Println(res)
	}
}
