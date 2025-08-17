package main

func selectSort(arr *[]int) {
	for k := 0; k <= len(*arr)-1; k++ {
		minimum := (*arr)[k]

		for i := k + 1; i <= len(*arr)-1; i++ {
			value := (*arr)[i]

			if (minimum) > value {
				(*arr)[i] = minimum
				(*arr)[k] = value

				minimum = value
			}
		}
	}
}

// func main() {
// 	arr := []int{1, 100, 56, 3, 78, 21, 89, 54}
// 	selectSort(&arr)
// 	fmt.Printf("%v\n", arr)
// }
