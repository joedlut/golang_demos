package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	if right < left {
		fmt.Println("Not found")
		return -1
	}
	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			fmt.Println("Found at position: ", mid)
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1

}

func main() {

	arr := []int{1, 5, 100, 0, -100, 15, 4, 102, 30, 1000}
	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)

	binarySearch(arr, 0)

	// Find the number by looking at the center of the array, choosing
	// the left or right side depending on the value and then continue
	// to halve until the result has been found.
}
