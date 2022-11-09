package main

func ExamplePrintBoard() {
	board := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"6", "8", "9"},
	}
	PrintBoard(board)

	// Output:
	// +---+---+---+
	// | 1 | 2 | 3 |
	// +---+---+---+
	// | 4 | 5 | 6 |
	// +---+---+---+
	// | 6 | 8 | 9 |
	// +---+---+---+
}
