package DSA

import (
	"errors"
	"fmt"
	"strconv"
)

type ArrayStack struct {
	maxSize int
	stack   []int
	top     int
}

func NewArrayStack(maxSize int) ArrayStack {
	return ArrayStack{maxSize: maxSize, stack: make([]int, maxSize), top: -1}
}

//栈是否空
func (r *ArrayStack) isEmpty() bool {
	if r.top == -1 {
		return true
	}
	return false
}

//栈是否满
func (r *ArrayStack) isFull() bool {
	if r.top == r.maxSize-1 {
		return true
	}
	return false
}

//入站
func (r *ArrayStack) push(n int) error {

	if r.isFull() {
		return errors.New("栈已满")
	}

	r.top++
	r.stack[r.top] = n

	return nil
}

func (r *ArrayStack) size() int {
	return r.top + 1
}

//出栈
func (r *ArrayStack) pop() (int, error) {

	if r.isEmpty() {
		return 0, errors.New("空栈")
	}

	res := r.stack[r.top]
	r.top--

	return res, nil
}

func (r *ArrayStack) list() error {

	for i := r.top; i >= 0; i-- {
		fmt.Println(r.stack[i])
	}

	return nil
}

func ArrayStackDemo() {
	//实现加减法
	//str := "2*2+1+3="

	//var judgSymbol func(str1, str2 string) int
	//judgSymbol = func(str1, str2 string) int { //判断运算符优先级
	//	if str1 == "*" || str1 == "/" || str2 == "*" || str2 == "/" {
	//		if str1 == "*" || str1 == "/" {
	//			if str2 == "-" || str2 == "+" {
	//				return 1
	//			} else {
	//				return 0
	//			}
	//		} else {
	//			if str2 == "*" || str2 == "/" {
	//				return -1
	//			} else {
	//				return 0
	//			}
	//		}
	//	} else {
	//		return 0
	//	}
	//}

	var isSymbol = func(str string) bool {

		return str == "*" || str == "/" || str == "-" || str == "+"
	}

	////符号栈和数栈
	numStack := NewArrayStack(100)
	//operStack := NewArrayStack(100)
	//for i := range str {
	//	if str[i:i+1] == "=" {
	//		break
	//	}
	//	//判断是否是符号
	//	if isSymbol(str[i : i+1]) { //是符号
	//		//符号栈是否为空
	//		if operStack.isEmpty() {
	//			operStack.push(int(str[i]))
	//		} else { //判断当前符号是否小于等于栈内的符号
	//			chInt, _ := operStack.pop()
	//			ch := string(int32(chInt))
	//			if judgSymbol(str[i:i+1], ch) <= 0 { //当前符号小于等于
	//				//数栈弹出两个数字计算
	//				num1, _ := numStack.pop()
	//				num2, _ := numStack.pop()
	//				numStack.push(calculateSwitch(num1, num2, ch))
	//				operStack.push(int(str[i]))
	//			} else {
	//				operStack.push(chInt)
	//				operStack.push(int(str[i]))
	//			}
	//		}
	//	} else {
	//		number, _ := strconv.Atoi(str[i : i+1])
	//		numStack.push(number)
	//	}
	//}
	//
	////计算剩下符号
	//for true {
	//	if operStack.isEmpty() {
	//		break
	//	}
	//
	//	chInt, _ := operStack.pop()
	//	ch := string(int32(chInt))
	//	num1, _ := numStack.pop()
	//	num2, _ := numStack.pop()
	//	fmt.Println(calculateSwitch(num1, num2, ch), operStack.top)
	//	numStack.push(calculateSwitch(num1, num2, ch))
	//
	//}
	//
	//test, err := numStack.pop()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(test)

	//中缀转后缀
	//2*2+2+3
	str := "2*2+1+3="
	s := ""
	last := ""
	flag := 1
	tempStack := NewArrayStack(100)
	for i := range str {
		if str[i:i+1] == "=" {
			break
		}
		if isSymbol(str[i : i+1]) { //是否是运算符
			//保存当前运算符
			last = str[i : i+1]

		} else {           //不是运算符
			if flag == 1 { //弹出一个数
				if last != "" {
					ch, _ := tempStack.pop()
					chStr := strconv.Itoa(ch)
					s += chStr + str[i:i+1] + last

					flag = 0
				}else {
					num,_:=strconv.Atoi(str[i:i+1])
					tempStack.push(num)
				}
			} else {
				s += str[i:i+1] + last
			}
		}
	}

	// 后缀表达式
	for i := range s {
		if isSymbol(s[i : i+1]) {
			num1, _ := numStack.pop()
			num2, _ := numStack.pop()
			numStack.push(calculateSwitch(num1, num2, s[i:i+1]))
		} else {
			num, err := strconv.Atoi(s[i : i+1])
			if err != nil {
				fmt.Println(err)
				return
			}
			numStack.push(num)
		}
	}

	num1, _ := numStack.pop()
	fmt.Println(num1)

}

func calculateSwitch(num1, num2 int, str string) int {
	switch str {
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	case "-":
		return num1 - num2
	case "+":
		return num1 + num2

	default:
		return 0
	}
}
