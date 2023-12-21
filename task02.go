package main

import (
	"fmt"
	"math/rand"
	"time"
)

func task02() {
	//task 2:
	println("")
	numbers := [...]int{2, 4, 6, 8, 10}
	for _, i := range numbers {
		go sqrt(i)
	}

	sleep_time, _ := time.ParseDuration("250ms")
	time.Sleep(sleep_time)
}

func sqrt(i int) {
	sleep_time, _ := time.ParseDuration(fmt.Sprintf("%dms", (rand.Intn(100))))
	time.Sleep(sleep_time)
	println(i * i)
}
