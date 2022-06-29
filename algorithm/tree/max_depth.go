/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package tree

var res = 0
var depth = 0

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	depth++
	if root.Left == nil && root.Right == nil {
		if depth > res {
			res = depth
		}
	}
	traverse(root.Left)
	traverse(root.Right)
	depth--
}

func MaxDepth(root *TreeNode) int {
	traverse(root)
	return res
}
func MaxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := MaxDepth1(root.Left)
	rightMax := MaxDepth1(root.Right)
	if leftMax > rightMax {
		return leftMax + 1
	} else {
		return rightMax + 1
	}
}
