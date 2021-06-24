package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("MAIN1:")
	main1()

	fmt.Println("MAIN2:")
	main2()
}

func main2(){
	go SlowCounter1(1)
	time.Sleep(10 * time.Second)
}

func SlowCounter1(n int) {
	i := 0
	for {
		d := time.Duration(n) * time.Second //create a duration of n seconds
		t := time.NewTimer(d) //create a timer for this duration
		<-t.C //channel é chamado em cima de t, que só continua a execução depois de passado tempo t

		i++
		fmt.Println(i)
	}
}

func main1() {
	nc := make(chan int)
	stopc := make(chan bool)

	//time duration = 1
	go SlowCounter2(1, nc, stopc)
	time.Sleep(5 * time.Second)

	//time duration = 2
	nc <- 2
	time.Sleep(6 * time.Second)

	//asks to stop
	stopc <- true
	time.Sleep(1 * time.Second)
}

func SlowCounter2(n int, nc chan int, stopc chan bool) {
	i := 0
	d := time.Duration(n) * time.Second

Loop: //beginning label
	for {
		select {
		case <-time.After(d): //use time.After channel to wait for a time period
			i++
			fmt.Println(i)
		case n = <-nc:
			fmt.Println("Timer duration changed to", n)
			d = time.Duration(n) * time.Second
		case <-stopc:
			fmt.Println("Timer stopped")
			break Loop
		}
	}
//end label

	fmt.Println("Exiting Slow Counter")
}
