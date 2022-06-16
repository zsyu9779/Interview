/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

/*
	删除链表中间节点

*/
func removeMidNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		return head.Next
	}
	slow := head
	fast := head.Next.Next
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		//if fast.Next == nil {
		//	break
		//}
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
