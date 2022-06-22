/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package array

import "strings"

/*
	滑动窗口

	1、当移动 right 扩大窗口，即加入字符时，应该更新哪些数据？

	2、什么条件下，窗口应该暂停扩大，开始移动 left 缩小窗口？

	3、当移动 left 缩小窗口，即移出字符时，应该更新哪些数据？

	4、我们要的结果应该在扩大窗口时还是缩小窗口时进行更新？
*/

/*
	给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""
*/
const INT_MAX = int(^uint(0) >> 1)

func MinWindow(s string, t string) string {
	sArr := strings.Split(s, "")
	tArr := strings.Split(t, "")
	need, window := make(map[string]int, len(t)), make(map[string]int, len(t))

	for _, s2 := range tArr {
		need[s2]++
	}
	//valid 表示窗口中满足need条件的字符个数
	left, right, valid := 0, 0, 0
	//最小子串起始索引及长度
	start, Slen := 0, INT_MAX

	for right < len(s) {
		c := sArr[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		//如果窗口没包含全部目标字母，这个循环不走，如果包含，尝试缩小窗口
		for valid == len(need) {
			if right-left < Slen {
				//记录上一次的数据
				start = left
				Slen = right - left
			}
			d := sArr[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if Slen == INT_MAX {
		return ""
	} else {
		return strings.Join(sArr[start:start+Slen], "")
	}
}

/*
	给你两个字符串s1=和=s2 ，写一个函数来判断 s2 是否包含 s1=的排列。如果是，返回 true ；否则，返回 false 。
	换句话说，s1 的排列之一是 s2 的 子串 。
*/
func CheckInclusion(s1 string, s2 string) bool {
	sArr := strings.Split(s2, "")
	tArr := strings.Split(s1, "")
	need, window := make(map[string]int, len(s1)), make(map[string]int, len(s1))

	for _, s := range tArr {
		need[s]++
	}
	//valid 表示窗口中满足need条件的字符个数
	left, right, valid := 0, 0, 0

	for right < len(s2) {
		c := sArr[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		//如果窗口没包含全部目标字母，这个循环不走，如果包含，尝试缩小窗口
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := sArr[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
