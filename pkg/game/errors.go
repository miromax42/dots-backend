package game

import "github.com/cockroachdb/errors"

var (
	ErrCoords  = errors.New("coordinates out of board")
	ErrMove    = errors.New("unavailable move")
	ErrRequest = errors.New("bad request")
)
