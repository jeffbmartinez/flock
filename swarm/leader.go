package swarm

type Leader struct {
	*Element
}

func (l *Leader) Move() {
}

func NewLeader(startX, startY, direction, speed, personalRadius float64) *Leader {
	return &Leader{
		Element: NewElement(startX, startY, direction, speed, personalRadius),
	}
}
