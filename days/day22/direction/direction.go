package direction

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

func (d Direction) TurnLeft() Direction {
	return (d + 3) % 4
}

func (d Direction) TurnRight() Direction {
	return (d + 1) % 4
}

func (d Direction) Reverse() Direction {
	return (d + 2) % 4
}

func (d Direction) String() string {
	switch d {
	case Right:
		return "right"
	case Down:
		return "down"
	case Left:
		return "left"
	case Up:
		return "up"
	}
	panic("invalid direction")
}