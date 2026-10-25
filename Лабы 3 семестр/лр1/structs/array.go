package structs

import (
	"fmt"
	"strconv"
	"strings"
)

// Array — динамический массив строк
type Array struct {
	Data []string
}

func NewArray() *Array {
	return &Array{Data: []string{}}
}

func (a *Array) PushBack(v string) {
	a.Data = append(a.Data, v)
}

func (a *Array) InsertAt(index int, v string) {
	if index < 0 || index > len(a.Data) {
		return
	}
	a.Data = append(a.Data, "")
	copy(a.Data[index+1:], a.Data[index:])
	a.Data[index] = v
}

func (a *Array) Get(index int) string {
	if index < 0 || index >= len(a.Data) {
		return ""
	}
	return a.Data[index]
}

func (a *Array) Set(index int, v string) {
	if index >= 0 && index < len(a.Data) {
		a.Data[index] = v
	}
}

func (a *Array) RemoveAt(index int) {
	if index < 0 || index >= len(a.Data) {
		return
	}
	a.Data = append(a.Data[:index], a.Data[index+1:]...)
}

func (a *Array) Length() int {
	return len(a.Data)
}

func (a *Array) Print() {
	fmt.Printf("Массив:  [%s]\n", strings.Join(a.Data, ", "))
}

func (a *Array) PrintDetails() {
	fmt.Println("Массив (адреса ячеек):")
	for i, v := range a.Data {
		fmt.Printf("  [%d] addr=%p value=\"%s\"\n", i, &a.Data[i], v)
	}
}

func (a *Array) Save() string {
	return strings.Join(a.Data, " ")
}

func (a *Array) Load(parts []string) {
	a.Data = parts
}

// Хелпер для парсинга индексов (замена stoi)
func ParseIndex(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
