package main

import (
	"fmt"
	"strings"
)

// Erwartet ein Spielfeld und gibt es auf der Konsole aus.
// Dabei sollen die Zeilen des Spielfelds einzeln ausgegeben und
// geeignete Trennzeichen verwendet werden (siehe Tests).
func PrintBoard(board [][]string) {
	fmt.Println("+---+---+---+")
	for _, row := range board {
		fmt.Println("|", strings.Join(row, " | "), "|")
		fmt.Println("+---+---+---+")
	}
}
