/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package linkedList

/**
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/partition-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func partition(node ListNode, x int) ListNode {
	dummy1 := ListNode{1, nil}
	dummy2 := ListNode{1, nil}
	p1 := dummy1
	p2 := dummy2
	p := &node
	for p != nil {
		if p.Val > x {
			p1.Next = p
			p1 = *p1.Next
		} else {
			p2.Next = p
			p2 = *p2.Next
		}
		temp := p.Next
		p.Next = nil
		p = temp
	}
	p1.Next = dummy2.Next
	return *dummy1.Next
}
