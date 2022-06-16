/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverse(head.Next)
	//保证每层递归head.Next.Next都是尾部，所以把head挂到尾，head，Next置为尾
	head.Next.Next = head
	head.Next = nil
	return last
}

var successor *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}

func reverseBetween(head *ListNode, n int, m int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}
