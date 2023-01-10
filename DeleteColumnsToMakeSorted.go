package main

import "fmt"

func minDeletionSize(strs []string) int {
	ret := 0
	for i := 0; i < len(strs[0]); i++ {
		tmp := strs[0][i]
		for _, s := range strs {
			if s[i] < tmp {
				ret += 1
				break
			}
			tmp = s[i]
		}
	}
	return ret
}

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
}
