package main

import "fmt"

func mergeSortedLists(arr1 []int, arr2 []int) []int {
	i := 0
	j := 0
	sortedArr := []int{}
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			sortedArr = append(sortedArr, arr1[i])
			i++
		} else {
			sortedArr = append(sortedArr, arr2[j])
			j++
		}
	}

	if i < len(arr1) {
		slicedArr := arr1[:len(arr1)-i]
		sortedArr = append(sortedArr, slicedArr...)
	}
	if j < len(arr2) {
		slicedArr := arr2[:len(arr2)-j]
		sortedArr = append(sortedArr, slicedArr...)

	}

	return sortedArr

}

func MergeSort(array []int) []int {
	var mergeDivideAndConq func([]int, int, int) []int

	mergeDivideAndConq = func(arr []int, start int, end int) []int {
		if start == end {
			return []int{arr[start]}
		}

		mid_pt := (start + end) / 2
		firstHalf := mergeDivideAndConq(arr, start, mid_pt)
		secondHalf := mergeDivideAndConq(arr, mid_pt+1, end)

		return mergeSortedLists(firstHalf, secondHalf)
	}

	return mergeDivideAndConq(array, 0, len(array)-1)

}

func main() {
	inputArray := []int{12, 66, 25, 73, 114}

	fmt.Println(MergeSort(inputArray))
}
