/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package array

/*
	删除排序数组中的重复数据
*/

func removeDuplicates(nums []int) int {
	fast := 0
	slow := 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
