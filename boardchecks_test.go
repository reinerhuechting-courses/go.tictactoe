package main

import "fmt"

func testboard1() [][]string {
	return [][]string{
		{"*", " ", " "},
		{"*", "*", " "},
		{"*", "*", "*"}}
}

func ExampleRowContainsOnly() {
	fmt.Println(RowContainsOnly(testboard1(), 0, "*"))
	fmt.Println(RowContainsOnly(testboard1(), 1, "*"))
	fmt.Println(RowContainsOnly(testboard1(), 2, "*"))

	// Output:
	// false
	// false
	// true
}

func ExampleColumnContainsOnly() {
	fmt.Println(ColumnContainsOnly(testboard1(), 0, "*"))
	fmt.Println(ColumnContainsOnly(testboard1(), 1, "*"))
	fmt.Println(ColumnContainsOnly(testboard1(), 2, "*"))

	// Output:
	// true
	// false
	// false
}

func ExampleNumberOfOccurrences() {
	fmt.Println(NumberOfOccurrences(testboard1(), " "))
	fmt.Println(NumberOfOccurrences(testboard1(), "*"))

	// Output:
	// 3
	// 6
}
