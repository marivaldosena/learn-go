package main

import (
	"container/list"
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func usingContainerList() {
	var x list.List

	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}

func orderListByName(people []Person) {
	sort.Sort(ByName(people))
}

func main() {
	usingContainerList()

	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
		{"Joe", 8},
		{"Julie", 7},
	}

	orderList(kids)
	fmt.Println(kids)
}
