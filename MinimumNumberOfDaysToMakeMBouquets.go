package main

import (
	"fmt"
)

func MinMax(array []int) (int, int) {
	min, max := array[0], array[0]

	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}

	makeBloom := func(days int) bool {
		boquetCount, bloomedCount := 0, 0
		for _, i := range bloomDay {
			if i <= days {
				bloomedCount++
			} else {
				bloomedCount = 0
			}
			if bloomedCount >= k {
				bloomedCount = 0
				boquetCount++
			}
		}
		return boquetCount >= m
	}

	left, right := MinMax(bloomDay)
	for left < right {
		mid := (left + right) / 2
		if makeBloom(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 1))
}
