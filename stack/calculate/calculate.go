package calculate

import (
	stack "data-structure/stack/arraystack"
	"strconv"
)

// Calculate 问题: 计算 类似： 3+2*5+(1+1)*2
// 1. 从左向右遍历计算式；利用栈 将 计算式 分为 数字、符号。 分别压入，数字栈、符号栈；便利完成后弹出计算。
//    将结果压入数字栈，用于参与后续计算。
// 2. 考虑到 运算法则 先乘除后加减：则需要考虑压入符号栈时的优先级。
//	  如果待压入符号 优先级低于 栈顶符号。则需将栈内此符号消耗掉。再次尝试符号压栈（递归）
// 3. 考虑到 （），则可将小括号看作对符号。
//    左括号直接入栈，并且定义默认最低优先级，即 不会影响 后续运算符 直接入栈。
//    右括号定义为特殊符号，但遇到‘）’,即开始弹栈计算，直到弹出‘（’为止。  此时一对（）以处理完毕。
func Calculate(calString string) int {

	numStack := stack.NewArrayStack(20)
	symbolStack := stack.NewArrayStack(20)

	numString := ""

	for i := 0; i < len(calString); i++ {
		c := calString[i]
		//fmt.Println(string(c))
		if IsNum(c) {
			numString += string(c)
		} else if IsSymbol(c) {
			// if is charAt symbol: num is end
			if "" != numString {
				num, _ := strconv.Atoi(numString)
				numStack.Push(num)
				numString = ""
			}

			// push symbol:
			pushSymbol(c, numStack, symbolStack)

		} else if isParenthesesL(c) {
			symbolStack.Push(int(c))
		} else if isParenthesesR(c) {

			// if is charAt symbol: num is end
			num, _ := strconv.Atoi(numString)
			numStack.Push(num)
			numString = ""

			for {
				if isParenthesesL(byte(symbolStack.Peek())) {
					break
				}
				var num1 = numStack.Pop()
				//fmt.Println(num1)
				var num2 = numStack.Pop()
				//fmt.Println(num2)
				var symbolInStack = byte(symbolStack.Pop())
				//fmt.Println(symbolInStack)
				var res = cal(num1, num2, symbolInStack)
				//fmt.Println(res)
				numStack.Push(res)
			}
			// pop ( ,remove '('
			symbolStack.Pop()
		}

		// if is end with num , push numStack
		if i == len(calString)-1 && IsNum(c) {
			num, _ := strconv.Atoi(numString)
			numStack.Push(num)
			numString = ""
		}

	}

	//打印此时切片内容
	//fmt.Println(numStack.Stack[0:numStack.Top+1])
	//fmt.Println(symbolStack.Stack[0:symbolStack.Top+1])

	// pop two nums and one symbol , calculate. until symbolStack is empty
	for {
		if symbolStack.IsEmpty() {
			break
		}
		var num1 = numStack.Pop()
		var num2 = numStack.Pop()
		var symbolInStack = byte(symbolStack.Pop())
		var res = cal(num1, num2, symbolInStack)
		numStack.Push(res)
	}

	return numStack.Pop()

}

func isParenthesesL(c byte) bool {
	return c == '('
}

func isParenthesesR(c byte) bool {
	return c == ')'
}

func pushSymbol(c byte, numStack *stack.ArrayStack, symbolStack *stack.ArrayStack) {

	// if symbolStack empty:push directly
	if symbolStack.IsEmpty() {
		symbolStack.Push(int(c))
		return
	}

	// the stack top char
	stackSymbol := byte(symbolStack.Peek())

	// compare symbol priority
	priorityC := getPriority(c)
	priorityStack := getPriority(stackSymbol)

	if priorityC >= priorityStack {
		symbolStack.Push(int(c))
		return
	} else {
		var num1 = numStack.Pop()
		var num2 = numStack.Pop()
		var symbolInStack = byte(symbolStack.Pop())
		var res = cal(num1, num2, symbolInStack)
		numStack.Push(res)
	}
	// recursion
	pushSymbol(c, numStack, symbolStack)

}

func cal(num1 int, num2 int, inStack byte) int {

	switch inStack {
	case '+':
		return num2 + num1
	case '-':
		return num2 - num1
	case '*':
		return num2 * num1
	default:
		return num2 / num1
	}

}

func getPriority(symbol byte) int {
	switch symbol {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0
	}

}

// IsSymbol 是否时运算符
func IsSymbol(c byte) bool {

	switch c {
	case '+':
		return true
	case '-':
		return true
	case '*':
		return true
	case '/':
		return true
	default:
		return false
	}

}

// IsNum 是否是数字
func IsNum(c byte) bool {
	return c >= '0' && c <= '9'
}
