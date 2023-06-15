package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var ans []int
	map1 := make(map[int]int)
	for i, num := range nums {
		if _, ok := map1[target-num]; !ok {
			map1[num] = i
		} else {
			ans = append(ans, map1[target-num], i)
		}
	}
	return ans
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(twoSum(arr, 6))
}
