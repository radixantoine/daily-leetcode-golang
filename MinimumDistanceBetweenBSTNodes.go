package main

import (
	"daily-leetcode-golang/utils"
	"fmt"
	"math"
)

type TreeNode = utils.TreeNode

func minDiffInBST(root *TreeNode) int {
	var prev *TreeNode
	minDiff := math.MaxInt64

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil {
			tmp := node.Val - prev.Val
			if tmp < 0 {
				tmp *= -1
			}
			if tmp < minDiff {
				minDiff = tmp
			}

		}
		prev = node
		inorder(node.Right)
	}
	inorder(root)
	return minDiff
}

func main() {
	fmt.Println(minDiffInBST(
		utils.InitBinaryTree([]int{90, 69, 10, 49, 88, 1, 52})))
}
