package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var InsertNodeToTree func(tree *TreeNode, i int)

func InitBinaryTree(values []int) (root *TreeNode) {
	valuesLen := len(values)
	if valuesLen < 1 {
		return &TreeNode{}
	}
	tree := TreeNode{Val: values[0], Left: nil, Right: nil}

	InsertNodeToTree = func(tree *TreeNode, i int) {
		left := i * 2
		right := left + 1

		if left <= valuesLen {
			tree.Left = &TreeNode{Val: values[left-1]}
			InsertNodeToTree(tree.Left, left)
		}

		if right <= valuesLen {
			tree.Right = &TreeNode{Val: values[right-1]}
			InsertNodeToTree(tree.Right, right)
		}
	}

	InsertNodeToTree(&tree, 1)
	return &tree
}
