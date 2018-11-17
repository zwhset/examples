// 栈 Stack 后进先出
package main

import "fmt"

type Stack struct {
	Head *StackNode
	Size int
}

type StackNode struct {
	Pre  *StackNode
	Data interface{}
}

func NewStack() *Stack {
	return new(Stack)
}

// 推出栈
func (s *Stack) Pop() interface{} {
	// is Empty return nil
	if s.Head == nil {
		return nil
	}

	n := s.Head
	s.Head = s.Head.Pre
	s.Size--
	fmt.Println("Call Pop() => ", n.Data)
	return n

}

// 推入栈
func (s *Stack) Push(v interface{}) {
	n := new(StackNode)
	n.Data = v
	// if nil Init a Node
	if s.Head == nil {
		s.Head = n
		s.Size = 1
	} else {
		n.Pre = s.Head
		s.Head = n
		s.Size++
	}
}

func main() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
}
