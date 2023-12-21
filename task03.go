package main

import (
	"sync"
)

func task03() {
	numbers := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	main_ch := make(chan int)
	wg.Add(1)
	recurrent_sum_sqr(numbers, main_ch, &wg)

	wg.Wait()
	x := <-main_ch
	println(x)
}

func recurrent_sum_sqr(input []int, channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	if len(input) == 1 {
		channel <- (input[0] * input[0])
	} else {
		ch1 := make(chan int)
		ch2 := make(chan int)

		wg.Add(2)
		recurrent_sum_sqr(input[:len(input)/2], ch1, wg)
		recurrent_sum_sqr(input[len(input)/2:], ch2, wg)

		x, y := <-ch1, <-ch2

		channel <- x + y
	}
	close(channel)
}