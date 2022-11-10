package main

import (
	"fmt"
)

// Erwartet einen Spieler-String und ein Spielfeld.
// Fragt den Spieler nach seinem Zug und führt diesen aus,
// sofern der Zug erlaubt und möglich ist.
// Ist der Zug ungültig, muss der Spieler erneut gefragt werden.
func MakeMove(player string, board [][]string) {
	fmt.Printf("Spieler %v, bitte wähle dein Spielfeld (1-9): ", player)
	var input int
	fmt.Scanln(&input)
	row := (input - 1) / 3
	col := (input - 1) % 3

	if MoveAllowed(board, row, col) {
		board[row][col] = player
		return
	}
	fmt.Println("Der Zug ist ungültig, bitte wiederholen!")
	MakeMove(player, board)
}

// Erwartet das Spielfeld eines beendeten Spiels und gibt das Ergebnis
// auf der Konsole aus. Diese Funktion hat nur dann ein sinnvolles Verhalten,
// wenn das Spiel wirklich beendet ist.
func PrintResult(board [][]string) {
	fmt.Println("Das Spiel ist aus!")
	if PlayerXWins(board) {
		fmt.Println("Spieler X hat gewonnen.")
	}
	if PlayerOWins(board) {
		fmt.Println("Spieler O hat gewonnen.")
	}
	if Draw(board) {
		fmt.Println("Unentschieden.")
	}
	PrintBoard(board)
}
