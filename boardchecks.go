package main

// Erwartet ein Spielfeld. Liefert true, falls das Spiel beendet ist.
func GameOver(board [][]string) bool {
	return PlayerXWins(board) || PlayerOWins(board) || Draw(board)
}

// Erwartet ein Spielfeld. Liefert true, falls Spieler X gewonnen hat.
func PlayerXWins(board [][]string) bool {
	return RowContainsOnly(board, 0, "X") || RowContainsOnly(board, 1, "X") || RowContainsOnly(board, 2, "X") ||
		ColumnContainsOnly(board, 0, "X") || ColumnContainsOnly(board, 1, "X") || ColumnContainsOnly(board, 2, "X") ||
		Diag1ContainsOnly(board, "X") || Diag2ContainsOnly(board, "X")
}

// Erwartet ein Spielfeld. Liefert true, falls Spieler O gewonnen hat.
func PlayerOWins(board [][]string) bool {
	return RowContainsOnly(board, 0, "O") || RowContainsOnly(board, 1, "O") || RowContainsOnly(board, 2, "O") ||
		ColumnContainsOnly(board, 0, "O") || ColumnContainsOnly(board, 1, "O") || ColumnContainsOnly(board, 2, "O") ||
		Diag1ContainsOnly(board, "O") || Diag2ContainsOnly(board, "O")
}

// Erwartet ein Spielfeld. Liefert true, falls das Spiel unentschieden ist.
func Draw(board [][]string) bool {
	return !PlayerXWins(board) && !PlayerOWins(board) && NumberOfOccurrences(board, " ") == 0
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

// Erwartet ein Spielfeld und ein Symbol.
// Liefert true, falls die Diagonale von links oben nach rechts unten
// nur dieses Symbol enthält.
func Diag1ContainsOnly(board [][]string, symbol string) bool {
	for i, row := range board {
		if row[i] != symbol {
			return false
		}
	}
	return true
}

// Erwartet ein Spielfeld und ein Symbol.
// Liefert true, falls die Diagonale von links unten nach rechts oben
// nur dieses Symbol enthält.
func Diag2ContainsOnly(board [][]string, symbol string) bool {
	for i, row := range board {
		if row[2-i] != symbol {
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

// Erwartet ein Spielfeld, eine Zeilen- und eine Spaltennummer.
// Liefert true, falls ein Zug an der angegebenen Stelle erlaubt ist.
func MoveAllowed(board [][]string, row, col int) bool {
	return row >= 0 && row <= 2 && col >= 0 && col <= 2 && board[row][col] == " "
}
