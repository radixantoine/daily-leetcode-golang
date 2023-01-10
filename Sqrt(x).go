package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}

	left := 0.0
	right := float64(x)

	for {
		mid := (left + right) / 2
		tmp := int(math.Floor(mid * mid))
		if tmp == x {
			return int(math.Floor(mid))
		}
		if tmp < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

}

func main() {
	fmt.Println(mySqrt(8))
}
