/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

/*
	删除链表中间节点

*/
func removeMidNode(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	if head.next.next == nil {
		return head.next
	}
	slow := head
	fast := head.next.next
	for slow.next != nil && fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		//if fast.next == nil {
		//	break
		//}
		slow = slow.next
	}
	slow.next = slow.next.next
	return head
}
