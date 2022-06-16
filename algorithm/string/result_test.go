/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package string

import "testing"

func Test(t *testing.T) {
	s := "ADOBECODEBANC"
	T := "ABC"
	t.Log(minWindow(s, T))
}
