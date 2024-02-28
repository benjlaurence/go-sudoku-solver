package main

import (
	"fmt"
	"strconv"
)

func getBoxNumber(r, c uint8) uint8 {
	return ((r/3)*3 + c/3)
}

func getBoxCoords(b, n uint8) Coord {
	return Coord{b/3*3 + n/3, b%3*3 + n%3}
}

func getLinkedValues(board *Board, r, c uint8) CoordSet {
	var b = getBoxNumber(r, c)
	var coord = Coord{r, c}
	set := make(CoordSet)
	var n uint8
	for ; n < 9; n++ {
		for _, coord2 := range [3]Coord{{r, n}, {n, c}, getBoxCoords(b, n)} {
			r2, c2 := coord2[0], coord2[1]
			if board.values[r2][c2] == 0 && coord2 != coord {
				set[coord2] = true
			}
		}
	}
	return set
}

func countTrue(bools *[9]bool) uint8 {
	var sum uint8
	for _, v := range *bools {
		if v {
			sum++
		}
	}
	return sum
}

func getOnlyValue(bools *[9]bool) uint8 {
	for n, v := range *bools {
		if v {
			return uint8(n) + 1
		}
	}
	return 0
}

func ValuesToString(values *Values) string {
	var str string
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			str += strconv.Itoa(int(values[r][c])) + " "
			if c == 2 || c == 5 {
				str += "│ "
			}
		}
		str += "\n"
		if r == 2 || r == 5 {
			str += "──────┼───────┼──────\n"
		}
	}
	return str
}

func PrintValues(values *Values) {
	fmt.Println("\n" + ValuesToString(values))
}

func ifRCBNContains(board *Board, r, c, b, n, v uint8) bool {
	return (board.rowsNPossibles[r][n] == v || board.columnsNPossibles[c][n] == v || board.boxesNPossibles[b][n] == v)
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	newMap := make(map[K]V)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}
