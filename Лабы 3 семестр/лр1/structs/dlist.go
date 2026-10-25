package structs

import (
	"fmt"
	"strings"
)

type DNode struct {
	Data string
	Prev *DNode
	Next *DNode
}

type DList struct {
	Head  *DNode
	Tail  *DNode
	Count int
}

func NewDList() *DList { return &DList{} }

func (l *DList) PushHead(v string) {
	n := &DNode{Data: v, Next: l.Head}
	if l.Head != nil {
		l.Head.Prev = n
	}
	l.Head = n
	if l.Tail == nil {
		l.Tail = n
	}
	l.Count++
}

func (l *DList) PushTail(v string) {
	n := &DNode{Data: v, Prev: l.Tail}
	if l.Tail != nil {
		l.Tail.Next = n
	}
	l.Tail = n
	if l.Head == nil {
		l.Head = n
	}
	l.Count++
}

func (l *DList) PushAfter(target, v string) {
	p := l.Head
	for p != nil && p.Data != target {
		p = p.Next
	}
	if p == nil {
		return
	}
	n := &DNode{Data: v, Prev: p, Next: p.Next}
	if p.Next != nil {
		p.Next.Prev = n
	} else {
		l.Tail = n
	}
	p.Next = n
	l.Count++
}

func (l *DList) PushBefore(target, v string) {
	p := l.Head
	for p != nil && p.Data != target {
		p = p.Next
	}
	if p == nil {
		return
	}
	n := &DNode{Data: v, Prev: p.Prev, Next: p}
	if p.Prev != nil {
		p.Prev.Next = n
	} else {
		l.Head = n
	}
	p.Prev = n
	l.Count++
}

func (l *DList) PopHead() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
	if l.Head != nil {
		l.Head.Prev = nil
	} else {
		l.Tail = nil
	}
	l.Count--
}

func (l *DList) PopTail() {
	if l.Tail == nil {
		return
	}
	l.Tail = l.Tail.Prev
	if l.Tail != nil {
		l.Tail.Next = nil
	} else {
		l.Head = nil
	}
	l.Count--
}

func (l *DList) PopAfter(target string) {
	p := l.Head
	for p != nil && p.Data != target {
		p = p.Next
	}
	if p == nil || p.Next == nil {
		return
	}
	t := p.Next
	p.Next = t.Next
	if t.Next != nil {
		t.Next.Prev = p
	} else {
		l.Tail = p
	}
	l.Count--
}

func (l *DList) PopBefore(target string) {
	p := l.Head
	for p != nil && p.Data != target {
		p = p.Next
	}
	if p == nil || p.Prev == nil {
		return
	}
	t := p.Prev
	p.Prev = t.Prev
	if t.Prev != nil {
		t.Prev.Next = p
	} else {
		l.Head = p
	}
	l.Count--
}

func (l *DList) RemoveByValue(v string) {
	p := l.Head
	for p != nil && p.Data != v {
		p = p.Next
	}
	if p == nil {
		return
	}
	if p.Prev != nil {
		p.Prev.Next = p.Next
	} else {
		l.Head = p.Next
	}
	if p.Next != nil {
		p.Next.Prev = p.Prev
	} else {
		l.Tail = p.Prev
	}
	l.Count--
}

func (l *DList) Find(v string) bool {
	for p := l.Head; p != nil; p = p.Next {
		if p.Data == v {
			return true
		}
	}
	return false
}

func (l *DList) PrintForward() {
	var res []string
	for p := l.Head; p != nil; p = p.Next {
		res = append(res, p.Data)
	}
	fmt.Printf("Вперёд:   [%s]\n", strings.Join(res, " <-> "))
}

func (l *DList) PrintBackward() {
	var res []string
	for p := l.Tail; p != nil; p = p.Prev {
		res = append(res, p.Data)
	}
	fmt.Printf("Назад:    [%s]\n", strings.Join(res, " <-> "))
}

func (l *DList) PrintDetails() {
	fmt.Println("Подробный вывод:")
	i := 0
	for p := l.Head; p != nil; p = p.Next {
		fmt.Printf("  [%d] addr=%p prev=%p data=\"%s\" next=%p\n", i, p, p.Prev, p.Data, p.Next)
		i++
	}
}

func (l *DList) Save() string {
	var res []string
	for p := l.Head; p != nil; p = p.Next {
		res = append(res, p.Data)
	}
	return strings.Join(res, " ")
}

func (l *DList) Load(parts []string) {
	for _, v := range parts {
		l.PushTail(v)
	}
}
