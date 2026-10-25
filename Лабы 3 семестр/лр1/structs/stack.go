package structs

import (
	"fmt"
	"strings"
)

type StackNode struct {
	Data string
	Next *StackNode
}

type Stack struct {
	Top   *StackNode
	Count int
}

func NewStack() *Stack { return &Stack{} }

func (s *Stack) Push(v string) {
	s.Top = &StackNode{Data: v, Next: s.Top}
	s.Count++
}

func (s *Stack) Pop() string {
	if s.Top == nil {
		return ""
	}
	v := s.Top.Data
	s.Top = s.Top.Next
	s.Count--
	return v
}

func (s *Stack) TopVal() string {
	if s.Top == nil {
		return ""
	}
	return s.Top.Data
}

func (s *Stack) Print() {
	var res []string
	s.printRec(s.Top, &res)
	top := s.TopVal()
	fmt.Printf("Стек:    [%s]  (вершина = %s)\n", strings.Join(res, ", "), top)
}

func (s *Stack) printRec(n *StackNode, res *[]string) {
	if n == nil {
		return
	}
	s.printRec(n.Next, res)
	*res = append(*res, n.Data)
}

func (s *Stack) PrintDetails() {
	fmt.Println("Подробный вывод стека:")
	pos := 0
	for p := s.Top; p != nil; p = p.Next {
		fmt.Printf("  поз=%d addr=%p data=\"%s\" next=%p\n", pos, p, p.Data, p.Next)
		pos++
	}
}

func (s *Stack) Save() string {
	var res []string
	var collect func(*StackNode)
	collect = func(n *StackNode) {
		if n == nil {
			return
		}
		collect(n.Next)
		res = append(res, n.Data)
	}
	collect(s.Top)
	return strings.Join(res, " ")
}

func (s *Stack) Load(parts []string) {
	for _, v := range parts {
		s.Push(v)
	}
}
