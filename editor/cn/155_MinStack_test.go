package cn

import (
	"testing"
)

//设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//
// 
// push(x) -- 将元素 x 推入栈中。 
// pop() -- 删除栈顶的元素。 
// top() -- 获取栈顶元素。 
// getMin() -- 检索栈中的最小元素。 
// 
//
// 示例: 
//
// MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
// 
// Related Topics 栈 设计

//leetcode submit region begin(Prohibit modification and deletion)
// 辅助栈
type MinStack struct {
	min   int
	stack []int
}

/** initialize your data structure here. */
func Constructor_MinStack() MinStack {
	return MinStack{
		min:   int(^uint(0) >> 1),
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	if x <= this.min {
		// 上个最小值压入栈中
		this.stack = append(this.stack, this.min)
		this.min = x
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	if this.stack[len(this.stack)-1] == this.min {
		this.min = this.stack[len(this.stack)-2] // 返回上一个最小值
		this.stack = this.stack[:len(this.stack)-2]
	} else {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestMinStack(t *testing.T) {
	obj := Constructor_MinStack()
	obj.Push(2)
	obj.Push(-1)
	obj.Push(300)
	obj.Push(3)
	obj.Push(1)
	obj.Push(-1)
	obj.Pop()
	//obj.Push(1)
}

// 辅助栈
//type MinStack struct {
//	min   []int
//	stack []int
//}
//
///** initialize your data structure here. */
//func Constructor_MinStack() MinStack {
//	return MinStack{
//		min:   make([]int, 0),
//		stack: make([]int, 0),
//	}
//}
//
//func (this *MinStack) Push(x int) {
//	this.stack = append(this.stack, x)
//	if len(this.min) == 0 {
//		this.min = append(this.min, x)
//		return
//	}
//	if x <= this.min[len(this.min)-1] {
//		this.min = append(this.min, x)
//	} else {
//		this.min = append(this.min, this.min[len(this.min)-1])
//	}
//}
//
//func (this *MinStack) Pop() {
//	if len(this.stack) == 0 {
//		return
//	}
//	this.stack = this.stack[:len(this.stack)-1]
//	this.min = this.min[:len(this.min)-1]
//}
//
//func (this *MinStack) Top() int {
//	if len(this.stack) == 0 {
//		return 0
//	}
//	return this.stack[len(this.stack)-1]
//}
//
//func (this *MinStack) GetMin() int {
//	if len(this.min) == 0 {
//		return 0
//	}
//	return this.min[len(this.min)-1]
//}
