package main

import (
	"fmt"
)

func InsertionSort(array []int) []int {
	for i, _ := range array {
		for j := i; j > 0 && array[j] < array[j-1]; {
			array[j], array[j-1] = array[j-1], array[j]
			j--
		}

	}
	return array
}

//driver code to invoke insertion sort.
func main() {
	inputArray := []int{12, 66, 25, 73, 114}

	fmt.Println(InsertionSort(inputArray))
}
