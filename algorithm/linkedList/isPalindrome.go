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
func isPalindrome(head *ListNode) bool {
	if head == nil || head.next == nil {
		return true
	}
	var stack stack
	rightIndex := head.next
	cur := head
	for cur.next != nil && cur.next.next != nil {
		rightIndex = rightIndex.next
		cur = cur.next.next
	}
	for rightIndex != nil{
		stack.push(rightIndex)
		rightIndex = rightIndex.next
	}
	for !stack.isEmpty() {
		if head.value != stack.pop().value {
			return false
		}
		head = head.next
	}
	return true
}
