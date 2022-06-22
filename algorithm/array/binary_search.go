/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package array

//基本二分搜索（搜一个数）

func binarySearch(nums []int, target int) int {
	//左闭右闭区间
	left, right := 0, len(nums)-1

	for left <= right {
		//防止整型值溢出
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			//mid已经搜索过 所以 -1/+1
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

//寻找左侧边界：在基础方法for循环内返回值时处理

func binarySearchLeftEdge(nums []int, target int) int {
	//左闭右闭区间
	left, right := 0, len(nums)-1

	for left <= right {
		//防止整型值溢出
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			//mid已经搜索过 所以 -1/+1
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	//for循环最后一次满足条件时 left==right 出循环时left可能大于right的初始值，即数组越界,左边界可能不是目标值（即不存在目标值）排除这两种情况
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	//存在target的时候最后left一定等于right
	return left
}

//寻找右侧边界：在左侧边界的基础上 收缩左侧指针

func binarySearchRightEdge(nums []int, target int) int {
	//左闭右闭区间
	left, right := 0, len(nums)-1

	for left <= right {
		//防止整型值溢出
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			//mid已经搜索过 所以 -1/+1
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	//for循环最后一次满足条件时 left==right 出循环时right可能小于left的初始值，即数组越界,右边界可能不是目标值（即不存在目标值）排除这两种情况
	if right < 0 || nums[right] != target {
		return -1
	}
	//存在target的时候最后left一定等于right
	return right
}
