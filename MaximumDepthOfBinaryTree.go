package main

import (
	"daily-leetcode-golang/utils"
	"fmt"
	"math"
)

type TreeNode = utils.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}

func main() {
	fmt.Println(
		maxDepth(utils.InitBinaryTree([]int{1, 0, 2})))
}
