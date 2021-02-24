/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

type DoubleNode struct {
	value int
	next  *DoubleNode
	pre  *DoubleNode
}

func removeLastKthNode(head *Node, K int) *Node{
	k1 := K
	head1 := head
	for head1 !=nil {
		k1--
		head1 = head1.next
	}
	if k1 == 0 {
		return head
	}
	head1 = head
	for k1 != 0 {
		k1++
		head1 = head1.next
	}
	head1.next = head1.next.next
	return head
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
