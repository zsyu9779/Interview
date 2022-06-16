/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package linkedList

/**
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

*/
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var heap = make(minHeap, 1)
	dummy := ListNode{Val: -1}
	p := &dummy
	for _, head := range lists {
		if head != nil {
			heap.insert(head)
		}
	}
	for !heap.isEmpty() {
		node := heap.delTop()
		p.Next = node
		if node.Next != nil {
			heap.insert(node.Next)
		}
		p = p.Next
	}
	return dummy.Next
}

type minHeap []*ListNode

func (h minHeap) top() *ListNode {
	return h[1]
}

func (h minHeap) isEmpty() bool {
	return len(h) == 1
}

func (h *minHeap) insert(node *ListNode) {
	*h = append(*h, node)
	h.swim(len(*h) - 1)
}

func (h *minHeap) delTop() *ListNode {
	top := (*h)[1]
	h.swap(1, len(*h)-1)
	*h = (*h)[:len(*h)-1]
	h.sink(1)
	return top
}

func (h *minHeap) swim(x int) {
	for x > 1 && (*h)[x].Val < (*h)[parent(x)].Val {
		h.swap(x, parent(x))
		x = parent(x)
	}
}

func (h *minHeap) sink(x int) {

	//左孩子idx大于数组长度则右孩子一定大于
	for left(x) < len(*h) {
		//先假设左孩子比父节点大
		older := left(x)
		if right(x) < len(*h) && (*h)[x].Val > (*h)[right(x)].Val {
			older = right(x)
		}
		if (*h)[x].Val <= (*h)[older].Val {
			break
		}
		h.swap(x, older)
		x = older
	}
}

func (h *minHeap) swap(i int, j int) {
	temp := (*h)[i]
	(*h)[i] = (*h)[j]
	(*h)[j] = temp
}
func parent(i int) int {
	return i / 2
}
func left(i int) int {
	return i * 2
}
func right(i int) int {
	return i*2 + 1
}
