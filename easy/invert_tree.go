package easy

import "neetcode/types"

func invertTree(root *types.TreeNode) *types.TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	return root
}
