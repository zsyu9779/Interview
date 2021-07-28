/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package binary_tree

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

func Traverse(node *TreeNode) {
	fmt.Println(node.val)
	if node.left != nil {
		Traverse(node.left)
	}
	// fmt.Println(node.val)在这里输出就是中序
	if node.right != nil {
		Traverse(node.right)
	}
	// fmt.Println(node.val) 在这里输出是后序
}

/*
	非递归方式实现前序中序遍历
*/
type stack []*TreeNode

func preOrderTraverse(head *TreeNode) {
	if head != nil {
		nodeStack :=stack{}
		nodeStack = append(nodeStack, head)
		for len(nodeStack)>0 {
			head = nodeStack[len(nodeStack)-1]
			nodeStack = nodeStack[0: len(nodeStack)-1]
			fmt.Println(head.val)
			if head.right != nil {
				nodeStack = append(nodeStack,head.right)
			}
			if head.left != nil {
				nodeStack = append(nodeStack,head.left)
			}
		}
	}
}
