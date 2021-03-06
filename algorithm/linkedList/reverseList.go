/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

func reverseList(head *Node) *Node {
	var pre,next *Node
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}