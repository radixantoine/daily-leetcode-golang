package utils

import "fmt"

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			fmt.Printf("%d ", node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		fmt.Println()
	}
}
