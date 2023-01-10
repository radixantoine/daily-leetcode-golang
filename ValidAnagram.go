package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ret := make(map[int]int)
	for i := 0; i < len(s); i++ {
		ret[int(s[i])] += 1
		ret[int(t[i])] -= 1
	}

	for _, v := range ret {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
