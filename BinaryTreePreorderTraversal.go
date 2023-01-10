package main

import (
	"daily-leetcode-golang/utils"
	"fmt"
)

type TreeNode = utils.TreeNode

func rec(root *TreeNode, ret []int) []int {
	if root == nil {
		return ret
	}

	ret = append(ret, root.Val)

	ret = rec(root.Left, ret)
	ret = rec(root.Right, ret)
	return ret
}

func preorderTraversal(root *TreeNode) []int {
	return rec(root, []int{})
}

func main() {
	fmt.Println(preorderTraversal(
		utils.InitBinaryTree([]int{1, 20, 2, 4, 10, 3, 4, 5})))
}
