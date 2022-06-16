/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

/*
	检测链表是否存在环路
*/

func detectCycle(head *ListNode) bool{
	slow := head.Next
	fast := head.Next.Next
	for fast.Next !=nil&&fast.Next.Next !=nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

/*
	检测链表是否存在环路,若有，返回环的入口节点，若无 nil
*/

func detectCycle1(head *ListNode) *ListNode{
	slow := head.Next
	fast := head.Next.Next
	times:=0
	for fast.Next !=nil&&fast.Next.Next !=nil {
		if fast == slow {
			times++
		}
		if times ==2 {
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}