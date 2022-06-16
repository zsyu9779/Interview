/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

import "testing"

func TestK(t *testing.T) {
	head := &ListNode{}
	head.Val = 2
	head.Next = &ListNode{
		Val: 4,
	}
	head.Next.Next = &ListNode{
		Val: 3,
	}

	head1 := &ListNode{}
	head1.Val = 5
	head1.Next = &ListNode{
		Val: 6,
	}
	//head1.Next.Next = &ListNode{
	//	Val: 4,
	//}

	//head.Next.Next.Next.Next.Next = head.Next.Next
	result := addTwoNumbers(head, head1)
	t.Log(result)
}
func TestConstructor(t *testing.T) {
	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	t.Log(list.Get(1))
	list.DeleteAtIndex(1)
	t.Log(list.Get(1))
}
func TestMergeKLists(t *testing.T) {
	head := &ListNode{}
	head.Val = 2
	head.Next = &ListNode{
		Val: 3,
	}
	head.Next.Next = &ListNode{
		Val: 4,
	}

	head1 := &ListNode{}
	head1.Val = 2
	head1.Next = &ListNode{
		Val: 6,
	}
	head3 := &ListNode{}
	head3.Val = 2
	head3.Next = &ListNode{
		Val: 7,
	}
	head4 := &ListNode{}
	head4.Val = 2
	head4.Next = &ListNode{
		Val: 8,
	}
	var list []*ListNode
	list = append(list, head)
	list = append(list, head1)
	list = append(list, head3)
	list = append(list, head4)
	MergeKLists(list)
}
