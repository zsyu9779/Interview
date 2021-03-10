/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

func reverseList(head *ListNode) *ListNode {
	var pre,next *ListNode
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}