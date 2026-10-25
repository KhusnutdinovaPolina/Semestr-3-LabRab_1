package structs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Глобальные хранилища
var (
	Arrays = make(map[string]*Array)
	SLists = make(map[string]*SList)
	DLists = make(map[string]*DList)
	Stacks = make(map[string]*Stack)
	Queues = make(map[string]*Queue)
	Trees  = make(map[string]*CBT)
)

func SaveFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	for name, arr := range Arrays {
		fmt.Fprintf(f, "M %s: %s\n", name, arr.Save())
	}
	for name, l := range SLists {
		fmt.Fprintf(f, "F %s: %s\n", name, l.Save())
	}
	for name, l := range DLists {
		fmt.Fprintf(f, "L %s: %s\n", name, l.Save())
	}
	for name, s := range Stacks {
		fmt.Fprintf(f, "S %s: %s\n", name, s.Save())
	}
	for name, q := range Queues {
		fmt.Fprintf(f, "Q %s: %s\n", name, q.Save())
	}
	for name, t := range Trees {
		fmt.Fprintf(f, "T %s: %s\n", name, t.Save())
	}
}

func LoadFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		colon := strings.Index(line, ":")
		if colon == -1 {
			continue
		}

		head := strings.Fields(line[:colon])
		if len(head) < 2 {
			continue
		}
		typ, name := head[0], head[1]

		valsStr := strings.TrimSpace(line[colon+1:])
		var vals []string
		if valsStr != "" {
			vals = strings.Fields(valsStr)
		}

		switch typ {
		case "M":
			if Arrays[name] == nil { Arrays[name] = NewArray() }
			Arrays[name].Load(vals)
		case "F":
			if SLists[name] == nil { SLists[name] = NewSList() }
			SLists[name].Load(vals)
		case "L":
			if DLists[name] == nil { DLists[name] = NewDList() }
			DLists[name].Load(vals)
		case "S":
			if Stacks[name] == nil { Stacks[name] = NewStack() }
			Stacks[name].Load(vals)
		case "Q":
			if Queues[name] == nil { Queues[name] = NewQueue() }
			Queues[name].Load(vals)
		case "T":
			if Trees[name] == nil { Trees[name] = NewCBT() }
			Trees[name].Load(vals)
		}
	}
}
