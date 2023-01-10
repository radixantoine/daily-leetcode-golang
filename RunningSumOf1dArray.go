package main

import "fmt"

func runningSum(nums []int) []int {
	ret := []int{}
	sum := 0

	for _, n := range nums {
		sum += n
		ret = append(ret, sum)
	}
	return ret
}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
}
