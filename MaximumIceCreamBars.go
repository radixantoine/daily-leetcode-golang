package main

import (
	"fmt"
	"sort"
)

func maxIceCream(costs []int, coins int) int {
	tmp := 0
	count := 0

	sort.Ints(costs)

	for _, cost := range costs {
		tmp += cost
		if tmp > coins {
			return count
		}
		count++
	}
	return count
}

func main() {
	fmt.Println(maxIceCream([]int{1, 3, 2, 4, 1}, 4))
}
