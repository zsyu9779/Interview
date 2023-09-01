/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package linkedList

/*
*
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
			heap.push(head)
		}
	}
	for !heap.isEmpty() {
		node := heap.delTop()
		p.Next = node
		if node.Next != nil {
			heap.push(node.Next)
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

func (h minHeap) less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *minHeap) push(node *ListNode) {
	*h = append(*h, node)
	h.swim(len(*h) - 1)
}

func (h *minHeap) delTop() *ListNode {
	min := (*h)[1]
	h.swap(1, len(*h)-1)
	*h = (*h)[:len(*h)-1]
	h.sink(1)
	return min
}

func (h *minHeap) swim(x int) {
	for x > 1 && !h.less(parent(x), x) {
		h.swap(parent(x), x)
		x = parent(x)
	}
}

func (h *minHeap) sink(x int) {
	for left(x) <= len(*h)-1 {
		if h.less(x, left(x)) {
			break
		}
		// 下沉和两个子节点最大的交换
		h.swap(x, left(x))
		// 如果是和右子节点交换 则左右节点再次交换
		if right(x) <= len(*h)-1 && h.less(right(x), left(x)) {
			h.swap(right(x), left(x))
			x = right(x)
		} else {
			x = left(x)
		}
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
