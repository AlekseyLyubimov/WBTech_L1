package main

func task03() {
	numbers := []int{2, 4, 6, 8, 10}
	ch := recurrent_sum_sqr(numbers)
	println(<-ch)
}

func recurrent_sum_sqr(input []int) <-chan int {
	ch := make(chan int)
	go func() {
		if len(input) == 1 {
			ch <- (input[0] * input[0])
		} else {
			ch1 := recurrent_sum_sqr(input[:len(input)/2])
			ch2 := recurrent_sum_sqr(input[len(input)/2:])
			x, y := <-ch1, <-ch2
			ch <- x + y
		}
	}()

	return ch
}