package main

import "fmt"

func insert1(arr []int, value int) {
	arr = append(arr, value)
}

func main() {
	arr := []int{0}
	fmt.Println(arr)
	insert1(arr, 10)
	fmt.Println(arr)
	arr = append(arr, 10)
	fmt.Println(arr)

	arr1 := make([]int, 1, 100)
	//arr1[10] = 100
	fmt.Println(arr1)
}
