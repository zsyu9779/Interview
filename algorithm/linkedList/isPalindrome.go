/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package linkedList

type stack struct {
	stackList []*ListNode
}

func (s *stack) push(node *ListNode)  {
	s.stackList = append(s.stackList, node)
}
func (s *stack) pop() *ListNode{
	node := s.stackList[len(s.stackList)-1]
	s.stackList = s.stackList[len(s.stackList)-2:len(s.stackList)-1]
	return node
}
func (s *stack) isEmpty() bool{
	return len(s.stackList)>0
}

/**
	判断一个链表是否为回文结构
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var stack stack
	rightIndex := head.Next
	cur := head
	for cur.Next != nil && cur.Next.Next != nil {
		rightIndex = rightIndex.Next
		cur = cur.Next.Next
	}
	for rightIndex != nil{
		stack.push(rightIndex)
		rightIndex = rightIndex.Next
	}
	for !stack.isEmpty() {
		if head.Val != stack.pop().Val {
			return false
		}
		head = head.Next
	}
	return true
}
