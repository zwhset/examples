package main

import (
	"errors"
	"fmt"
)

type List struct {
	head *Node // 头节点
	tail *Node // 尾结点
	size int   // 长度
}

type Node struct {
	Data       interface{} // 值
	next, prev *Node       // 后缀与前驱
}

func New() *List {
	l := &List{}
	l.size = 0
	l.head = nil
	l.tail = nil

	return l
}

func (l *List) Append(d interface{}) {
	n := new(Node)
	n.Data = d

	if l.size == 0 {
		l.head = n
		l.tail = n
		n.prev = nil
		n.next = nil
	} else {
		n.prev = l.tail
		n.next = nil
		l.tail.next = n
		l.tail = n
	}
	l.size++
}

// 在索引后面插入元素
func (l *List) Insert(index int, element interface{}) {
	n := new(Node)
	n.Data = element

	// 如果是在最后面
	if index >= l.size-1 {
		l.Append(element)
	} else { // p -> A    =>    p -> B -> A
		i := 0
		for p := l.head; p != nil; p = p.next {
			if i == index {
				n.prev = p
				n.next = p.next
				p.next.prev = n
				p.next = n

				break
			}
			i++
		}
	}
}

// 统计元素出现的次数
func (l *List) Count(element interface{}) int {
	n := 0
	for p := l.head; p != nil; p = p.next {
		if p.Data == element {
			n++
		}
	}
	return n
}

// 在元素尾部插入 List
func (l *List) Extend(el *List) {
	for p := el.head; p != nil; p = p.next {
		l.Append(p.Data)
	}
}

// 查找值的索引， 如果未找到返回错误
func (l *List) Index(element interface{}) (n int, err error) {
	n = 0
	for p := l.head; p != nil; p = p.next {
		if p.Data == element {
			return n, nil
		}
		n++
	}
	return 0, errors.New(fmt.Sprintf("Not Fund element %v", element))
}

// 删除并返回最后一个元素
func (l *List) Pop() interface{} {
	// a b c => a b return c
	l.tail.prev.next = l.head
	fmt.Printf("next: %v head: %v \n", l.tail.prev.Data, l.head.Data)

	return l.tail.Data
}

// GET
func (l *List) Get(index int) interface{} {
	i := 0
	for p := l.head; p != nil; p = p.next {
		if index == i {
			return p.Data
		}
		i++
	}
	return -1
}

// SET
func (l *List) Set(index int, element interface{}) {
	i := 0
	for p := l.head; p != nil; p = p.next {
		if i == index {
			p.Data = element
			return
		}
	}
}

// reverse
func (l *List) Reverse() *List {
	// copy tmp
	tmp := new(List)

	for p := l.tail; p != nil; p = p.prev {
		tmp.Append(p)
	}
	return tmp
}

func (l *List) Print() string {
	s := ""

	for p := l.head; p != nil; {
		s = fmt.Sprintf("%s %v", s, p.Data)
		fmt.Println(s)
		p = p.next
	}
	return s
}

func main() {
	l := new(List)
	l.Append("A")
	l.Append("B")
	l.Append("B")
	l.Append("B")
	l.Append("A")
	l.Insert(0, "abc")

	o := new(List)
	o.Append(0)
	o.Append(1)
	o.Append(2)
	o.Append(3)
	l.Extend(o)

	// Index
	fmt.Printf("Size: %d Elements: [%s] Count['B']: %d Count['A']: %d \n", l.size, l.Print(), l.Count("B"), l.Count("A"))
	if n, err := l.Index(0); err != nil {
		fmt.Println("Don't fund element.")
	} else {
		fmt.Printf("Find 0, index: %d\n", n)
	}

	// Get
	fmt.Println(l.Get(3))

	// Set
	l.Set(0, "AA")
	fmt.Println(l.Print())

	// Reverse
	tmp := l.Reverse()
	fmt.Printf("Print: %s", tmp.Print())
}
