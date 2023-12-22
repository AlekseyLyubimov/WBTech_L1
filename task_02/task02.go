package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	println("Task 2:")
	numbers := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	wg.Add(len(numbers))

	for _, i := range numbers {
		go sqrt(i, wg)
	}

	wg.Wait()
}

func sqrt(i int, wg *sync.WaitGroup) {
	sleep_time, _ := time.ParseDuration(fmt.Sprintf("%dms", (rand.Intn(100))))
	time.Sleep(sleep_time)
	println(i * i)
	wg.Done()
}
