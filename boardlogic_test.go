package main

import "fmt"

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
