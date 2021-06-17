package main

import "fmt"

func main() {
	c := make(chan bool)
	go waitAndSay(c, "world")
	fmt.Println("Hello")

	c <- true

	if b := <-c; b {  //<-c
		fmt.Println("Program now signalled to exit")
	}

	fmt.Printf("\nMain2:\n")
	main2()
}

func waitAndSay(c chan bool, s string) {
	if b := <-c; b {
		fmt.Println(s)
	}
	c <- true
}

func main2() {
	c := make(chan string)
	go sayHelloMultipleTimes(c, 5)
	for s := range c {
		fmt.Println(s)
	}

	_, ok := <-c //<-c
	fmt.Println("Channel closed? ", !ok)

}

func sayHelloMultipleTimes(c chan string, n int) {
	for i := 0; i < n; i++ {
		c <- "Hello"
	}
	close(c)
}
