package main

func newBoard(values Values) *Board {
	var board Board
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
				setValue(&board, r, c, values[r][c])
			}
		}
	}

	return &board
}

func setNotPossible(board *Board, r, c, v uint8) bool {
	if board.possibles[r][c][v] {
		var b = getBoxNumber(r, c)
		board.possibles[r][c][v] = false
		board.rowsNPossibles[r][v] -= 1
		board.columnsNPossibles[c][v] -= 1
		board.boxesNPossibles[b][v] -= 1

		if board.rowsNPossibles[r][v] == 0 || board.columnsNPossibles[c][v] == 0 || board.boxesNPossibles[b][v] == 0 {
			return false
		}
	}
	return true
}
func setValue(board *Board, r, c, v uint8) *CoordSet {
	v_p := v - 1

	if !board.possibles[r][c][v_p] || board.knowns[r][c] {
		return nil
	}
	var n uint8
	for ; n < 9; n++ {
		if n != v_p {
			setNotPossible(board, r, c, n)
		}
	}

	board.values[r][c] = v
	board.knowns[r][c] = true
	var linked_values = getLinkedValues(&board.knowns, r, c)

	for coord2 := range *linked_values {
		r2, c2 := coord2[0], coord2[1]
		setNotPossible(board, r2, c2, v_p)
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
			if !board.knowns[r][c] {
				return false
			}
		}
	}
	return true
}
