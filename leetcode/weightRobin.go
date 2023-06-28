package main

import (
	"fmt"
	"math/rand"
)

func getNext(m map[string]int) string {
	sum := 0
	for _, weight := range m {
		sum += weight
	}
	arr := make([]string, 0)
	for key, value := range m {
		for i := 0; i < value; i++ {
			arr = append(arr, key)
		}
	}
	index := rand.Intn(sum)
	return arr[index]
}

func main() {
	m := make(map[string]int)
	m["a"] = 2000
	m["b"] = 3000
	m["c"] = 4000
	count := map[string]int{}
	for i := 0; i < 9000; i++ {
		count[getNext(m)]++
	}
	fmt.Println(count)
}
