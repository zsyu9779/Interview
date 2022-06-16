/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package linkedList

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result =&ListNode{}
	var indexNode  = result
	addNum := 0
	for l1 != nil || l2 != nil {
		var sumNum int
		if l2 == nil {
			sumNum = l1.Val+addNum
			if l1.Next != nil {
				l1 = l1.Next
			}else {
				l1 = nil
			}
		}else if l1 == nil {
			sumNum = l2.Val+addNum

			if l2.Next != nil {
				l2 = l2.Next
			}else {
				l2 = nil
			}
		}else {
			sumNum = l1.Val +l2.Val+addNum

			if l2.Next != nil {
				l2 = l2.Next
			}else {
				l2 = nil
			}
			if l1.Next != nil {
				l1 = l1.Next
			}else {
				l1 = nil
			}
		}
		if sumNum >=10 {
			sumNum = sumNum%10
			addNum = 1
			if (l1 == nil && l2.Next == nil) || (l2 == nil && l1.Next == nil){
				indexNode.Val = sumNum
				indexNode.Next = &ListNode{Val: 1}
				indexNode = indexNode.Next
				continue
			}
		}else {
			addNum = 0
		}
		indexNode.Val = sumNum
		fmt.Printf("数值：%d\t",sumNum)
		if l1.Next != nil && l2.Next != nil {
			indexNode.Next = &ListNode{}
			indexNode = indexNode.Next
		}


	}
	return result
}
