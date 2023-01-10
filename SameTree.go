package main

import (
	"daily-leetcode-golang/utils"
	"fmt"
)

type TreeNode = utils.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	fmt.Println(isSameTree(
		utils.InitBinaryTree([]int{0, 1, 4, 2}),
		utils.InitBinaryTree([]int{0, 1, 2})))
}
