package game

import (
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
