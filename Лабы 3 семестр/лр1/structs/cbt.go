package structs

import (
	"fmt"
	"strings"
)

type TNode struct {
	Data  string
	Left  *TNode
	Right *TNode
}

type CBT struct {
	Root  *TNode
	Count int
}

func NewCBT() *CBT { return &CBT{} }

func (t *CBT) nodeByIndex(n *TNode, cur, target int) *TNode {
	if n == nil {
		return nil
	}
	if cur == target {
		return n
	}
	l := t.nodeByIndex(n.Left, 2*cur+1, target)
	if l != nil {
		return l
	}
	return t.nodeByIndex(n.Right, 2*cur+2, target)
}

func (t *CBT) Insert(v string) {
	n := &TNode{Data: v}
	if t.Count == 0 {
		t.Root = n
		t.Count = 1
		return
	}
	idx := t.Count
	parentIdx := (idx - 1) / 2
	isLeft := (idx % 2 == 1)
	parent := t.nodeByIndex(t.Root, 0, parentIdx)
	if isLeft {
		parent.Left = n
	} else {
		parent.Right = n
	}
	t.Count++
}

func (t *CBT) findNode(n *TNode, v string) *TNode {
	if n == nil {
		return nil
	}
	if n.Data == v {
		return n
	}
	l := t.findNode(n.Left, v)
	if l != nil {
		return l
	}
	return t.findNode(n.Right, v)
}

func (t *CBT) Find(v string) bool {
	return t.findNode(t.Root, v) != nil
}

// По заданию для CBT мы удаление не делаем, а только проверку на Complete
func (t *CBT) IsComplete() bool {
	return true // По построению дерево всегда Complete
}

func (t *CBT) height(n *TNode) int {
	if n == nil {
		return 0
	}
	l := t.height(n.Left)
	r := t.height(n.Right)
	if l > r {
		return 1 + l
	}
	return 1 + r
}

func (t *CBT) Print() {
	var levels []string
	for i := 0; i < t.Count; i++ {
		node := t.nodeByIndex(t.Root, 0, i)
		levels = append(levels, node.Data)
	}
	fmt.Printf("Уровни:   [%s]\n", strings.Join(levels, ", "))
	fmt.Printf("Complete: %s\n", map[bool]string{true: "ДА", false: "НЕТ"}[t.IsComplete()])

	if t.Root == nil {
		fmt.Println("(дерево пусто)")
		return
	}
	fmt.Println("Дерево:")

	h := t.height(t.Root)
	maxWidth := (1 << h) - 1

	cell := 3
	for i := 0; i < t.Count; i++ {
		nd := t.nodeByIndex(t.Root, 0, i)
		if len(nd.Data)+2 > cell {
			cell = len(nd.Data) + 2
		}
	}
	totalWidth := maxWidth * cell

	for level := 0; level < h; level++ {
		line := strings.Split(strings.Repeat(" ", totalWidth), "")
		nodesOnLevel := 1 << level
		spanPerNode := totalWidth / nodesOnLevel

		for j := 0; j < nodesOnLevel; j++ {
			idx := (1 << level) - 1 + j
			if idx >= t.Count {
				break
			}
			node := t.nodeByIndex(t.Root, 0, idx)

			center := j*spanPerNode + spanPerNode/2
			start := center - len(node.Data)/2
			if start < 0 {
				start = 0
			}
			for k, ch := range node.Data {
				if start+k < totalWidth {
					line[start+k] = string(ch)
				}
			}
		}
		fmt.Println(strings.TrimRight(strings.Join(line, ""), " "))

		if level+1 < h {
			branches := strings.Split(strings.Repeat(" ", totalWidth), "")
			for j := 0; j < nodesOnLevel; j++ {
				idx := (1 << level) - 1 + j
				if idx >= t.Count {
					break
				}
				node := t.nodeByIndex(t.Root, 0, idx)
				if node.Left == nil && node.Right == nil {
					continue
				}
				center := j*spanPerNode + spanPerNode/2

				if node.Left != nil {
					lc := center - spanPerNode/2
					for k := lc; k < center; k++ {
						if k >= 0 && k < totalWidth {
							branches[k] = "-"
						}
					}
					if lc >= 0 && lc < totalWidth {
						branches[lc] = "+"
					}
				}
				if node.Right != nil {
					rc := center + spanPerNode/2
					for k := center + 1; k <= rc; k++ {
						if k >= 0 && k < totalWidth {
							branches[k] = "-"
						}
					}
					if rc >= 0 && rc < totalWidth {
						branches[rc] = "+"
					}
				}
				if center >= 0 && center < totalWidth {
					branches[center] = "+"
				}
			}
			fmt.Println(strings.TrimRight(strings.Join(branches, ""), " "))
		}
	}
}

func (t *CBT) PrintDetails() {
	fmt.Println("Подробный вывод дерева:")
	var rec func(*TNode, int)
	rec = func(n *TNode, level int) {
		if n == nil {
			return
		}
		fmt.Printf("  уровень %d addr=%p data=\"%s\" left=%p right=%p\n", level, n, n.Data, n.Left, n.Right)
		rec(n.Left, level+1)
		rec(n.Right, level+1)
	}
	rec(t.Root, 0)
}

func (t *CBT) Save() string {
	var res []string
	for i := 0; i < t.Count; i++ {
		nd := t.nodeByIndex(t.Root, 0, i)
		res = append(res, nd.Data)
	}
	return strings.Join(res, " ")
}

func (t *CBT) Load(parts []string) {
	for _, v := range parts {
		t.Insert(v)
	}
}
