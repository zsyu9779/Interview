/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package StackAndQueue

import "fmt"

/*
	给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

	输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
	输出: [3,3,5,5,6,7]
	解释:

	  滑动窗口的位置                最大值
	---------------               -----
	[1  3  -1] -3  5  3  6  7       3
	 1 [3  -1  -3] 5  3  6  7       3
	 1  3 [-1  -3  5] 3  6  7       5
	 1  3  -1 [-3  5  3] 6  7       5
	 1  3  -1  -3 [5  3  6] 7       6
	 1  3  -1  -3  5 [3  6  7]      7

*/
func getMaxWindow(arr []int, w int) []int {
	if len(arr) == 0 || w < 1 || len(arr) < w {
		return nil
	}
	var qmax, res []int
	for i, num := range arr {
		//如果qmax为空或qmax队尾元素小于等于当前元素，将qmax队尾元素弹出
		for len(qmax) > 0 && arr[qmax[len(qmax)-1]] <= num {
			qmax = qmax[:len(qmax)-1]

		}
		//qmax队尾元素大于当前元素，将当前元素插入qmax队尾
		qmax = append(qmax, i)
		//当qmax的队头元素下标不在当前窗口，则视为队头元素过期，将其弹出
		if qmax[0] == i-w {
			qmax = qmax[1:]
		}
		//qmax进入存在完整窗口的区间，取qmax头下标对应的原数组元素
		if i >= w-1 {
			res = append(res, arr[qmax[0]])
		}
	}
	return res
}
func TestGetMaxWindow() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(getMaxWindow(nums, 3))
}
