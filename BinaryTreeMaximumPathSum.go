package main

import (
	"daily-leetcode-golang/utils"
	"fmt"
)

type TreeNode = utils.TreeNode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxValue := root.Val

	var getMax func(root *TreeNode) int

	getMax = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0, getMax(root.Left))
		right := max(0, getMax(root.Right))
		sum := root.Val + left + right
		if sum > maxValue {
			maxValue = sum
		}

		return max(left, right) + root.Val
	}

	getMax(root)
	return maxValue
}

func main() {
	fmt.Println(maxPathSum(
		utils.InitBinaryTree([]int{1, 10, 4, 9, 3, 2, 22})))

}
