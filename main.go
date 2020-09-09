package main

import (
	"bufio"
	"fmt"
	_board "noughts-crosses/board"
	_game "noughts-crosses/game"
	_players "noughts-crosses/players"
	"os"
	"strconv"
	"strings"
)

func main() {

	playerOne := _players.New("Guilherme", 'X')
	playerTwo := _players.New("Guilherme The Second", 'Z')

	game := _game.New(3, *playerOne, *playerTwo)

	for !game.IsGameOver() {
		game.DrawBoard()
		fmt.Printf("\n \n It's %s's turn! (ROW, COL): ", game.GetPlayerInTurn().Name())

		for {
			success, err := game.Play(ReadInput())

			if success {
				break
			}

			fmt.Print("Your play failed: ", err, " Please, try again: ")
		}
	}
}

func ReadInput() _board.Coordinate {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.ReplaceAll(input, "\n", "")

	coordinate := strings.Split(input, ",")

	coordX, _ := strconv.Atoi(coordinate[0])
	coordY, _ := strconv.Atoi(coordinate[1])

	return _board.Coordinate{X: coordX, Y: coordY}
}
