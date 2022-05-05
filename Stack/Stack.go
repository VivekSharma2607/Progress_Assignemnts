// Implementing Stack using interfaces in Golang 


package main

import "fmt"

type Stack []int

func (s *Stack) IsEmpty() bool {
	if len(*s) == 0{
		return true
	}else{
		return false
	}
}

func (s *Stack) Push(a int) {
	*s = append(*s, a)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}else {
		i := len(*s)-1
		temp := (*s)[i]
		*s = (*s)[:i]
		return temp
	}
}

func main() {
	var stack Stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println(stack)
	stack.Pop()
	fmt.Println("Stack after poping the element = " , stack)
}