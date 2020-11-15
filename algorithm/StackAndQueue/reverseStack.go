/*
 * Copyright (C) 2020 Baidu, Inc. All Rights Reserved.
 */
package StackAndQueue

import "fmt"

/*
	仅使用递归函数逆序一个栈，不能使用其他数据结构
*/

//第一步：获取栈底元素并删除

func getAndRemoveLastElement(stack []int) int {
	result := stack[len(stack)-1]
	stack = stack[0:len(stack)-1]
	if len(stack) == 0 {
		return result
	}else {
		last := getAndRemoveLastElement(stack)
		stack = append(stack, result)
		return last
	}
}

//递归调用上面的func 使压栈顺序倒转
func reverse(stack []int)  {
	if len(stack) == 0 {
		return
	}
	i := getAndRemoveLastElement(stack)
	reverse(stack)
	stack = append(stack, i)
}

func TestReverse() {
	stack := []int{1,2,3,4,5}
	reverse(stack)
	fmt.Println(stack)
}