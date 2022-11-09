package main

// Erwartet ein Spielfeld. Liefert true, falls das Spiel beendet ist.
func GameOver(board [][]string) bool {
	// TODO
	return true
}

// Erwartet ein Spielfeld. Liefert true, falls Spieler X gewonnen hat.
func PlayerXWins(board [][]string) bool {
	return RowContainsOnly(board, 0, "X") || RowContainsOnly(board, 1, "X") || RowContainsOnly(board, 2, "X") ||
		ColumnContainsOnly(board, 0, "X") || ColumnContainsOnly(board, 1, "X") || ColumnContainsOnly(board, 2, "X")
}

// Erwartet ein Spielfeld. Liefert true, falls Spieler O gewonnen hat.
func PlayerOWins(board [][]string) bool {
	return RowContainsOnly(board, 0, "O") || RowContainsOnly(board, 1, "O") || RowContainsOnly(board, 2, "O") ||
		ColumnContainsOnly(board, 0, "O") || ColumnContainsOnly(board, 1, "O") || ColumnContainsOnly(board, 2, "O")
}

// Erwartet ein Spielfeld. Liefert true, falls das Spiel unentschieden ist.
func Draw(board [][]string) bool {
	// TODO
	return true
}

// Hilfsfunktion: Erwartet ein Spielfeld, eine Zeilennummer und ein Symbol.
// Liefert true, falls die entsprechende Zeile nur das Symbol enthält.
func RowContainsOnly(board [][]string, row int, symbol string) bool {
	for _, sym := range board[row] {
		if sym != symbol {
			return false
		}
	}
	return true
}

// Hilfsfunktion: Erwartet ein Spielfeld, eine Spaltennummer und ein Symbol.
// Liefert true, falls die entsprechende Zeile nur das Symbol enthält.
func ColumnContainsOnly(board [][]string, col int, symbol string) bool {
	for _, row := range board {
		if row[col] != symbol {
			return false
		}
	}
	return true
}

// Hilfsfunktion: Erwartet ein Spielfeld und ein Symbol.
// Liefert die Anzahl der Vorkommen des Symbols auf dem Spielfeld.
func NumberOfOccurrences(board [][]string, symbol string) int {
	counter := 0
	for _, row := range board {
		for _, sym := range row {
			if sym == symbol {
				counter++
			}
		}
	}
	return counter
}
