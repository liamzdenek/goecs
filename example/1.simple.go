package main

import (
	"./.."
	"strconv"
)

type MyNode struct{}

func (mn *MyNode) Satisfies(e *ecs.Entity) bool {
	if !e.HasComponent(&MyComponent{}) {
		println("Does not satisfy")
		return false
	}
	println("Satisfies")
	return true
}

func (mn *MyNode) Tick(e *ecs.Entity) {
	mycomponent := e.GetComponent(&MyComponent{}).(*MyComponent)
	mycomponent.Counter++
	println("Hello #" + strconv.Itoa(mycomponent.Counter) + " from tick: " + mycomponent.EmbeddedData)
}

type MyComponent struct {
	EmbeddedData string
	Counter      int
}

func main() {
	s := ecs.NewSystem()
	s.AddNode(&MyNode{})

	e := ecs.NewEntity()
	e.Add(&MyComponent{"This data is hidden from the node", 0})

	s.AddEntity(e)
	s.Tick() // prints

	e.Destroy()
	s.Tick() // does not print

	s.AddEntity(e)
	s.Tick() // prints
}
