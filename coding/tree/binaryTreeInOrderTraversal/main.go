package binaryTreeInOrderTraversal

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var arr []int
	traversal(root, &arr)
	return arr
}

func traversal(node *TreeNode, arr *[]int)  {
	if node != nil {
		traversal(node.Left, arr)
		*arr = append(*arr, node.Val)
		traversal(node.Right, arr)
	}
}