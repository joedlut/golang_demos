package main

import "fmt"

func test(arr []int) {
	arr[len(arr)-1] = 1
}

func main() {
	arr := make([]int, 10)
	test(arr)
	fmt.Println(arr)
}
