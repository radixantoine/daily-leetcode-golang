package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	s = strings.ToUpper(s)
	for l < r {
		for unicode.IsLetter(rune(s[l])) == false && unicode.IsNumber(rune(s[l])) == false {
			if l == r {
				return true
			}
			l++
		}
		for unicode.IsLetter(rune(s[r])) == false && unicode.IsNumber(rune(s[r])) == false {
			if r == l {
				return true
			}
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
