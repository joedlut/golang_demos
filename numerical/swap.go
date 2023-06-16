package main

import "fmt"

func main() {
	x, y := 1, 2
	x, y = y, x
	fmt.Println(x)
	fmt.Println(y)

}
