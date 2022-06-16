/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package array

import "strings"

/*
	最长回文字串
*/
func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		s1 := palindrome(i, i, s)
		s2 := palindrome(i, i+1, s)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func palindrome(left, right int, s string) string {
	strArr := strings.Split(s, "")
	for left >= 0 && right < len(s) && strArr[left] == strArr[right] {
		left--
		right++
	}
	return strings.Join(strArr[left+1:right], "")
}
