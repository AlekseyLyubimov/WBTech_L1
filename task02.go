package main

import (
	"fmt"
	"math/rand"
	"time"
)

func task02() {
	println("")
	println("Task 2:")
	numbers := [...]int{2, 4, 6, 8, 10}
	for _, i := range numbers {
		go sqrt(i)
	}

	sleep_time, _ := time.ParseDuration("200ms")
	time.Sleep(sleep_time)
}

func sqrt(i int) {
	sleep_time, _ := time.ParseDuration(fmt.Sprintf("%dms", (rand.Intn(100))))
	time.Sleep(sleep_time)
	println(i * i)
}
