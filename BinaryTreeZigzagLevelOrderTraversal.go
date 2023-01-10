package main

import (
	"daily-leetcode-golang/utils"
	"fmt"
)

type TreeNode = utils.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	queue := []*TreeNode{root}
	b := false
	for len(queue) > 0 {
		n := len(queue)
		level := make([]int, n)
		b = !b
		for i := 0; i < n; i++ {
			if b == true {
				level[i] = queue[i].Val
			} else {
				level[n-i-1] = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, level)
		queue = queue[n:]
	}
	return res
}

func main() {
	fmt.Println(zigzagLevelOrder(
		utils.InitBinaryTree([]int{3, 9, 20, 0, 0, 15, 7})))
}
