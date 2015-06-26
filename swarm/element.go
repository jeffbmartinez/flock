package swarm

import (
	"github.com/jeffbmartinez/swarm/physics"
)

type Element struct {
	physics.Velocity
	physics.Point

	PersonalRadius float64
}

type Mover interface {
	Move()
}

func NewElement(startX, startY, direction, speed, personalRadius float64) *Element {
	return &Element{
		Velocity: physics.Velocity{
			Direction: direction,
			Speed:     speed,
		},
		Point: physics.Point{
			X: startX,
			Y: startY,
		},
		PersonalRadius: personalRadius,
	}
}
