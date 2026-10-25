package main

import (
	"fmt"
	"lab1-go/structs"
	"os"
	"strings"
)

func main() {
	filename := ""
	query := ""

	// 1. РАЗБОР АРГУМЕНТОВ КОМАНДНОЙ СТРОКИ
	for i := 1; i < len(os.Args); i++ {
		a := os.Args[i]
		if a == "--file" && i+1 < len(os.Args) {
			filename = os.Args[i+1]
			i++
		} else if a == "--query" && i+1 < len(os.Args) {
			query = os.Args[i+1]
			i++
		} else if !strings.HasPrefix(a, "-") && filename == "" {
			filename = a
		}
	}

	if filename == "" || query == "" {
		fmt.Println("Использование: ./dbms --file data.db --query 'КОМАНДА'")
		return
	}

	// 2. ЗАГРУЗКА ДАННЫХ ИЗ ФАЙЛА=
	structs.LoadFile(filename)

	// 3. РАЗБОР КОМАНДЫ
	p := strings.Fields(query)
	if len(p) == 0 {
		return
	}

	c := p[0]
	changed := false
	// 4. ОБРАБОТКА КОМАНД

	// МАССИВ (M) 
	if c == "MPUSH" && len(p) == 3 {
		if structs.Arrays[p[1]] == nil {
			structs.Arrays[p[1]] = structs.NewArray()
		}
		structs.Arrays[p[1]].PushBack(p[2])
		changed = true
	} else if c == "MPUSH" && len(p) == 4 {
		if structs.Arrays[p[1]] == nil {
			structs.Arrays[p[1]] = structs.NewArray()
		}
		structs.Arrays[p[1]].InsertAt(structs.ParseIndex(p[2]), p[3])
		changed = true
	} else if c == "MGET" && len(p) == 3 {
		if structs.Arrays[p[1]] != nil {
			fmt.Println(structs.Arrays[p[1]].Get(structs.ParseIndex(p[2])))
		}
	} else if c == "MSET" && len(p) == 4 {
		if structs.Arrays[p[1]] != nil {
			structs.Arrays[p[1]].Set(structs.ParseIndex(p[2]), p[3])
		}
		changed = true
	} else if c == "MDEL" && len(p) == 3 {
		if structs.Arrays[p[1]] != nil {
			structs.Arrays[p[1]].RemoveAt(structs.ParseIndex(p[2]))
		}
		changed = true
	} else if c == "MLEN" && len(p) == 2 {
		if structs.Arrays[p[1]] != nil {
			fmt.Println(structs.Arrays[p[1]].Length())
		}

	//  ОДНОСВЯЗНЫЙ СПИСОК (F) 
	} else if c == "FPUSH" && len(p) == 4 {
		if structs.SLists[p[1]] == nil {
			structs.SLists[p[1]] = structs.NewSList()
		}
		L := structs.SLists[p[1]]
		if p[2] == "head" {
			L.PushHead(p[3])
		} else if p[2] == "tail" {
			L.PushTail(p[3])
		}
		changed = true
	} else if c == "FPUSH" && len(p) == 5 {
		if structs.SLists[p[1]] == nil {
			structs.SLists[p[1]] = structs.NewSList()
		}
		L := structs.SLists[p[1]]
		if p[2] == "after" {
			L.PushAfter(p[3], p[4])
		} else if p[2] == "before" {
			L.PushBefore(p[3], p[4])
		}
		changed = true
	} else if c == "FDEL" && len(p) == 3 {
		if structs.SLists[p[1]] != nil {
			L := structs.SLists[p[1]]
			if p[2] == "head" {
				L.PopHead()
			} else if p[2] == "tail" {
				L.PopTail()
			} else {
				L.RemoveByValue(p[2])
			}
		}
		changed = true
	} else if c == "FDEL" && len(p) == 4 {
		if structs.SLists[p[1]] != nil {
			L := structs.SLists[p[1]]
			if p[2] == "after" {
				L.PopAfter(p[3])
			} else if p[2] == "before" {
				L.PopBefore(p[3])
			}
		}
		changed = true
	} else if c == "FGET" && len(p) == 3 {
		if structs.SLists[p[1]] != nil {
			L := structs.SLists[p[1]]
			if p[2] == "forward" {
				L.PrintForward()
			} else if p[2] == "reverse" {
				L.PrintReverse()
			} else if p[2] == "details" {
				L.PrintDetails()
			} else {
				if L.Find(p[2]) {
					fmt.Println("ИСТИНА")
				} else {
					fmt.Println("ЛОЖЬ")
				}
			}
		}

	//  ДВУСВЯЗНЫЙ СПИСОК (L) 
	} else if c == "LPUSH" && len(p) == 4 {
		if structs.DLists[p[1]] == nil {
			structs.DLists[p[1]] = structs.NewDList()
		}
		L := structs.DLists[p[1]]
		if p[2] == "head" {
			L.PushHead(p[3])
		} else if p[2] == "tail" {
			L.PushTail(p[3])
		}
		changed = true
	} else if c == "LPUSH" && len(p) == 5 {
		if structs.DLists[p[1]] == nil {
			structs.DLists[p[1]] = structs.NewDList()
		}
		L := structs.DLists[p[1]]
		if p[2] == "after" {
			L.PushAfter(p[3], p[4])
		} else if p[2] == "before" {
			L.PushBefore(p[3], p[4])
		}
		changed = true
	} else if c == "LDEL" && len(p) == 3 {
		if structs.DLists[p[1]] != nil {
			L := structs.DLists[p[1]]
			if p[2] == "head" {
				L.PopHead()
			} else if p[2] == "tail" {
				L.PopTail()
			} else {
				L.RemoveByValue(p[2])
			}
		}
		changed = true
	} else if c == "LDEL" && len(p) == 4 {
		if structs.DLists[p[1]] != nil {
			L := structs.DLists[p[1]]
			if p[2] == "after" {
				L.PopAfter(p[3])
			} else if p[2] == "before" {
				L.PopBefore(p[3])
			}
		}
		changed = true
	} else if c == "LGET" && len(p) == 3 {
		if structs.DLists[p[1]] != nil {
			L := structs.DLists[p[1]]
			if p[2] == "forward" {
				L.PrintForward()
			} else if p[2] == "backward" {
				L.PrintBackward()
			} else if p[2] == "details" {
				L.PrintDetails()
			} else {
				if L.Find(p[2]) {
					fmt.Println("ИСТИНА")
				} else {
					fmt.Println("ЛОЖЬ")
				}
			}
		}

	//  СТЕК (S)
	} else if c == "SPUSH" && len(p) == 3 {
		if structs.Stacks[p[1]] == nil {
			structs.Stacks[p[1]] = structs.NewStack()
		}
		structs.Stacks[p[1]].Push(p[2])
		changed = true
	} else if c == "SPOP" && len(p) == 2 {
		if structs.Stacks[p[1]] != nil {
			fmt.Println(structs.Stacks[p[1]].Pop())
		}
		changed = true
	} else if c == "SGET" && len(p) == 2 {
		if structs.Stacks[p[1]] != nil {
			fmt.Println(structs.Stacks[p[1]].TopVal())
		}

    // ОЧЕРЕДЬ

	} else if c == "QPUSH" && len(p) == 3 {
		if structs.Queues[p[1]] == nil {
			structs.Queues[p[1]] = structs.NewQueue()
		}
		structs.Queues[p[1]].Push(p[2])
		changed = true
	} else if c == "QPOP" && len(p) == 2 {
		if structs.Queues[p[1]] != nil {
			fmt.Println(structs.Queues[p[1]].Pop())
		}
		changed = true
	} else if c == "QGET" && len(p) == 2 {
		if structs.Queues[p[1]] != nil {
			fmt.Println(structs.Queues[p[1]].FrontVal())
		}

	// ДЕРЕВО (T)
	} else if c == "TINSERT" && len(p) == 3 {
		if structs.Trees[p[1]] == nil {
			structs.Trees[p[1]] = structs.NewCBT()
		}
		structs.Trees[p[1]].Insert(p[2])
		changed = true
	} else if c == "TGET" && len(p) == 3 {
		if structs.Trees[p[1]] != nil {
			if structs.Trees[p[1]].Find(p[2]) {
				fmt.Println("ИСТИНА")
			} else {
				fmt.Println("ЛОЖЬ")
			}
		}
	} else if c == "TCHECK" && len(p) == 3 {
		if structs.Trees[p[1]] != nil && p[2] == "complete" {
			if structs.Trees[p[1]].IsComplete() {
				fmt.Println("ИСТИНА")
			} else {
				fmt.Println("ЛОЖЬ")
			}
		}

	//  (PRINT / DETAILS) 
	} else if c == "PRINT" && len(p) == 2 {
		n := p[1]
		if structs.Arrays[n] != nil {
			structs.Arrays[n].Print()
		} else if structs.SLists[n] != nil {
			structs.SLists[n].PrintForward()
		} else if structs.DLists[n] != nil {
			structs.DLists[n].PrintForward()
		} else if structs.Stacks[n] != nil {
			structs.Stacks[n].Print()
		} else if structs.Queues[n] != nil {
			structs.Queues[n].Print()
		} else if structs.Trees[n] != nil {
			structs.Trees[n].Print()
		} else {
			fmt.Println("не найдено")
		}
	} else if c == "DETAILS" && len(p) == 2 {
		n := p[1]
		if structs.Arrays[n] != nil {
			structs.Arrays[n].PrintDetails()
		} else if structs.SLists[n] != nil {
			structs.SLists[n].PrintDetails()
		} else if structs.DLists[n] != nil {
			structs.DLists[n].PrintDetails()
		} else if structs.Stacks[n] != nil {
			structs.Stacks[n].PrintDetails()
		} else if structs.Queues[n] != nil {
			structs.Queues[n].PrintDetails()
		} else if structs.Trees[n] != nil {
			structs.Trees[n].PrintDetails()
		} else {
			fmt.Println("не найдено")
		}

	} else {
		fmt.Println("Неизвестная команда:", c)
		return
	}

	// 5. СОХРАНЕНИЕ
	if changed {
		structs.SaveFile(filename)
	}
}
