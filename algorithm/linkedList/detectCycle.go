/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

/*
	检测链表是否存在环路
*/

func detectCycle(head *ListNode) bool{
	slow := head.next
	fast := head.next.next
	for fast.next !=nil&&fast.next.next !=nil {
		if fast == slow {
			return true
		}
		slow = slow.next
		fast = fast.next.next
	}
	return false
}

/*
	检测链表是否存在环路,若有，返回环的入口节点，若无 nil
*/

func detectCycle1(head *ListNode) *ListNode{
	slow := head.next
	fast := head.next.next
	times:=0
	for fast.next !=nil&&fast.next.next !=nil {
		if fast == slow {
			times++
		}
		if times ==2 {
			return slow
		}
		slow = slow.next
		fast = fast.next.next
	}
	return nil
}