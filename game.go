package main

// Haupt-Funktion des Spiels.
// Initialisiert ein Spielfeld und eine Variable für den aktuellen Spieler.
// Startet dann die Hauptschleife, in der so lange Züge gemacht werden, bis
// das Spiel vorbei ist.
func main() {
	// Initialisierung des Spiels
	board := [][]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
	currentPlayer := "X"

	// Hauptschleife des Spiels ("Game Loop")
	for !GameOver(board) {
		// Einmal das Spielfeld ausgeben.
		PrintBoard(board)

		// Fragt den aktuellen Spieler nach seinem Zug und führt ihn durch.
		MakeMove(currentPlayer, board)

		// Spieler wechseln.
		currentPlayer = nextPlayer(currentPlayer)
	}

	// Nach dem Ende des Spiels das Ergebnis ausgeben.
	PrintResult(board)
}

// Erwartet den String des aktuellen Spielers ("X" oder "O").
// Liefert den String des nächsten Spielers (also den jeweils anderen).
func nextPlayer(currentPlayer string) string {
	if currentPlayer == "X" {
		return "O"
	}
	return "X"
}
