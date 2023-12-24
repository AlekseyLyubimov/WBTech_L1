package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	sleep(time.Second * 3)
	fmt.Println(time.Now())
}

func sleep(value time.Duration) {
	<-time.After(value)
}