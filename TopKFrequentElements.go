package main

import (
	"fmt"
	"sort"
)

type Value struct {
	counter int
	value   int
}

func topKFrequent(nums []int, k int) []int {

	occurrence := make(map[int]Value)
	for _, n := range nums {
		tmp := occurrence[n].counter + 1
		occurrence[n] = Value{tmp, n}
	}

	tmp := [][]int{}
	for _, o := range occurrence {
		tmp = append(tmp, []int{o.counter, o.value})
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i][0] < tmp[j][0]
	})

	tmp = tmp[len(tmp)-k:]
	ret := []int{}
	for _, val := range tmp {
		ret = append(ret, val[1])
	}
	return ret
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
