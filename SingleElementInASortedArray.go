package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	for i := 0; i < len(nums); i += 2 {
		if i == len(nums)-1 || nums[i+1] != nums[i] {
			return nums[i]
		}
	}
	return 0
}

func main() {
	fmt.Println(singleNonDuplicate(
		[]int{3, 3, 7, 7, 10, 11, 11}))
}
