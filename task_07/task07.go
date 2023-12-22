package main

import (
	"math/rand"
	"sync"
)

func main() {

	m := MapWrapper{
		Mutex: sync.Mutex{},
		Entry: make(map[int]int),
	}
	wg := sync.WaitGroup{}
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(wg *sync.WaitGroup) {
			for i := 0; i < 100; i++ {
				m.SyncWriter(rand.Intn(10), rand.Intn(100))
			}
			wg.Done()
		}(&wg)
	}

	wg.Wait()
}

type MapWrapper struct {
	Mutex sync.Mutex
	Entry map[int]int
}

func (m *MapWrapper) SyncWriter(key int, value int) {
	m.Mutex.Lock()
	m.Entry[key] = value
	m.Mutex.Unlock()
}