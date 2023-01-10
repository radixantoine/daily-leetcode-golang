package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1] ||
			(points[i][1] == points[j][1] && points[i][0] < points[j][0])
	})

	tmp := points[0][1]
	shots := 1

	for i := 1; i < len(points); i++ {
		if tmp < points[i][0] {
			tmp = points[i][1]
			shots++
		}
	}
	return shots
}

func main() {
	fmt.Println(
		findMinArrowShots([][]int{{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7}}),
	)
}
