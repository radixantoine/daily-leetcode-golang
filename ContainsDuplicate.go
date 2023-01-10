package main

import "fmt"

func containsDuplicate(nums []int) bool {
	t := make(map[int]bool)
	for _, n := range nums {
		if t[n] {
			return true
		}
		t[n] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4, 5, 6, 1, 7}))
}
