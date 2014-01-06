package ecs

import (
	"reflect"
)

type System struct {
	Entities []*Entity
	Nodes    []Node
}

func NewSystem() *System {
	return &System{}
}

func (s *System) AddNode(n Node) {
	kind := reflect.TypeOf(n).Kind()

	for _, tn := range s.Nodes {
		tkind := reflect.TypeOf(tn).Kind()
		if kind == tkind {
			panic("Attempting to add the same node (type: '" + kind.String() + "') multiple times to a single system")
		}
	}

	s.Nodes = append(s.Nodes, n)
}

func (s *System) AddEntity(e *Entity) {
	s.Entities = append(s.Entities, e)
}

func (s *System) Tick() {
	for _, n := range s.Nodes {
		for _, e := range s.Entities {
			if n.Satisfies(e) {
				n.Tick(e);
			}
		}
	}
}
