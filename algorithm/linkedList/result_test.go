/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

import "testing"

func TestK(t *testing.T) {
	head := &Node{}
	head.value = 1
	head.next = &Node{
		value: 2,
	}
	head.next.next = &Node{
		value: 3,
	}
	head.next.next.next = &Node{
		value: 4,
	}
	head.next.next.next.next = &Node{
		value: 5,
	}
	head = reverseList(head)
	t.Log(head)
}