package board

type Coordinate struct {
	X, Y int
}

func (c *Coordinate) isValid(boardSize int) bool {
	return c.X < boardSize && c.X >= 0 && c.Y < boardSize && c.Y >= 0
}
