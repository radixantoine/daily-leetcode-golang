package main

import (
	"daily-leetcode-golang/utils"
)

type TreeNode = utils.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// swap left and right subtrees
	root.Left, root.Right = root.Right, root.Left
	// invert left and right subtrees recursively
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func main() {
	utils.PrintTree(invertTree(
		utils.InitBinaryTree([]int{4, 2, 7, 1, 3, 6, 9})))
}
