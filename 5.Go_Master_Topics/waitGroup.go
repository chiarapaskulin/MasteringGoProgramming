package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1) //increment the WaitGroup counter. (identify the number of goroutines)

		go func(i int) { //launch a goroutine
			defer wg.Done() ///decrement the counter when the goroutine completes.

			//do some work
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Println("Work done for ", i)
		}(i)
	}

	wg.Wait() //main waits to finish until all the wg's have finished
}