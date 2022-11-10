package main

import "fmt"

func ExamplePlayerXWins() {
	fmt.Println(PlayerXWins(testboardX1()))
	fmt.Println(PlayerXWins(testboardX2()))
	fmt.Println(PlayerXWins(testboardX3()))
	fmt.Println(PlayerXWins(testboardO1()))
	fmt.Println(PlayerXWins(testboardO2()))
	fmt.Println(PlayerXWins(testboardO3()))
	fmt.Println(PlayerXWins(testboardDraw()))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
	// false
}

func ExamplePlayerOWins() {
	fmt.Println(PlayerOWins(testboardX1()))
	fmt.Println(PlayerOWins(testboardX2()))
	fmt.Println(PlayerOWins(testboardX3()))
	fmt.Println(PlayerOWins(testboardO1()))
	fmt.Println(PlayerOWins(testboardO2()))
	fmt.Println(PlayerOWins(testboardO3()))
	fmt.Println(PlayerOWins(testboardDraw()))

	// Output:
	// false
	// false
	// false
	// true
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

func ExampleMoveAllowed() {
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

func ExampleNextPlayer() {
	fmt.Println(NextPlayer("X"))
	fmt.Println(NextPlayer("O"))

	// Output:
	// O
	// X
}
