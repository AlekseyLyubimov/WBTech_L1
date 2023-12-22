package main

import "sync"

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(3)
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	numbers := []int{2, 4, 6, 8, 10, 12, 14}

	go array_sender(numbers, ch1, wg)
	go squarer(ch1, ch2, wg)
	go outputter(ch2, wg)

	wg.Wait()
}

func array_sender(numbers []int, ch1 chan int, wg *sync.WaitGroup) {
	for _, i := range numbers {
		ch1 <- i
	}
	close(ch1)
	wg.Done()
}

func squarer(ch1 chan int, ch2 chan int, wg *sync.WaitGroup) {
	for i := range ch1 {
		ch2 <- (i*i)
	}
	close(ch2)
	wg.Done()
}

func outputter(ch2 chan int, wg *sync.WaitGroup) {
	for i := range ch2 {
		println(i)
	}
	wg.Done()
}
