package swarm

type Swarm struct {
	Elements []Mover

	CurrentFrame int64
}

// Start the simulation
func (s *Swarm) Start() {
	for i := 0; i < 100; i++ {
		s.Step()
	}
}

// Step one frame in the simulation
func (s *Swarm) Step() {
	for _, mover := range s.Elements {
		mover.Move()
		s.CurrentFrame++
	}
}
