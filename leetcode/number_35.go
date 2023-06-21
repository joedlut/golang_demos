package main

import "fmt"

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/search-insert-position
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
2分查找
*/

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	//left <= right
	for left <= right {
		mid = (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}
	/*if right < left {
		fmt.Println(left, right, mid)
		//最小元素
		if right < 0 {
			return 0
		}
		//最大元素
		if left >= len(nums) {
			return len(nums)
		}
	}*/
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, -1))
}
