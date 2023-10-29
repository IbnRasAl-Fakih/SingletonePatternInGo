package main

import (
	"fmt"
)

type VisitorCounter struct {
	count int
}

var instance *VisitorCounter

func GetInstance() *VisitorCounter {
	if instance == nil {
		instance = &VisitorCounter{}
		return instance
	}
	return instance
}

func (vc *VisitorCounter) Increment() {
	vc.count++
}

func (vc *VisitorCounter) Decrement() {
	vc.count--
}

func (vc *VisitorCounter) GetCount() int {
	return vc.count
}

func main() {
	counter := GetInstance()

	counter.Increment()
	counter.Increment()
	counter.Increment()

	count := counter.GetCount()

	fmt.Printf("Текущее количество посетителей: %d\n", count)

	anotherCounter := GetInstance()
	anotherCounter.Increment()
	anotherCounter.Decrement()
	anotherCounter.Decrement()

	anotherCount := anotherCounter.GetCount()

	fmt.Printf("Текущее количество посетителей: %d\n", anotherCount)
}
