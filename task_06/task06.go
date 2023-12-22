package main

import (
	"context"
	"time"
)

func main() {

	ch1 := make(chan string)
	go for_each_with_channel(ch1)

	ctx, cancel := context.WithCancel(context.Background())
	go cancel_with_context(ctx)

	trigger := true
	go shared_wariable(&trigger)

	ch2 := make(chan bool)
	go channel_signal(ch2)

	time.Sleep(time.Second)

	close(ch1)
	cancel()
	ch2 <- true
	trigger = false
}

func for_each_with_channel(channel chan string) {
	for message := range channel {
		println(message)
	}
	println("for-each loop with channel exits when channel is closed")
}

func cancel_with_context(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
		 return
		default:
		 println("goroutine can be stopped through context")
		 time.Sleep(time.Second)
		}
	}
}

func shared_wariable(trigger *bool) {
	if *trigger {
		println("goroutine can be stopped through shared variable")
		time.Sleep(time.Second)
	}
}

func channel_signal(channel chan bool) {
	for {
		select {
		case <-channel:
			println("goroutine can be stopped as explicit reaction to channel signal")
		default:
			time.Sleep(time.Second)
		}
	}
}