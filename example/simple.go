package main

import (
	"./.."
)

type MyNode struct{
}

func (mn *MyNode) Satisfies(e *ecs.Entity) bool {
	if !e.HasComponent(&MyComponent{}) {
		println("Does not satisfy");
		return false
	}
	println("Satisfies");
	return true
}

func (mn *MyNode) Tick(e *ecs.Entity) {
	mycomponent := e.GetComponent(&MyComponent{}).(*MyComponent);
	println("Hello from tick: "+mycomponent.EmbeddedData);
}

type MyComponent struct{
	EmbeddedData string
}

func main() {
	s := ecs.NewSystem()
	s.AddNode(&MyNode{})

	e := ecs.NewEntity()
	e.Add(&MyComponent{"This data is hidden from the node"})

	s.AddEntity(e)
	s.Tick();

	e.Destroy();
	s.Tick();

	s.AddEntity(e);
	s.Tick();
}
