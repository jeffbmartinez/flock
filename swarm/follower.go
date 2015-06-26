package swarm

type Follower struct {
	*Element
	Leader *Leader
}

func (f *Follower) Move() {
}

func NewFollower(startX, startY, direction, speed, personalRadius float64, leader *Leader) *Follower {
	return &Follower{
		Element: NewElement(startX, startY, direction, speed, personalRadius),
		Leader:  leader,
	}
}
