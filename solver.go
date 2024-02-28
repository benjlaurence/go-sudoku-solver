package main

func inferValues(board *Board, coords CoordSet) (bool, bool) {
	var found bool
	for coord := range coords {
		r, c := coord[0], coord[1]

		if board.values[r][c] == 0 {
			trues := countTrue(&board.possibles[r][c])
			if trues == 0 {
				return false, false
			}
			if trues == 1 {
				val := getOnlyValue(&board.possibles[r][c])
				linked_coords := setValue(board, r, c, val)
				if linked_coords == nil {
					return false, false
				}
				okay, _ := inferValues(board, linked_coords)
				if !okay {
					return false, false
				}
				found = true
			} else {
				b := getBoxNumber(r, c)
				for _, ival := range getPossiblesAsIndexes(board, r, c) {
					if ifRCBNContains(board, r, c, b, ival, 1) {
						linked_coords := setValue(board, r, c, ival+1)
						if linked_coords == nil {
							return false, false
						}
						found = true
						break
					}
				}
			}
		}
	}
	return true, found
}

func getNextValueToTry(board *Board) (bool, uint8, uint8, uint8) {
	var minC, minR uint8
	var minUnknownN uint8 = 9
	var r, c uint8
	for coord := range board.unknowns {
		r, c = coord[0], coord[1]
		if board.values[r][c] == 0 {
			nPos := getNPossibles(board, r, c)
			if nPos == 0 {
				return false, 0, 0, 0
			}
			if nPos == 2 {
				return true, r, c, getNextPossible(board, r, c)
			} else {
				if nPos < minUnknownN {
					minR, minC, minUnknownN = r, c, nPos
				}
			}
		}
	}

	return true, minR, minC, getNextPossible(board, minR, minC)

}

func inferToCompletion(board *Board) bool {
	for solving := true; solving; {
		var solvable bool

		new_unknowns := CopyMap(board.unknowns)
		solvable, solving = inferValues(board, new_unknowns)
		if !solvable {
			return false
		}
	}
	return true
}

func Solve(board *Board) *Board {
	if !inferToCompletion(board) {
		return nil
	}
	for !checkSolved(board) {
		solvable, r, c, v := getNextValueToTry(board)
		if !solvable {
			return nil
		}
		var test_board = CopyBoard(board)
		setValue(test_board, r, c, v)
		result := Solve(test_board)
		if result != nil {
			return result
		}
		setNotPossible(board, r, c, v-1)
		if !inferToCompletion(board) {
			return nil
		}
	}

	return board
}
