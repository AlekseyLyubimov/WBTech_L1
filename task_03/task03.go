package main

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	main_ch := make(chan int)
	go recurrent_sum_sqr(numbers, main_ch)

	x := <-main_ch
	println(x)
}

func recurrent_sum_sqr(input []int, channel chan int) {
	if len(input) == 1 {
		channel <- (input[0] * input[0])
	} else {
		ch1 := make(chan int)
		ch2 := make(chan int)

		go recurrent_sum_sqr(input[:len(input)/2], ch1)
		go recurrent_sum_sqr(input[len(input)/2:], ch2)

		x, y := <-ch1, <-ch2
		channel <- x + y
	}
}

/*

alternative implementation

func main() {
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
*/