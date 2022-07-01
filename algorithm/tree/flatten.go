/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package tree

/*
	原地将一个二叉树展开成一个链表，下一个节点在上一个节点的right指针
*/

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	//把左子树移到右子树的位置
	left := root.Left
	right := root.Right
	root.Left = nil
	root.Right = left
	//把原本的右子树移到现有右子树的尾部
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
