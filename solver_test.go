package main

import (
	"fmt"
	"testing"
)

var EASY Values = Values{
	{0, 3, 0, 0, 0, 0, 1, 0, 0},
	{5, 0, 0, 0, 9, 6, 0, 2, 0},
	{2, 8, 0, 3, 4, 0, 0, 0, 0},
	{0, 9, 4, 0, 0, 0, 8, 0, 0},
	{0, 2, 5, 8, 7, 9, 0, 4, 3},
	{0, 7, 6, 5, 3, 0, 9, 1, 2},
	{9, 0, 0, 6, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 8, 3, 5, 0, 1},
	{0, 5, 0, 0, 0, 0, 0, 3, 6},
}

var EASY_SOLVED Values = Values{
	{6, 3, 9, 7, 2, 8, 1, 5, 4},
	{5, 4, 7, 1, 9, 6, 3, 2, 8},
	{2, 8, 1, 3, 4, 5, 7, 6, 9},
	{3, 9, 4, 2, 6, 1, 8, 7, 5},
	{1, 2, 5, 8, 7, 9, 6, 4, 3},
	{8, 7, 6, 5, 3, 4, 9, 1, 2},
	{9, 1, 3, 6, 5, 2, 4, 8, 7},
	{7, 6, 2, 4, 8, 3, 5, 9, 1},
	{4, 5, 8, 9, 1, 7, 2, 3, 6},
}

var MEDIUM Values = Values{
	{0, 2, 0, 6, 0, 8, 0, 0, 0},
	{5, 8, 0, 0, 0, 9, 7, 0, 0},
	{0, 0, 0, 0, 4, 0, 0, 0, 0},
	{3, 7, 0, 0, 0, 0, 5, 0, 0},
	{6, 0, 0, 0, 0, 0, 0, 0, 4},
	{0, 0, 8, 0, 0, 0, 0, 1, 3},
	{0, 0, 0, 0, 2, 0, 0, 0, 0},
	{0, 0, 9, 8, 0, 0, 0, 3, 6},
	{0, 0, 0, 3, 0, 6, 0, 9, 0},
}

var MEDIUM_SOLVED Values = Values{
	{1, 2, 3, 6, 7, 8, 9, 4, 5},
	{5, 8, 4, 2, 3, 9, 7, 6, 1},
	{9, 6, 7, 1, 4, 5, 3, 2, 8},
	{3, 7, 2, 4, 6, 1, 5, 8, 9},
	{6, 9, 1, 5, 8, 3, 2, 7, 4},
	{4, 5, 8, 7, 9, 2, 6, 1, 3},
	{8, 3, 6, 9, 2, 4, 1, 5, 7},
	{2, 1, 9, 8, 5, 7, 4, 3, 6},
	{7, 4, 5, 3, 1, 6, 8, 9, 2},
}

var HARD Values = Values{
	{0, 2, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 6, 0, 0, 0, 0, 3},
	{0, 7, 4, 0, 8, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 3, 0, 0, 2},
	{0, 8, 0, 0, 4, 0, 0, 1, 0},
	{6, 0, 0, 5, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 7, 8, 0},
	{5, 0, 0, 0, 0, 9, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 4, 0},
}

var HARD_SOLVED Values = Values{
	{1, 2, 6, 4, 3, 7, 9, 5, 8},
	{8, 9, 5, 6, 2, 1, 4, 7, 3},
	{3, 7, 4, 9, 8, 5, 1, 2, 6},
	{4, 5, 7, 1, 9, 3, 8, 6, 2},
	{9, 8, 3, 2, 4, 6, 5, 1, 7},
	{6, 1, 2, 5, 7, 8, 3, 9, 4},
	{2, 6, 9, 3, 1, 4, 7, 8, 5},
	{5, 4, 8, 7, 6, 9, 2, 3, 1},
	{7, 3, 1, 8, 5, 2, 6, 4, 9},
}

var UNSOLVABLE Values = Values{
	{8, 2, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 6, 0, 0, 0, 0, 3},
	{0, 7, 4, 0, 8, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 3, 0, 0, 2},
	{0, 8, 0, 0, 4, 0, 0, 1, 0},
	{6, 0, 0, 5, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 7, 8, 0},
	{5, 0, 0, 0, 0, 9, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 4, 0},
}

var BROKEN Values = Values{
	{1, 3, 0, 0, 0, 0, 1, 0, 0},
	{5, 0, 0, 0, 9, 6, 0, 2, 0},
	{2, 8, 0, 3, 4, 0, 0, 0, 0},
	{0, 9, 4, 0, 0, 0, 8, 0, 0},
	{0, 2, 5, 8, 7, 9, 0, 4, 3},
	{0, 7, 6, 5, 3, 0, 9, 1, 2},
	{9, 0, 0, 6, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 8, 3, 5, 0, 1},
	{0, 5, 0, 0, 0, 0, 0, 3, 6},
}

func TestEasy(t *testing.T) {
	board := newBoard(&EASY)
	fmt.Println("Easy:\n" + ValuesToString(&EASY))

	solved_board := Solve(board)

	if solved_board == nil {
		t.Error("Can't solve sudoku")
	}
	if solved_board.values != EASY_SOLVED {
		t.Error("Wrong answer, expected:\n" + ValuesToString(&EASY_SOLVED) + "Result:\n" + BoardToString(solved_board))
	} else {
		fmt.Println("Solution:")
		PrintBoard(solved_board)
	}
}

func TestMedium(t *testing.T) {
	board := newBoard(&MEDIUM)
	fmt.Println("Medium:\n" + ValuesToString(&MEDIUM))
	solved_board := Solve(board)
	if solved_board == nil {
		t.Error("Can't solve sudoku")
	}
	if solved_board.values != MEDIUM_SOLVED {
		t.Error("Wrong answer, expected:\n" + ValuesToString(&MEDIUM_SOLVED) + "Result:\n" + BoardToString(solved_board))
	} else {
		fmt.Println("Solution:")
		PrintBoard(solved_board)
	}
}

func TestHARD(t *testing.T) {
	board := newBoard(&HARD)
	fmt.Println("Hard:\n" + ValuesToString(&HARD))
	solved_board := Solve(board)
	if solved_board == nil {
		t.Error("Can't solve sudoku")
	} else if solved_board.values != HARD_SOLVED {
		t.Error("Wrong answer, expected:\n" + ValuesToString(&HARD_SOLVED) + "Result:\n" + BoardToString(solved_board))
	} else {
		fmt.Println("Solution:")
		PrintBoard(solved_board)
	}

}
