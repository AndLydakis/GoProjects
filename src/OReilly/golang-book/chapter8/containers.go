package main

import (
	"fmt"
	"container/list"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

type ByAge []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps ByAge) Len() int {
	return len(ps)
}

func (ps ByAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}

func (ps ByAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	var x list.List
	x.PushBack(3)
	x.PushBack(2)
	x.PushBack(1)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	kids := []Person{
		{"Jill", 10},
		{"Pearson", 8},
	}

	sort.Sort(ByName(kids))
	fmt.Println(kids)
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}
