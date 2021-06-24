package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	mc := MapCounter{m: make(map[int]int)}
	go runWriter(mc, 10)
	go runReaders(mc, 10)
	go runReaders(mc, 10)
	time.Sleep(15 * time.Second)
}

type MapCounter struct {
	m map[int]int
	sync.RWMutex
}

func runWriter(mc MapCounter, n int) {
	for i := 0; i <= n; i++ {
		mc.Lock()		//WRITE LOCK
		mc.m[i] = i * 10
		mc.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func runReaders(mc MapCounter, n int) {
	for {
		mc.RLock()		//READ LOCK
		v := mc.m[rand.Intn(n)]
		mc.RUnlock()
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
}
