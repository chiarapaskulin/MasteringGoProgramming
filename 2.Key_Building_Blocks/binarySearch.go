package main

import "fmt"

func binarySearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}

	mid := int((lowIndex + highIndex) / 2)

	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

func iterBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int

	for startIndex < endIndex {
		mid = int((startIndex + endIndex)/2)
		if array[mid] > target {
			endIndex = mid
		} else if array[mid] < target {
			startIndex = mid
		} else {
			return mid
		}
	}

	return -1
}

func main(){
	arr := []int{2,2,4,6,8,10,10,12,12,14,16,18,20,20}

	res1 := binarySearch(arr, 8, 0, len(arr)-1)
	fmt.Println("Index of the element: ", res1)

	res2 := iterBinarySearch(arr, 8, 0, len(arr)-1)
	fmt.Println("Index of the element: ", res2)
}