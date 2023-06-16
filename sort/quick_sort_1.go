package main

import (
	"fmt"
	"math/rand"
)

func quick_sort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	//随机找一个开始的temp元素
	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = quick_sort(lowPart)
	highPart = quick_sort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}

func main() {
	arr := []int{12, 4, 67, 8, 9, 90}
	arr1 := quick_sort(arr)
	fmt.Println(arr1)
}
