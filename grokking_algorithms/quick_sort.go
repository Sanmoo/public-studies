package main

import "fmt"

func qSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	lower := make([]int, 0, len(arr))
	higher := make([]int, 0, len(arr))

	for i := 1; i < len(arr); i++ {
		value := arr[i]

		if value < pivot {
			lower = append(lower, value)
		}

		if value > pivot {
			higher = append(higher, value)
		}
	}

	lower = qSort(lower)
	higher = qSort(higher)

	return append(append(lower, pivot), higher...)
}

// func main() {
// 	arr := []int{1, 100, 56, 3, 78, 21, 89, 54}
// 	arr = qSort(arr)
// 	fmt.Printf("%v\n", arr)
// }
