package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Board Board
	Turn  string
}

func NewGame() *Game {
	return &Game{
		Board: NewBoard(),
		Turn:  PlayerX,
	}
}

func (g *Game) PlayGame() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Plansza gry:\n")
		g.printBoard()

		if g.CheckWin() {
			fmt.Printf("Gracz %s wygrywa!\n", g.Turn)
			break
		}

		if g.Board.IsFull() {
			fmt.Println("Gra zakończona remisem!")
			break
		}

		fmt.Printf("Ruch gracza %s. Wprowadź pozycję (x y): ", g.Turn)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		coords := strings.Split(input, " ")

		if len(coords) != 2 {
			fmt.Println("Nieprawidłowy format. Wprowadź ponownie.")
			continue
		}

		x, errX := strconv.Atoi(coords[0])
		y, errY := strconv.Atoi(coords[1])

		if errX != nil || errY != nil || !g.isValidPosition(x, y) {
			fmt.Println("Nieprawidłowe współrzędne. Wprowadź ponownie.")
			continue
		}

		err := g.Board.PlaceMark(x, y, g.Turn)
		if err != nil {
			fmt.Println("Błąd:", err)
			continue
		}

		if g.CheckWin() {
			fmt.Printf("Gracz %s wygrywa!\n", g.Turn)
			break
		}

		g.ChangeTurn()
	}
}

func (g *Game) CheckWin() bool {
	// Sprawdzanie wierszy i kolumn
	for i := 0; i < 3; i++ {
		if g.Board[i][0] == g.Turn && g.Board[i][1] == g.Turn && g.Board[i][2] == g.Turn {
			return true
		}
		if g.Board[0][i] == g.Turn && g.Board[1][i] == g.Turn && g.Board[2][i] == g.Turn {
			return true
		}
	}

	// Sprawdzanie przekątnych
	if g.Board[0][0] == g.Turn && g.Board[1][1] == g.Turn && g.Board[2][2] == g.Turn {
		return true
	}
	if g.Board[0][2] == g.Turn && g.Board[1][1] == g.Turn && g.Board[2][0] == g.Turn {
		return true
	}

	return false
}

func (g *Game) ChangeTurn() {
	if g.Turn == PlayerX {
		g.Turn = PlayerO
	} else {
		g.Turn = PlayerX
	}
}

func (g *Game) printBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(" %s ", g.Board[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		if i < 2 {
			fmt.Println("\n-----------")
		} else {
			fmt.Println()
		}
	}
}

func (g *Game) isValidPosition(x, y int) bool {
	return x >= 0 && x < 3 && y >= 0 && y < 3
}
