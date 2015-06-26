/*
Simulating a swarm of something

Basic terminology:

swarm: The group of things
element: A single thing or object that is part of the swarm. For example,
	in a flock of birds, a single bird would be an element, the flock
	would be the swarm
leader: An element that other elements seem to follow, the followers'
	behavior is largely determined by that of the leader(s).
follower: An element that follows one or more leaders.
*/

package main

import (
	"fmt"
	"math/rand"

	"github.com/jeffbmartinez/swarm/swarm"
)

const (
	swarmSize = 10

	worldWidth  = 500
	worldHeight = 500

	leaderPersonalSpace   = 2.0
	followerPersonalSpace = 5.0
)

func main() {
	fmt.Printf("Swarm of size %v\n", swarmSize)

	sim := buildSwarmSimulation(swarmSize)
	sim.Start()
}

func buildSwarmSimulation(size int) *swarm.Swarm {
	sim := &swarm.Swarm{}

	startX := randomFloat64InRange(0.0, worldWidth)
	startY := randomFloat64InRange(0.0, worldHeight)
	startDirection := 0.0
	startSpeed := 2.0
	personalSpace := leaderPersonalSpace

	leader := swarm.NewLeader(startX, startY, startDirection, startSpeed, personalSpace)
	sim.Elements = append(sim.Elements, leader)

	for i := 0; i < size; i++ {
		startX = randomFloat64InRange(0.0, worldWidth)
		startY = randomFloat64InRange(0.0, worldHeight)
		startDirection = 0.0
		startSpeed = 2.0
		personalSpace = followerPersonalSpace

		follower := swarm.NewFollower(startX, startY, startDirection, startSpeed, personalSpace, leader)
		sim.Elements = append(sim.Elements, follower)
	}

	return sim
}

// Returns a random float64 in the range of [low, high)
// low is inclusive, high is exclusive
func randomFloat64InRange(low, high float64) float64 {
	randomRange := high - low

	return rand.Float64()*randomRange + low
}
