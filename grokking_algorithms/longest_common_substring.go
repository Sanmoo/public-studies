package main

import (
	"fmt"
	"strings"
)

func longest(str1 string, str2 string) [][]int {
	dynamicGrid := make([][]int, len(str1))

	for i, iChar := range strings.Split(str1, "") {
		dynamicGrid[i] = make([]int, len(str2))

		for j, jChar := range strings.Split(str2, "") {
			if iChar != jChar {
				dynamicGrid[i][j] = 0
			} else if i == 0 || j == 0 {
				dynamicGrid[i][j] = 1
			} else {
				dynamicGrid[i][j] = dynamicGrid[i-1][j-1] + 1
			}

			fmt.Printf("%v-%v ", iChar, jChar)
		}

		fmt.Print("\n")
	}

	return dynamicGrid
}

func main() {
	resp := longest("fish", "fosh")
	fmt.Print("\n\n")
	prettyPrinty(resp)
}

func prettyPrinty(table [][]int) {
	for i := range table {
		for _, value := range table[i] {
			fmt.Printf("%v ", value)
		}
		fmt.Print("\n")
	}
}
