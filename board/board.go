package board

import (
	"errors"
	"fmt"
)

const EMPTY_VALUE = '-'

type Board struct {
	fields    [][]byte
	boardSize int
}

func New(boardSize int) *Board {
	board := Board{boardSize: boardSize}

	board.StartBoard()
	board.ClearBoard()

	return &board
}

func (b *Board) BoardSize() int {
	return b.boardSize
}

func (b *Board) Draw() {
	for _, rows := range b.fields {
		for _, cell := range rows {
			fmt.Printf("| %c |", cell)
		}

		fmt.Println("")
	}
}

func (b *Board) GetFieldValue(coord Coordinate) (byte, error) {
	if !coord.isValid(b.boardSize) {
		return EMPTY_VALUE, InvalidCoordinates()
	}

	return b.fields[coord.X][coord.Y], nil
}

func (b *Board) SetFieldValue(coord Coordinate, symbol byte) (bool, error) {
	if coord.isValid(b.boardSize) {
		if !b.isFieldEmpty(coord) {
			return false, CoordinateAlreadyFilledIn()
		}

		b.fields[coord.X][coord.Y] = symbol
		return true, nil
	}

	return false, InvalidCoordinates()
}

func (b *Board) ClearBoard() {
	for i, rows := range b.fields {
		for j, _ := range rows {
			b.fields[i][j] = EMPTY_VALUE
		}
	}
}

func (b *Board) isFieldEmpty(coord Coordinate) bool {
	value, _ := b.GetFieldValue(coord)
	return value == EMPTY_VALUE
}

func (b *Board) StartBoard() {
	b.fields = make([][]byte, b.boardSize)
	for i := range b.fields {
		b.fields[i] = make([]byte, b.boardSize)
	}
}

func InvalidCoordinates() error {
	return errors.New("Invalid coordinate!")
}

func CoordinateAlreadyFilledIn() error {
	return errors.New("Coordinate already filled in")
}
