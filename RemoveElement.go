package main

import "fmt"

func removeElement(nums []int, val int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] == val {
			n--
			nums[i] = nums[n]
		} else {
			i++
		}
	}
	return n
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
