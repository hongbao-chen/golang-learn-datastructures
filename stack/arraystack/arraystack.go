package stack

// ArrayStack 以数组(用定长的切片代替)实现的栈
type ArrayStack struct {
	Size  int
	Stack []int
	Top   int
}

func NewArrayStack(size int) *ArrayStack {
	return &ArrayStack{
		Size:  size,
		Stack: make([]int, size, size),
		Top:   -1,
	}
}

func (a *ArrayStack) IsEmpty() bool {
	return a.Top == -1
}

func (a *ArrayStack) IsFull() bool {
	return a.Top == a.Size-1
}

func (a *ArrayStack) Push(val int) {

	if a.IsFull() {
		panic("stack is full err")
	}

	a.Top++
	a.Stack[a.Top] = val

}

func (a *ArrayStack) Pop() int {

	if a.IsEmpty() {
		panic("stack is empty")
	}

	result := a.Stack[a.Top]
	a.Top--

	return result

}

func (a *ArrayStack) Peek() int {

	if a.IsEmpty() {
		panic("stack is empty")
	}

	return a.Stack[a.Top]

}
