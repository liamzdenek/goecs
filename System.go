package ecs

import (
	"reflect"
)

type System struct {
	Entities map[*Entity][]Node
	Nodes    map[Node][]*Entity
}

func NewSystem() *System {
	return &System{
		Nodes:    make(map[Node][]*Entity),
		Entities: make(map[*Entity][]Node),
	}
}

func (s *System) AddNode(n Node) {
	kind := reflect.TypeOf(n)

	for tn, _ := range s.Nodes {
		tkind := reflect.TypeOf(tn)
		if kind == tkind {
			panic("Attempting to add the same node (type: '" + kind.String() + "') multiple times to a single system")
		}
	}
	s.Nodes[n] = make([]*Entity, 0)
	for e := range s.Entities {
		if n.Satisfies(e) {
			s.register(n, e)
		}
	}
}

func (s *System) AddEntity(e *Entity) {
	if e.system != nil {
		panic("Attempting to add the same entity to multiple systems")
	}

	e.system = s

	s.Entities[e] = make([]Node, 0)

	for n := range s.Nodes {
		if n.Satisfies(e) {
			s.register(n, e)
		}
	}
}

func (s *System) GetEntitiesByNode(n Node) []*Entity {
	return s.Nodes[n]
}

func (s *System) GetNodesByEntity(e *Entity) []Node {
	return s.Entities[e]
}

func (s *System) register(n Node, e *Entity) {
	s.Nodes[n] = append(s.Nodes[n], e)
	s.Entities[e] = append(s.Entities[e], n)
}

func (s *System) unregisterNode(n Node) {
	for _, e := range s.Nodes[n] {
		n_list := s.Entities[e]
		for i, tn := range n_list {
			if tn == n {
				if i == 0 {
					n_list = n_list[1:]
				} else if i == len(n_list) {
					n_list = n_list[:i]
				} else {
					n_list = append(n_list[:i], n_list[i+1:]...)
				}
				s.Entities[e] = n_list
				break
			}
		}
	}

	delete(s.Nodes, n)
}

func (s *System) unregisterEntity(e *Entity) {
	for _, n := range s.Entities[e] {
		e_list := s.Nodes[n]
		for i, te := range e_list {
			if te == e {
				if i == 0 {
					e_list = e_list[1:]
				} else if i == len(e_list) {
					e_list = e_list[:i]
				} else {
					e_list = append(e_list[:i], e_list[i+1:]...)
				}
				s.Nodes[n] = e_list
				break
			}
		}
	}

	delete(s.Entities, e)
}

func (s *System) Tick() {
	for n, e_list := range s.Nodes {
		for _, e := range e_list {
			if e.changed {
				e.changed = false;
				if n.Satisfies(e) {
					n.Tick(e)
				}
			} else {
				n.Tick(e)
			}
		}
	}
}
