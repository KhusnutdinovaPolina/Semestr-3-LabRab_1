package structs

import (
	"fmt"
	"strings"
)

type QueueNode struct {
	Data string
	Next *QueueNode
}

type Queue struct {
	Front *QueueNode
	Rear  *QueueNode
	Count int
}

func NewQueue() *Queue { return &Queue{} }

func (q *Queue) Push(v string) {
	n := &QueueNode{Data: v}
	if q.Rear == nil {
		q.Front = n
		q.Rear = n
	} else {
		q.Rear.Next = n
		q.Rear = n
	}
	q.Count++
}

func (q *Queue) Pop() string {
	if q.Front == nil {
		return ""
	}
	v := q.Front.Data
	q.Front = q.Front.Next
	if q.Front == nil {
		q.Rear = nil
	}
	q.Count--
	return v
}

func (q *Queue) FrontVal() string {
	if q.Front == nil {
		return ""
	}
	return q.Front.Data
}

func (q *Queue) Print() {
	var res []string
	for p := q.Front; p != nil; p = p.Next {
		res = append(res, p.Data)
	}
	fmt.Printf("Очередь: [%s]  (начало = %s)\n", strings.Join(res, " -> "), q.FrontVal())
}

func (q *Queue) PrintDetails() {
	fmt.Println("Подробный вывод очереди:")
	pos := 0
	for p := q.Front; p != nil; p = p.Next {
		fmt.Printf("  поз=%d addr=%p data=\"%s\" next=%p\n", pos, p, p.Data, p.Next)
		pos++
	}
}

func (q *Queue) Save() string {
	var res []string
	for p := q.Front; p != nil; p = p.Next {
		res = append(res, p.Data)
	}
	return strings.Join(res, " ")
}

func (q *Queue) Load(parts []string) {
	for _, v := range parts {
		q.Push(v)
	}
}
