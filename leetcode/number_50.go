package main

import (
	"fmt"
	"math"
)

/*
	超时了

	func myPow(x float64, n int) float64 {
		flag := false
		if n == 0 {
			return 1
		}
		if n < 0 {
			flag = true
			n = -n
		}
		var result float64 = 1.0
		for i := 1; i <= n; i++ {
			result *= x
		}
		if !flag {
			return result
		} else {
			return 1.0 / result
		}

}
*/
func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	} else {
		return y * y * x
	}
}

func myPow(x float64, n int) float64 {
	if n > 0 {
		return quickMul(x, n)
	} else {
		return 1 / quickMul(x, -n)
	}
}

func main() {
	fmt.Println(myPow(2.0, -3))
	fmt.Println(math.Pow(2.0, -3))
}
