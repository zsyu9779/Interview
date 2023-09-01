/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package array

func search(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	//res := make([]int, 2)
	//var leftBound func([]int, int)int
	//leftBound = func(nums []int, target int) int{
	//	if len(nums) <= 0 {
	//		return -1
	//	}
	//	left, right := 0, len(nums)
	//	for left < right {
	//		mid := left + (right-left)/2
	//		if nums[mid] == target {
	//			right = mid
	//		} else if nums[mid] < target {
	//			left = mid + 1
	//		} else if nums[mid] > target {
	//			right = mid
	//		}
	//	}
	//	if left
	//}
	return nil
}
