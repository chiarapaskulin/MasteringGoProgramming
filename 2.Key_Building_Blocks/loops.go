package main

import "fmt"

func main() {
	//for
	n := 1
	for i := 10; i > 0; i-- {
		n *= i
	}
	fmt.Println("Result:", n)

	//while
	n = 1
	for n <= 50 {
		n += n
	}
	fmt.Println("\nResult:", n)

	//defer inside loop
	fmt.Println("\nCounting")
	for i := 1; i <= 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
