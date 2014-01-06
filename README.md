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

a Component is a struct, eg, "Velocity"

an Entity is an unnamed collection of components. for example, "Velocity", "Car", "Weight"

A node defines rules for entities depending on what components are loaded for them. eg, "Car + Velocity" might activate collision physics

A system keeps track of all entities and nodes, and is responsible for triggering node logic.
