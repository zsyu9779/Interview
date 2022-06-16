/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package array

func search(nums []int, target int) int{
	high := len(nums)-1
	low := 0
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			high = mid-1
		}else {
			low = mid+1
		}
	}
	return -1
}
