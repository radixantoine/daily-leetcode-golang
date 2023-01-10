package main

import "fmt"

func longestConsecutive(nums []int) int {
	tmp := make(map[int]bool)

	for _, v := range nums {
		tmp[v] = true
	}

	longest := 0
	tmpLongest := 0

	for n := range tmp {
		if tmp[n] == false {
			continue
		}
		tmpLongest++
		for i := n + 1; tmp[i] == true; i++ {
			tmpLongest++
			tmp[i] = false
		}
		for i := n - 1; tmp[i] == true; i-- {
			tmpLongest++
			tmp[i] = false
		}
		if tmpLongest > longest {
			longest = tmpLongest
		}
		tmpLongest = 0
	}
	return longest
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
