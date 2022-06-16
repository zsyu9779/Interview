/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package linkedList

/*
	在链表类中实现这些功能：

*/
type MyLinkedList struct {
	dummy *Node
}

type Node struct {
	Next *Node
	Pre  *Node
	Val  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/*
	get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
*/
func (l *MyLinkedList) Get(index int) int {
	var head Node
	head.Next = l.dummy
	for i := 0; i <= index; i++ {
		head = *head.Next
		if i == index {
			return head.Val
		}
		if head.Next == nil {
			break
		}
	}
	return -1
}

/*
	addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
*/
func (l *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Next: l.dummy,
		Pre:  nil,
		Val:  val,
	}
	l.dummy = node
}

/*
	addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
*/
func (l *MyLinkedList) AddAtTail(val int) {

	head := l.dummy
	for head != nil {
		if head.Next == nil {
			node := &Node{
				Next: nil,
				Pre:  head,
				Val:  val,
			}
			head.Next = node
			break
		}
		head = head.Next
	}
}

/*
	addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。
	如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
*/
func (l *MyLinkedList) AddAtIndex(index, val int) {
	if index <= 0 {
		l.AddAtHead(val)
		return
	}
	node := &Node{
		Val: val,
	}
	head := l.dummy
	for head != nil && index >= 0 {
		head = head.Next
		index--
		if index == 0 {
			//适配index等于链表长度
			if head == nil {
				l.AddAtTail(val)
				return
			}
			pre := head.Pre
			node.Pre = pre
			node.Next = head
			pre.Next = node
			if next := head.Next; next != nil {
				next.Pre = node
			}
			break
		}
	}
}

/*
	deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
*/
func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		l.dummy = l.dummy.Next
		return
	}
	head := l.dummy
	for head != nil && index > 0 {
		head = head.Next
		index--
		if index == 0 {
			if pre := head.Pre; pre != nil {
				pre.Next = head.Next
				if head.Next != nil {
					head.Next.Pre = pre
				}
			}
			break
		}
	}
}
