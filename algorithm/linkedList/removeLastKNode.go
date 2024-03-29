/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

type DoubleNode struct {
	value int
	next  *DoubleNode
	pre  *DoubleNode
}

func removeLastKthDoubleNode(head *DoubleNode, K int) *	DoubleNode{
	k1 := K
	head1 := head
	for head1 !=nil {
		k1--
		head1 = head1.next
	}
	if k1 >= 0 {
		return head
	}else {
		for k1 == k1+K {
			k1++
			head1 = head1.pre
		}
		head1.next = head1.next.next
	}
	return head
}

/*双指针(快慢指针)*/
func removeLastKthNode1(head *ListNode, K int) *ListNode {
	newHead := &ListNode{
		Next: head,
	}
	slow,fast := newHead,newHead
	for i:=0;i<K;i++{
		fast = fast.Next
	}
	for fast.Next !=nil{
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return newHead.Next
}
