/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package tree

/*
	二叉树的最大直径：即求某个节点左右子树的最大深度和
*/

func DiameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	maxDepth(root)
	return maxDiameter
}

var maxDiameter = 0

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)

	myDiameter := leftMax + rightMax
	maxDiameter = max(myDiameter, maxDiameter)
	return max(leftMax, rightMax) + 1

}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
