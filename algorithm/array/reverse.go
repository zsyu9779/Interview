/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package array

//	翻转字符串
func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}
