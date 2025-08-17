package main

import "fmt"

func binarySearch(arr []int, el int) int {
	lowerBound := 0
	upperBound := len(arr) - 1
	fmt.Printf("Lower bound: %d, Upper bound: %d\n", lowerBound, upperBound)
	for upperBound-lowerBound > 1 {
		currentIndex := (upperBound + lowerBound) / 2
		currentValue := arr[currentIndex]

		if currentValue == el {
			return currentIndex
		} else if currentValue > el {
			upperBound = currentIndex
		} else if currentValue < el {
			lowerBound = currentIndex
		}
		fmt.Printf("Lower bound: %d, Upper bound: %d\n", lowerBound, upperBound)
	}

	return -1
}

func createSortedArray(size int) []int {
	array := make([]int, size)

	for i := range array {
		array[i] = i
	}

	return array
}

// func main() {
// 	arr := createSortedArray(100000000)
// 	result := binarySearch(arr, 9999)
// 	println(result)
// }
