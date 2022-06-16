/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package array

/*
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值
target的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

*/
func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i, num := range nums {
		if index, ok := hash[target-num]; ok {
			return []int{i, index}
		}
		hash[num] = i
	}
	return nil
}

/**
给定一个排序数组，实现两数之和，但只能使用常量级的额外空间
*/

func twoSum2(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if target == sum {
			result := []int{left + 1, right + 1}
			return result
		} else if target > sum {
			right--
		} else if target < sum {
			left++
		}
	}
	return []int{-1, -1}
}
