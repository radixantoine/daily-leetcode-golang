package main

import "fmt"

func getRounds(n int, count int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return count
	}
	if n%3 == 0 {
		return (n / 3) + count
	}
	return getRounds(n-2, count+1)
}

func minimumRounds(tasks []int) int {
	done := make(map[int]int)

	for _, n := range tasks {
		done[n]++
	}

	rounds := 0
	for _, n := range done {
		ret := getRounds(n, 0)
		if ret == -1 {
			return -1
		}
		rounds += ret
	}
	return rounds
}

func main() {
	fmt.Println(minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}))
}
