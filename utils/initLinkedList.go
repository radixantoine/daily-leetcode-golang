package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitLinkedList(nums []int) *ListNode {
	var head *ListNode
	var tail *ListNode

	for _, n := range nums {
		newNode := &ListNode{n, nil}
		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = tail.Next
		}
	}

	return head
}
