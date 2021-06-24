package main

import (
	"fmt"
	"sync"
)

func main() {

	sc := new(safeCounter)

	for i := 0; i < 100; i++ {
		go sc.Increment()
		go sc.Decrement()
	}

	fmt.Println(sc.GetValue())
}

type safeCounter struct {
	i int
	sync.Mutex
}

func (sc *safeCounter) Increment() {
	sc.Lock()
	sc.i++
	sc.Unlock()
}

func (sc *safeCounter) Decrement() {
	sc.Lock()
	sc.i--
	sc.Unlock()
}

func (sc *safeCounter) GetValue() int {
	sc.Lock()
	v:= sc.i
	sc.Unlock()
	return v
}
