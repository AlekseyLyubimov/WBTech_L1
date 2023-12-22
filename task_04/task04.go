package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	worker_count := flag.Int("workers", 0, "")
	flag.Parse()

	channel := make(chan string)
	for i := 0; i < *worker_count; i++ {
		go read_worker(i, channel)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go write_worker(channel, wg)
	go interrupter()
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

func interrupter() {
	signalChannel := make(chan os.Signal, 2)
    signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT)
    for sig := range signalChannel {
		switch sig {
		case os.Interrupt:
			os.Exit(0)
		case syscall.SIGINT:
			os.Exit(0)
		}
	}
}