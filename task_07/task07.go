package main

import (
	"fmt"
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
		go worker(i, &wg, &m)
	}

	wg.Wait()
}

func worker(worker_id int, wg *sync.WaitGroup, syncMap *MapWrapper ) {
	for j := 0; j < 10; j++ {
		pos, val := rand.Intn(10), rand.Intn(100)
		println(fmt.Printf("Worker №%d is writing value %d to position %d\n", worker_id, val, pos))
		syncMap.SyncWriter(pos, val)
	}
	wg.Done()
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