package main

import "fmt"

/*
三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/three-steps-problem-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func waysToStep(n int) int {
	//m := make(map[int]int)
	//m[0] = 1
	//m[1] = 1
	//m[2] = 2
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	zero := 1
	one := 1
	two := 2
	three := 0
	for i := 3; i <= n; i++ {
		three = (zero + one + two) % 1000000007
		zero = one
		one = two
		two = three
	}
	return three
}

func main() {
	fmt.Println(waysToStep(78))
}
