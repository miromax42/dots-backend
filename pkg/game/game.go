package game

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

func (g *Game) Move() error {
	return nil
}
