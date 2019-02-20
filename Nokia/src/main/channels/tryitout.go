package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Duration(id) * time.Second)
		fmt.Println("worker", id, "finished job", j)
	}
}

func main() {
	jobs := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs)
	}

	for j := 1; j <= 12; j++ {
		jobs <- j
	}
	close(jobs)
	fmt.Println("jobs closed")
	time.Sleep(5*time.Second)

}