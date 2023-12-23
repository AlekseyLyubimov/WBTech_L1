package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := counter{}
	wg1 := &sync.WaitGroup{}

	c2 := counter{}
	wg2 := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		go counter_clicker(&c1, 100, wg1)
	}

	for i := 0; i < 100; i++ {
		go counter_clicker(&c2, 100, wg2)
	}

	wg1.Wait()
	wg2.Wait()

	fmt.Printf("Counter 1: %d\n", c1.Value)
	fmt.Printf("Counter 2: %d\n", c2.Value)
}

func counter_clicker(ctr *counter, count int, wg *sync.WaitGroup) {
	wg.Add(1)

	for i := 0; i < count; i++ {
		ctr.Increment()
	}

	wg.Done()
}

type counter struct {
	Value int32
	Lock sync.Mutex
}

func (c *counter) Increment() {
	c.Lock.Lock()
	c.Value++
	c.Lock.Unlock()
}