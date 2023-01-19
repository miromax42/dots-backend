package game

import (
	"github.com/cockroachdb/errors"
	"strconv"
	"strings"
)

type Board [10][10]Tile

func NewBoard() *Board {
	b := &Board{}
	b[0][0] = Tile{
		State: active,
		Side:  red,
	}

	b[9][9] = Tile{
		State: active,
		Side:  blue,
	}

	return b
}

func (b *Board) Render() string {
	buffer := strings.Builder{}

	buffer.WriteString(" 0123456789\n")

	for i := range b {
		buffer.WriteString(strconv.Itoa(i))

		for j := range b[i] {
			buffer.WriteString(b[i][j].String())
		}

		buffer.WriteRune('\n')
	}

	return buffer.String()
}

func (b *Board) Set(i, j int, t Tile) error {
	iValid := 0 <= i && i < 10
	jValid := 0 <= j && j < 10

	if !iValid || !jValid {
		return errors.Wrapf(ErrCoords, "i=%d,j=%d %v %v", i, j, iValid, jValid)
	}

	if !b[i][j].Available(t) {
		return errors.Wrapf(ErrMove, "%s->%s", b[i][j], t)
	}

	b[i][j] = t

	return nil
}
