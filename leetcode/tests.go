package main

import "fmt"

func main() {
	s := []string{"1", "2"}
	s = s[:0]
	fmt.Println(s)
}
