package utils

import "fmt"

func PrintList(l *ListNode) {
	if l == nil {
		return
	}
	for {
		fmt.Println(l.Val)
		if l.Next == nil {
			return
		}
		l = l.Next
	}
}
