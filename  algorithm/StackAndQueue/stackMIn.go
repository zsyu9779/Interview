/*
 * Copyright (C) 2020 Baidu, Inc. All Rights Reserved.
 */
package StackAndQueue

/*
	设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

	push(x) —— 将元素 x 推入栈中。
	pop() —— 删除栈顶的元素。
	top() —— 获取栈顶元素。
	getMin() —— 检索栈中的最小元素。
	 
*/

/*
	解法一：辅助栈法
*/
type MinStack struct {
	stackData []int
	stackMin  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stackData: []int{},
		stackMin:  []int{},
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stackMin) == 0 {
		this.stackMin = append(this.stackMin, x)
	} else if x <= this.GetMin() {
		this.stackMin = append(this.stackMin, x)
	}
	this.stackData = append(this.stackData, x)
}

func (this *MinStack) Pop() {
	if this.stackData[len(this.stackData)-1] == this.GetMin() {
		this.stackMin = this.stackMin[0 : len(this.stackMin)-1]
	}
	this.stackData = this.stackData[0 : len(this.stackData)-1]
}

func (this *MinStack) Top() int {
	return this.stackData[len(this.stackData)-1]
}

func (this *MinStack) GetMin() int {
	return this.stackMin[len(this.stackMin)-1]
}

/*
	解法二：使用一个栈，维护一个最小值变量
*/
type MinStackB struct {
	stackData []int
	Min       int
}

/** initialize your data structure here. */
func ConstructorB() MinStackB {
	return MinStackB{
		stackData: []int{},
	}
}

func (this *MinStackB) Push(x int) {
	//如果push的值 <= Min 则现将原Min值压入栈内，再将x压入
	if len(this.stackData) == 0 {
		this.Min = x
	} else if x <= this.GetMin() {
		this.stackData = append(this.stackData, this.Min)
		this.Min = x
	}
	this.stackData = append(this.stackData, x)
}

func (this *MinStackB) Pop() {
	//pop时栈顶值若等于Min则弹出两次，并将第二次弹出的值赋给min（之前暂存的次小值）
	if this.stackData[len(this.stackData)-1] == this.GetMin() {
		this.stackData = this.stackData[0 : len(this.stackData)-1]
		if len(this.stackData) != 0 {
			this.Min = this.stackData[len(this.stackData)-1]
		}else {
			return
		}
	}
	this.stackData = this.stackData[0 : len(this.stackData)-1]
}

func (this *MinStackB) Top() int {
	return this.stackData[len(this.stackData)-1]
}

func (this *MinStackB) GetMin() int {
	return this.Min
}
