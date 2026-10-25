package structs

import (
	"fmt"
	"strings"
)

type SNode struct {
	Data string
	Next *SNode
}

type SList struct {
	Head  *SNode
	Count int
}

func NewSList() *SList { return &SList{} }

func (l *SList) PushHead(v string) {
	l.Head = &SNode{Data: v, Next: l.Head}
	l.Count++
}

func (l *SList) PushTail(v string) {
	n := &SNode{Data: v}
	if l.Head == nil {
		l.Head = n
	} else {
		p := l.Head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = n
	}
	l.Count++
}

func (l *SList) PushAfter(target, v string) {
	p := l.Head
	for p != nil && p.Data != target {
		p = p.Next
	}
	if p == nil {
		return
	}
	p.Next = &SNode{Data: v, Next: p.Next}
	l.Count++
}

func (l *SList) PushBefore(target, v string) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == target {
		l.PushHead(v)
		return
	}
	p := l.Head
	for p.Next != nil && p.Next.Data != target {
		p = p.Next
	}
	if p.Next != nil {
		p.Next = &SNode{Data: v, Next: p.Next}
		l.Count++
	}
}

func (l *SList) PopHead() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
	l.Count--
}

func (l *SList) PopTail() {
	if l.Head == nil {
		return
	}
	if l.Head.Next == nil {
		l.Head = nil
		l.Count--
		return
	}
	p := l.Head
	for p.Next.Next != nil {
		p = p.Next
	}
	p.Next = nil
	l.Count--
}

func (l *SList) PopAfter(target string) {
	p := l.Head
	for p != nil && p.Data != target {
		p = p.Next
	}
	if p != nil && p.Next != nil {
		p.Next = p.Next.Next
		l.Count--
	}
}

func (l *SList) PopBefore(target string) {
	if l.Head == nil || l.Head.Data == target {
		return
	}
	if l.Head.Next != nil && l.Head.Next.Data == target {
		l.PopHead()
		return
	}
	p := l.Head
	for p.Next != nil && p.Next.Next != nil && p.Next.Next.Data != target {
		p = p.Next
	}
	if p.Next != nil && p.Next.Next != nil {
		p.Next = p.Next.Next
		l.Count--
	}
}

func (l *SList) RemoveByValue(v string) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == v {
		l.PopHead()
		return
	}
	p := l.Head
	for p.Next != nil && p.Next.Data != v {
		p = p.Next
	}
	if p.Next != nil {
		p.Next = p.Next.Next
		l.Count--
	}
}

func (l *SList) Find(v string) bool {
	for p := l.Head; p != nil; p = p.Next {
		if p.Data == v {
			return true
		}
	}
	return false
}

func (l *SList) ToSlice() []string {
	var res []string
	for p := l.Head; p != nil; p = p.Next {
		res = append(res, p.Data)
	}
	return res
}

func (l *SList) PrintForward() {
	fmt.Printf("Список:  [%s]\n", strings.Join(l.ToSlice(), " -> "))
}

func (l *SList) printReverseRec(n *SNode, res *[]string) {
	if n == nil {
		return
	}
	l.printReverseRec(n.Next, res)
	*res = append(*res, n.Data)
}

func (l *SList) PrintReverse() {
	var res []string
	l.printReverseRec(l.Head, &res)
	fmt.Printf("Назад:   [%s]\n", strings.Join(res, " <- "))
}

func (l *SList) PrintDetails() {
	fmt.Println("Подробный вывод списка:")
	i := 0
	for p := l.Head; p != nil; p = p.Next {
		fmt.Printf("  [%d] addr=%p data=\"%s\" next=%p\n", i, p, p.Data, p.Next)
		i++
	}
}

func (l *SList) Save() string {
	return strings.Join(l.ToSlice(), " ")
}

func (l *SList) Load(parts []string) {
	for _, v := range parts {
		l.PushTail(v)
	}
}
