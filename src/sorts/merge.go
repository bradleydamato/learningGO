package main

import "fmt"

// my attempt to implement merge sort in Go

func main() {
	test := []int{534, 534523, 34, 634, 23, 6, 43, 2356, 43, 41, 23, 5, 634, 43, 1}
	fmt.Println(test)
	q := MergeSort(test)
	fmt.Println(q)
}

func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	SliceToReturn := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 { // cleanup happens
			SliceToReturn[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 { //cleanup happens
			SliceToReturn[k] = left[i]
			i++
		} else if left[i] < right[j] { //comparison
			SliceToReturn[k] = left[i]
			i++
		} else { // comparison
			SliceToReturn[k] = right[j]
			j++
		}
	}
	return SliceToReturn
}
