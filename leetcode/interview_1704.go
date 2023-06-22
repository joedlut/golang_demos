package main

func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}
	return sum
}

//哈希解法
/*
func missingNumber(nums []int) int {
    m := make(map[int]bool)
    length := len(nums)
    for i := 0;i<len(nums);i++{
        m[nums[i]] = true
    }
    numPos := 0
    for numPos=0;numPos<=length;numPos++{
        if _,ok := m[numPos];!ok{
            break
        }
    }
    return numPos

}
*/

func main() {

}
