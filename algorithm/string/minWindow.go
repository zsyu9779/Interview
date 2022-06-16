/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package string

func minWindow(s string, t string) string {
	subSMap, tMap := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
		subSMap[t[i]]++
	}
	start, end := 0, 0
	for end < len(s) {
		num, ok := tMap[s[end]]
		if !ok {
			end++
			continue
		}
		if num == 1 {
			delete(tMap, s[end])
			subSMap[s[end]] = subSMap[s[end]] - 1
			if len(tMap) == 0 {
				break
			}
			end++
			continue
		}
		tMap[s[end]] = tMap[s[end]] - 1
		subSMap[s[end]] = subSMap[s[end]] - 1
		end++
	}
	for start < end {
		subSMap[s[start]] = subSMap[s[start]] + 1
		num := subSMap[s[start]]
		if num >= 0 {
			delete(subSMap, s[start])
		}
		if len(subSMap) <= 0 {
			break
		}
		start++
	}
	if len(tMap) != 0 {
		return ""
	} else {
		return s[start:end]
	}
}
