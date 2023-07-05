package main

import "fmt"

func insert1(arr []int, value int) {
	arr = append(arr, value)
}

func main() {
	a := 0.2
	b := 0.1
	fmt.Println(a + b)
}
