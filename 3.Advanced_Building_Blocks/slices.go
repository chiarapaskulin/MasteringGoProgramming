package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[2:4]
	c := a[:3]
	d := a[3:]

	fmt.Println("Slices a:", a, " b:", b, " c:", c, " d:", d)
	fmt.Printf("\nLength of b: %d \n", len(b))
	fmt.Println("b slice only:", b[:len(b)])
	fmt.Println("Capacity of b: ", cap(b))
	fmt.Println("What b actually sees: ", b[:cap(b)])

	fmt.Println("_______________________2")

	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("s1:", s1)
	s2 := s1[2:4]
	s2[0] = 10
	fmt.Println("S1: ", s1)

	fmt.Println("_______________________3")

	s3 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("s3:", s3)
	s4 := make([]int, 2)
	n := copy(s4, s3[2:4])
	fmt.Println("Number of items copied: ", n)
	s4[0] = 10
	fmt.Println("s3: ", s3)
	fmt.Println("s4: ", s4)

	fmt.Println("_______________________4")

	subslice := testSubSlice()
	fmt.Println(subslice, " length: ", "remaining underlying array: ", subslice[:cap(subslice)])

	fmt.Println("_______________________5")

	subslice2 := testCopySubSlice()
	fmt.Println(subslice2, "remaining underlying array: ", subslice2[:cap(subslice2)])

	fmt.Println("_______________________6")

	s5 := []int{1, 2, 3}
	s5 = append(s5, 4, 5, 6)
	s6 := []int{7, 8, 9}
	s5 = append(s5, s6...)
	fmt.Println(s5)

	fmt.Println("_______________________7")

	s7 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s7 = append(s7[:2], s7[4:]...)
	fmt.Println(s7)
}

func testSubSlice() []int {
	s := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}
	return s[1:4]
}

func testCopySubSlice() []int {
	s := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}
	sub := make([]int, 3)
	copy(sub, s[1:4])
	return sub
}
