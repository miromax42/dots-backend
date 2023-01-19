package game

type State int8

const (
	empty State = iota
	active
	grabbed
)

func (s State) String() string {
	return [...]string{
		" ",
		"o",
		"x",
	}[s]
}

type Side int8

const (
	red Side = iota + 1
	blue
)

func (s Side) String() string {
	return [...]string{
		"\u001b[0m",
		"\u001b[31m",
		"\u001b[34m",
	}[s]
}

type Tile struct {
	State State
	Side  Side
}

func (t Tile) String() string {
	return t.Side.String() + t.State.String() + Side(0).String()
}

func (t Tile) Available(next Tile) bool {
	switch {
	case t.State == empty:
		return next.State == active
	case t.State == active:
		return next.State == grabbed && t.State != next.State
	default:
		return false
	}
}
