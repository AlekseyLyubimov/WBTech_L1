package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

//cmd := flag.String("workers", "", "")

func main() {
	worker_count := flag.Int("workers", 0, "")
	worktime := flag.Duration("worktime", time.Duration(0), "")
	flag.Parse()

	channel := make(chan string)
	for i := 0; i < *worker_count; i++ {
		go read_worker(i, channel)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go write_worker(channel, wg)
	go interrupter(*worktime)
	wg.Wait()
}

func write_worker(channel chan string, wg *sync.WaitGroup) {
	sleep_time, _ := time.ParseDuration("250ms")
	
	for {
		//bulding random 10-character string
		rand.Seed(time.Now().UnixNano())
		b := make([]byte, 12)
		rand.Read(b)
		
		channel <- fmt.Sprintf("%x", b)[2 : 12]
		time.Sleep(sleep_time)
	}
}

func read_worker(readerID int, channel chan string) {
	for msg := range channel {
        println(fmt.Printf("Reader №%d got message %s\n", readerID, msg))
    }
}

func interrupter(worktime time.Duration) {
	timer := time.NewTimer(worktime)
	<- timer.C
	os.Exit(0)
}