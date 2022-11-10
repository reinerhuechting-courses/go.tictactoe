package main

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

func testboardX3() [][]string {
	return [][]string{
		{"X", "O", "O"},
		{"X", "X", "O"},
		{"O", "X", "X"}}
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

func testboardO3() [][]string {
	return [][]string{
		{"O", "X", "X"},
		{"O", "O", "X"},
		{"X", "O", "O"}}
}

func testboardDraw() [][]string {
	return [][]string{
		{"O", "X", "X"},
		{"X", "X", "O"},
		{"O", "O", "X"}}
}
