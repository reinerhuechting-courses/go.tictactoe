package main

import "fmt"

// Erwartet ein Spielfeld und gibt es auf der Konsole aus.
// Dabei sollen die Zeilen des Spielfelds einzeln ausgegeben und
// geeignete Trennzeichen verwendet werden (siehe Tests).
func PrintBoard(board [][]string) {
	fmt.Println(board)
}

// Erwartet einen Spieler-String und ein Spielfeld.
// Fragt den Spieler nach seinem Zug und führt diesen aus,
// sofern der Zug erlaubt und möglich ist.
// Ist der Zug ungültig, muss der Spieler erneut gefragt werden.
func MakeMove(player string, board [][]string) {
	// TODO
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
