/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

import "fmt"

/*
	给定两个有序链表的头指针head1和head2，打印两个链表的公共部分
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func printCommonPart(head1, head2 *ListNode) {
	for head1 != nil && head2 != nil {
		if head1.Val <head2.Val {
			head1 = head1.Next
		}else if head2.Val < head1.Val {
			head2 = head2.Next
		}else {
			fmt.Println(head1.Val)
			head1 = head1.Next
			head2 = head2.Next
		}
	}
}
