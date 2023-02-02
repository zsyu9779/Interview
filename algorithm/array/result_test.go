/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package array

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	s := CheckInclusion("adc", "dcda")
	fmt.Println(s)
}

func TestNewMinHeap(t *testing.T) {
	h := NewMinHeap(20)
	h.Insert(1)
	h.Insert(3)
	h.Insert(8)
	h.Insert(11)
	h.Insert(4)
	h.Insert(7)
	h.Insert(20)
	fmt.Println(h.pq)

	for {
		min := h.RemoveRoot()
		//fmt.Println(h.pq)
		if min == -1 {
			break
		}
		fmt.Println(min)
	}
}
