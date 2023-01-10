package main

import "fmt"

func checkSum(nums []int, target int, i int) (bool, int) {
	for j := i + 1; j < len(nums); j++ {
		if nums[i]+nums[j] == target {
			return true, j
		}
	}
	return false, 0
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		status, j := checkSum(nums, target, i)
		if status == true {
			return []int{i, j}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
