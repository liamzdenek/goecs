package ecs;

import (
	"reflect"
)

type Entity struct {
	system *System

	components []Component
}

func NewEntity() *Entity {
	return &Entity{};
}

func(e *Entity) GetComponent(c Component) Component {
	kind := reflect.TypeOf(c).Kind();
	for _, tc := range e.components {
		tkind := reflect.TypeOf(tc).Kind();
		if kind == tkind {
			return tc;
		}
	}
	return nil;
}

func(e *Entity) HasComponent(c Component) bool {
	return e.GetComponent(c) != nil;
}

func(e *Entity) Add(c Component) {
	e.components = append(e.components, c);
}

func(e *Entity) Remove(c Component) bool {return false}
