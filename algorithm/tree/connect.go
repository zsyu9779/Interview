/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
	填充每个节点的下一个右侧节点

	通过给每个节点 Next赋值
*/

/*
	思路：把每俩相邻节点看作一个节点 形成一颗逻辑上的三叉树 给三插两两关联
*/
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	traverse1(root.Left, root.Right)
	return root
}

func traverse1(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	traverse1(node1.Left, node1.Right)
	traverse1(node1.Right, node2.Left)
	traverse1(node2.Left, node2.Right)
}
