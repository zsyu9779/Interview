/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

import "testing"

func TestK(t *testing.T) {
	head := &ListNode{}
	head.value = 1
	head.next = &ListNode{
		value: 2,
	}
	head.next.next = &ListNode{
		value: 3,
	}
	head.next.next.next = &ListNode{
		value: 4,
	}
	head.next.next.next.next = &ListNode{
		value: 5,
	}
	head = reverseList(head)
	t.Log(head)
}