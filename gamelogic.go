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

// Erwartet ein Spielfeld, eine Zeilen- und eine Spaltennummer.
// Liefert true, falls ein Zug an der angegebenen Stelle erlaubt ist.
func MoveAllowed(board [][]string, row, col int) bool {
	return row >= 0 && row <= 2 && col >= 0 && col <= 2 && board[row][col] == " "
}

// Erwartet den String des aktuellen Spielers ("X" oder "O").
// Liefert den String des nächsten Spielers (also den jeweils anderen).
func NextPlayer(currentPlayer string) string {
	if currentPlayer == "X" {
		return "O"
	}
	return "X"
}
