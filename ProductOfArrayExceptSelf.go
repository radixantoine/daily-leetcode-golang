package main

import "fmt"

func productExceptSelf(nums []int) []int {
	l := len(nums)
	res := make([]int, l)

	tmp := 1
	for i := 0; i < l; i++ {
		res[i] = tmp
		tmp *= nums[i]
	}
	tmp = 1
	for i := l - 1; i >= 0; i-- {
		res[i] *= tmp
		tmp *= nums[i]
	}
	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
