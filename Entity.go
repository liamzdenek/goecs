package ecs

import (
	"reflect"
)

type Entity struct {
	system *System

	components []Component

	changed bool
}

func NewEntity() *Entity {
	return &Entity{}
}

func (e *Entity) GetSystem() *System {
	return e.system
}

func (e *Entity) GetComponent(c Component) Component {
	kind := reflect.TypeOf(c)
	for _, tc := range e.components {
		tkind := reflect.TypeOf(tc)
		if kind == tkind {
			return tc
		}
	}
	return nil
}

func (e *Entity) Destroy() {
	e.changed = true
	e.system.unregisterEntity(e)
	e.system = nil
}

func (e *Entity) HasComponent(c Component) bool {
	return e.GetComponent(c) != nil
}

func (e *Entity) Add(c Component) {
	e.changed = true
	e.components = append(e.components, c)
	e.recalculateNodes()
}

func (e *Entity) recalculateNodes() {
	// this is hax
	if e.system != nil {
		system := e.system
		e.Destroy()
		system.AddEntity(e)
	}
}

func (e *Entity) Remove(c Component) {
	e.changed = true
	tc := e.GetComponent(c)
	for i, tc2 := range e.components {
		if tc == tc2 {
			if i == 0 {
				e.components = e.components[1:]
			} else if i == len(e.components) {
				e.components = e.components[:i]
			} else {
				e.components = append(e.components[:i], e.components[i+1:]...)
			}
			break
		}
	}
	e.recalculateNodes()
}

func (e *Entity) ListComponents() string {
	str := ""
	for _, c := range e.components {
		str = str + reflect.TypeOf(c).String() + " "
	}
	return str
}
