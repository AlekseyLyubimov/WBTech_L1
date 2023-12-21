package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	task01()
	task02()
}

func pause() {
	//костыль для рандомной задержки вывода
	sleep_time, _ := time.ParseDuration(fmt.Sprintf("%dms", (rand.Intn(100))))
	time.Sleep(sleep_time)
}