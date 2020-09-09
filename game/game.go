package game

import (
	"math/rand"
	_board "noughts-crosses/board"
	_players "noughts-crosses/players"
	"time"
)

type Game struct {
	players         []_players.Player
	turnPlayerIndex int
	board           *_board.Board
}

func New(boardSize int, players ..._players.Player) *Game {
	rand.Seed(time.Now().UnixNano())

	game := Game{}

	game.board = _board.New(boardSize)
	game.players = players
	game.turnPlayerIndex = rand.Intn(len(players))

	return &game
}

func (g *Game) GetPlayerInTurn() *_players.Player {
	return &g.players[g.turnPlayerIndex]
}

func (g *Game) SetNextPlayerTurn() {
	g.turnPlayerIndex += 1

	if g.turnPlayerIndex >= len(g.players) {
		g.turnPlayerIndex = 0
	}
}

func (g *Game) Play(coord _board.Coordinate) (bool, error) {
	success, err := g.board.SetFieldValue(coord, g.GetPlayerInTurn().Symbol())

	if success {
		g.SetNextPlayerTurn()

		return true, nil
	}

	return false, err
}

func (g *Game) DrawBoard() {
	g.board.Draw()
}

func (g *Game) IsGameOver() bool {
	return false
}
