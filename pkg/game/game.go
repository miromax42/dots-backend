package game

import (
	"github.com/cockroachdb/errors"
	"strconv"
	"strings"
)

type Game struct {
	board *Board
}

func New() *Game {
	return &Game{
		board: NewBoard(),
	}
}

func (g *Game) Render() string {
	return g.board.Render()
}

func (g *Game) Move(c string) error {
	commands := strings.Split(c, " ")

	if len(commands) != 4 {
		return errors.Wrap(ErrRequest, "bad args count(need 4)")
	}

	i, err := strconv.Atoi(commands[0])
	if err != nil {
		return errors.Wrap(ErrRequest, "i is not a num")
	}

	j, err := strconv.Atoi(commands[1])
	if err != nil {
		return errors.Wrap(ErrRequest, "j is not a num")
	}

	var tile Tile
	switch commands[2] {
	case "1":
		tile.State = active
	case "2":
		tile.State = grabbed
	default:
		return errors.Wrapf(ErrRequest, "state %q unknown", commands[2])
	}

	switch commands[3] {
	case "r":
		tile.Side = red
	case "b":
		tile.Side = blue
	default:
		return errors.Wrapf(ErrRequest, "side %q unknown", commands[3])
	}

	return g.board.Set(i, j, tile)
}
