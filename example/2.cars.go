package main

import (
	"./.."
	"fmt"
	"math"
	"time"
)

type ComponentVelocity struct {
	DeltaX, DeltaY int
}

type ComponentPosition struct {
	X, Y int
}

type ComponentCar struct {
	Identifier   int
	HitboxRadius int
}

type ComponentFire struct{}

type NodeDriver struct{}

func (c *NodeDriver) Satisfies(e *ecs.Entity) bool {
	//fmt.Printf("Testing NodeDriver Satisfaction - %s\n", e.ListComponents())
	if !e.HasComponent(&ComponentCar{}) {
		return false
	}

	if !e.HasComponent(&ComponentPosition{}) {
		return false
	}

	if !e.HasComponent(&ComponentVelocity{}) {
		return false
	}
	return true
}

func (c *NodeDriver) Tick(e *ecs.Entity) {
	position := e.GetComponent(&ComponentPosition{}).(*ComponentPosition)
	velocity := e.GetComponent(&ComponentVelocity{}).(*ComponentVelocity)
	car := e.GetComponent(&ComponentCar{}).(*ComponentCar)

	position.X += velocity.DeltaX
	position.Y += velocity.DeltaY

	e_list := e.GetSystem().GetEntitiesByNode(c)
	for _, te := range e_list {
		if te == e {
			continue
		}
		tposition := te.GetComponent(&ComponentPosition{}).(*ComponentPosition)
		tvelocity := te.GetComponent(&ComponentVelocity{}).(*ComponentVelocity)
		tcar := te.GetComponent(&ComponentCar{}).(*ComponentCar)

		hitbox_dist := float64(tcar.HitboxRadius + car.HitboxRadius)
		actual_dist := math.Hypot(float64(position.X-tposition.X), float64(position.Y-tposition.Y))

		//fmt.Printf("DISTANCES: %d >= %d\n", hitbox_dist, actual_dist)

		if hitbox_dist >= actual_dist {
			fmt.Printf("Cars %d and %d have collided!\n", car.Identifier, tcar.Identifier)

			te.Remove(tvelocity)
			e.Remove(velocity)
			te.Add(&ComponentFire{})
			e.Add(&ComponentFire{})
		}
	}
}

type NodeAnnouncer struct{}

func (a *NodeAnnouncer) Satisfies(e *ecs.Entity) bool {
	if !e.HasComponent(&ComponentCar{}) {
		return false
	}

	if !e.HasComponent(&ComponentPosition{}) {
		return false
	}
	return true
}

func (a *NodeAnnouncer) Tick(e *ecs.Entity) {
	car := e.GetComponent(&ComponentCar{}).(*ComponentCar)
	position := e.GetComponent(&ComponentPosition{}).(*ComponentPosition)

	clause := ""

	if v, ok := e.GetComponent(&ComponentVelocity{}).(*ComponentVelocity); ok {
		clause = clause + fmt.Sprintf(". it is moving at %d, %d per tick", v.DeltaX, v.DeltaY)
	}

	if _, ok := e.GetComponent(&ComponentFire{}).(*ComponentFire); ok {
		clause = clause + ", and is on fire!"
	}

	fmt.Printf("Car #%d is at %d, %d%s\n", car.Identifier, position.X, position.Y, clause)
}

func main() {
	s := ecs.NewSystem()

	s.AddNode(&NodeDriver{})
	s.AddNode(&NodeAnnouncer{})

	e1 := ecs.NewEntity()
	e1.Add(&ComponentCar{Identifier: 69, HitboxRadius: 1})
	e1.Add(&ComponentPosition{-10, 0})

	e2 := ecs.NewEntity()
	e2.Add(&ComponentCar{Identifier: 1337, HitboxRadius: 1})
	e2.Add(&ComponentPosition{10, 0})

	s.AddEntity(e1)
	s.AddEntity(e2)

	counter := 0
	for {
		fmt.Printf("Tick!\n")
		s.Tick()
		counter++
		if counter == 5 {
			fmt.Printf("And they're off!\n")
			e1.Add(&ComponentVelocity{2, 0})
			e2.Add(&ComponentVelocity{-2, 0})
		}
		time.Sleep(time.Millisecond * 500)
	}
}
