/*
 * Copyright (C) 2020 Baidu, Inc. All Rights Reserved.
 */
package StackAndQueue

import "fmt"

func SortStackByStack(stack []int) []int{
	var help []int
	for len(stack)>0 {
		cur := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		for len(help)>0 && help[len(help)-1] <cur {
			stack = append(stack, help[len(help)-1])
			help = help[0:len(help)-1]
		}
		help = append(help, cur)
	}
	for len(help)>0 {
		stack = append(stack, help[len(help)-1])
		help = help[0:len(help)-1]
	}
	return stack
}
func TestSortStackByStack()  {
	a := []int{1,3,2,6,4,7,5}
	fmt.Println(SortStackByStack(a))
}