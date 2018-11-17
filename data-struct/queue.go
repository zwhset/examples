// 队列 先进先出 input: 1, 2, 3, 4 => [Head 1, 2, 3, 4 Tail] => first 1, 2, 3, 4

package main

import "fmt"

type Queue struct {
	Head *QueueNode
	Tail *QueueNode
	Size int
}

type QueueNode struct {
	Pre  *QueueNode
	Next *QueueNode
	Data interface{}
}

// 推出队列的值
func (q *Queue) Pop() interface{} {
	if q.Head == nil {
		return nil
	}

	n := q.Head
	q.Head = q.Head.Next
	fmt.Printf("Call Pop() => %v, Queue Size: %d\n", n.Data, q.Size)
	q.Size--

	return n
}

// 推入队列
func (q *Queue) Push(v interface{}) {
	n := new(QueueNode)
	n.Data = v

	// init A queue
	if q.Head == nil {
		q.Head = n
	} else {
		q.Tail.Next = n
		n.Pre = q.Tail
	}
	q.Tail = n
	q.Size++

}

func main() {
	q := Queue{}
	q.Push("a")
	q.Push("b")
	q.Push("c")
	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
}
