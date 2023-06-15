package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var ans []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ans = append(ans, i, j)
			}
		}
	}
	return ans
}

func main() {
	nums := []int{2, 7, 11, 15}
	ans := twoSum(nums, 9)
	fmt.Println(ans)
}
