Go Entity-Component-System
=

This is a tiny framework intended for reference when developing an entity-component-system (hereafter ECS) in Go.

This framework is far from efficient (maps are used in many cases where arrays should, out of laziness)

Use
-
ECS frameworks are particularly useful when you need to load and unload properties from many entities at runtime, and increases code modularity.

How ECS works
-
This ECS framework is made up of four types: Entity, Component, Node, and System.

A Component is a struct, eg, "Velocity". There may be methods on these structs, but they should be restricted to performing operations purely on the data within that struct.

An Entity is an unnamed collection of Components. These components can be added and removed from the entity at runtime. For example, "Velocity", "Car", "Position"

A Node defines rules for Entities depending on what components are loaded for them. eg, "Car + Position" might activate NodeCollidable

A System keeps track of all Entities and Nodes, and is responsible for triggering Node logic.
