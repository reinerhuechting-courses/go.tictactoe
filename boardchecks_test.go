package main

import "fmt"

func testboardGeneric1() [][]string {
	return [][]string{
		{"*", " ", " "},
		{"*", "*", " "},
		{"*", "*", "*"}}
}

func testboardGeneric2() [][]string {
	return [][]string{
		{" ", " ", "*"},
		{" ", "*", "*"},
		{"*", "*", "*"}}
}

func testboardX1() [][]string {
	return [][]string{
		{" ", "X", " "},
		{"O", "X", "X"},
		{"O", "X", "O"}}
}

func testboardX2() [][]string {
	return [][]string{
		{" ", "O", " "},
		{"X", "X", "X"},
		{"O", "X", "O"}}
}

func testboardO1() [][]string {
	return [][]string{
		{" ", "O", " "},
		{"X", "O", "O"},
		{"X", "O", "X"}}
}

func testboardO2() [][]string {
	return [][]string{
		{" ", "X", " "},
		{"O", "O", "O"},
		{"X", "O", "X"}}
}

func testboardDraw() [][]string {
	return [][]string{
		{"O", "X", "X"},
		{"X", "X", "O"},
		{"O", "O", "X"}}
}

func ExampleRowContainsOnly() {
	fmt.Println(RowContainsOnly(testboardGeneric1(), 0, "*"))
	fmt.Println(RowContainsOnly(testboardGeneric1(), 1, "*"))
	fmt.Println(RowContainsOnly(testboardGeneric1(), 2, "*"))

	// Output:
	// false
	// false
	// true
}

func ExampleColumnContainsOnly() {
	fmt.Println(ColumnContainsOnly(testboardGeneric1(), 0, "*"))
	fmt.Println(ColumnContainsOnly(testboardGeneric1(), 1, "*"))
	fmt.Println(ColumnContainsOnly(testboardGeneric1(), 2, "*"))

	// Output:
	// true
	// false
	// false
}

func ExampleDiag1ContainsOnly() {
	fmt.Println(Diag1ContainsOnly(testboardGeneric1(), "*"))
	fmt.Println(Diag1ContainsOnly(testboardGeneric1(), " "))

	// Output:
	// true
	// false
}

func ExampleDiag2ContainsOnly() {
	fmt.Println(Diag2ContainsOnly(testboardGeneric2(), "*"))
	fmt.Println(Diag2ContainsOnly(testboardGeneric2(), " "))

	// Output:
	// true
	// false
}

func ExampleNumberOfOccurrences() {
	fmt.Println(NumberOfOccurrences(testboardGeneric1(), " "))
	fmt.Println(NumberOfOccurrences(testboardGeneric1(), "*"))

	// Output:
	// 3
	// 6
}

func ExamplePlayerXWins() {
	fmt.Println(PlayerXWins(testboardX1()))
	fmt.Println(PlayerXWins(testboardX2()))
	fmt.Println(PlayerXWins(testboardO1()))
	fmt.Println(PlayerXWins(testboardO2()))
	fmt.Println(PlayerXWins(testboardDraw()))

	// Output:
	// true
	// true
	// false
	// false
	// false
}

func ExamplePlayerOWins() {
	fmt.Println(PlayerOWins(testboardX1()))
	fmt.Println(PlayerOWins(testboardX2()))
	fmt.Println(PlayerOWins(testboardO1()))
	fmt.Println(PlayerOWins(testboardO2()))
	fmt.Println(PlayerOWins(testboardDraw()))

	// Output:
	// false
	// false
	// true
	// true
	// false
}

func ExampleDraw() {
	fmt.Println(Draw(testboardX1()))
	fmt.Println(Draw(testboardX2()))
	fmt.Println(Draw(testboardO1()))
	fmt.Println(Draw(testboardO2()))
	fmt.Println(Draw(testboardDraw()))

	// Output:
	// false
	// false
	// false
	// false
	// true
}

func ExampleGameOver() {
	fmt.Println(GameOver(testboardX1()))
	fmt.Println(GameOver(testboardX2()))
	fmt.Println(GameOver(testboardO1()))
	fmt.Println(GameOver(testboardO2()))
	fmt.Println(GameOver(testboardDraw()))
	fmt.Println(GameOver(testboardGeneric1()))

	// Output:
	// true
	// true
	// true
	// true
	// true
	// false
}

func ExampleMoveAllowd() {
	fmt.Println(MoveAllowed(testboardGeneric1(), 0, 0))
	fmt.Println(MoveAllowed(testboardGeneric1(), 0, 1))
	fmt.Println(MoveAllowed(testboardGeneric1(), 1, 2))
	fmt.Println(MoveAllowed(testboardGeneric1(), 0, 3))
	fmt.Println(MoveAllowed(testboardGeneric1(), -1, 2))

	// Output:
	// false
	// true
	// true
	// false
	// false
}
