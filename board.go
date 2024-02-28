package main

import "fmt"

func newBoard(values *Values) *Board {
	board := Board{}
	var r, c uint8
	for r = 0; r < 9; r++ {
		for c = 0; c < 9; c++ {
			board.rowsNPossibles[r][c] = 9
			board.columnsNPossibles[r][c] = 9
			board.boxesNPossibles[r][c] = 9
			for v := 0; v < 9; v++ {
				board.possibles[r][c][v] = true
			}
		}
	}

	for r = 0; r < 9; r++ {
		for c = 0; c < 9; c++ {
			if values[r][c] != 0 {
				if setValue(&board, r, c, values[r][c]) == nil {
					return nil
				}
			}
		}
	}
	loadUnknowns(&board)

	return &board
}

func CopyBoard(board *Board) *Board {
	var newBoard = *board
	loadUnknowns(&newBoard)
	return &newBoard
}

func loadUnknowns(board *Board) {
	board.unknowns = make(CoordSet)
	var r, c uint8
	for r = 0; r < 9; r++ {
		for c = 0; c < 9; c++ {
			if board.values[r][c] == 0 {
				board.unknowns[Coord{r, c}] = true
			} else {
				delete(board.unknowns, Coord{r, c})
			}
		}
	}
}

func setNotPossible(board *Board, r, c, v uint8) bool {
	if board.possibles[r][c][v] {
		var b = getBoxNumber(r, c)
		board.possibles[r][c][v] = false
		board.rowsNPossibles[r][v] -= 1
		board.columnsNPossibles[c][v] -= 1
		board.boxesNPossibles[b][v] -= 1
		if ifRCBNContains(board, r, c, b, v, 0) {
			return false
		}
	}
	return true
}
func setValue(board *Board, r, c, v uint8) CoordSet {
	v_p := v - 1
	if (!board.possibles[r][c][v_p]) || board.values[r][c] != 0 {

		return nil
	}
	delete(board.unknowns, Coord{r, c})
	var n uint8
	for ; n < 9; n++ {
		if n != v_p {
			if !setNotPossible(board, r, c, n) {
				return nil
			}
		}
	}

	board.values[r][c] = v
	var linked_values = getLinkedValues(board, r, c)
	for coord2 := range linked_values {
		r2, c2 := coord2[0], coord2[1]
		if !setNotPossible(board, r2, c2, v_p) {
			return nil
		}
	}
	return linked_values
}

func getNPossibles(board *Board, r, c uint8) uint8 {
	return countTrue(&board.possibles[r][c])
}

func checkSolved(board *Board) bool {
	var r, c uint8
	for r = 0; r < 9; r++ {
		for c = 0; c < 9; c++ {
			if board.values[r][c] == 0 {
				return false
			}
		}
	}
	return true
}
func getNextPossible(board *Board, r, c uint8) uint8 {
	for i, v := range board.possibles[r][c] {
		if v {
			return uint8(i) + 1
		}
	}
	return 0
}

func getPossiblesAsIndexes(board *Board, r, c uint8) ValueSet {
	var set ValueSet
	for i, v := range board.possibles[r][c] {
		if v {
			set = append(set, uint8(i))
		}
	}
	return set
}

func BoardToString(board *Board) string {
	return ValuesToString(&board.values)
}
func PrintBoard(board *Board) {
	fmt.Println(BoardToString(board))
}
