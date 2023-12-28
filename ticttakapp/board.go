package main

import (
	"errors"
)

// Stałe reprezentujące graczy i puste pole.
const (
	PlayerX = "X"
	PlayerO = "O"
	Empty   = " "
)

// Board reprezentuje planszę gry w kółko i krzyżyk.
type Board [3][3]string

// NewBoard tworzy nową, pustą planszę gry.
func NewBoard() Board {
	return Board{{Empty, Empty, Empty}, {Empty, Empty, Empty}, {Empty, Empty, Empty}}
}

// PlaceMark umieszcza znak na planszy.
func (b *Board) PlaceMark(x, y int, mark string) error {
	if x < 0 || y < 0 || x >= 3 || y >= 3 {
		return errors.New("poza zakresem planszy")
	}
	if b[x][y] != Empty {
		return errors.New("pole już zajęte")
	}
	b[x][y] = mark
	return nil
}

// IsFull sprawdza, czy plansza jest pełna.
func (b *Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == Empty {
				return false
			}
		}
	}
	return true
}
