package ecs

type Node interface {
	// deterministic function to assess whether an entity satisfies a node
	Satisfies(e *Entity) bool

	Tick(e *Entity)
}
