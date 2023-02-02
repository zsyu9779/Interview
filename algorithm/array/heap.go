/*
 * Copyright (C) 2023 Baidu, Inc. All Rights Reserved.
 */
package array

//基于小顶堆的优先级队列

func parent(root int) int {
	return root / 2
}

func left(root int) int {
	return root * 2
}

func right(root int) int {
	return root*2 + 1
}

type SortHeap interface {
	Insert(val int)
	RemoveRoot() int
	swim(index int)
	sink(index int)
}

type MinHeap struct {
	pq   []int
	size int
}

func NewMinHeap(cap int) *MinHeap {
	pq := make([]int, cap+1)
	for i := 0; i < len(pq); i++ {
		pq[i] = -1
	}
	return &MinHeap{pq: pq, size: 0}
}

func (m *MinHeap) Insert(val int) {
	m.size++
	m.pq[m.size] = val
	m.swim(m.size)
}

func (m *MinHeap) RemoveRoot() int {
	min := m.pq[1]
	m.swap(1, m.size)
	m.pq[m.size] = -1
	m.size--
	//此时头结点是最初的尾结点 让其下沉到正确位置
	m.sink(1)
	return min
}

func (m *MinHeap) swim(index int) {
	for index > 1 && m.less(index, parent(index)) {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

func (m *MinHeap) sink(index int) {
	//沉到堆底结束
	for left(index) <= m.size {
		//假设左子节点较小
		min := left(index)
		if right(index) <= m.size && m.less(right(index), min) {
			min = right(index)
		}
		//如果此节点比两个子节点都小，停止下沉
		if m.less(index, min) {
			break
		}
		m.swap(min, index)
		index = min
	}
}

func (m *MinHeap) min() int {
	return m.pq[1]
}

func (m *MinHeap) swap(i, j int) {
	temp := m.pq[i]
	m.pq[i] = m.pq[j]
	m.pq[j] = temp
}

func (m *MinHeap) less(i, j int) bool {
	if m.pq[i] < m.pq[j] {
		return true
	}
	return false
}
