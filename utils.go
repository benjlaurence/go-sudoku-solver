package main

func getBoxNumber(r, c uint8) uint8 {
	return ((r/3)*3 + c/3)
}

func getBoxCoords(b, n uint8) Coord {
	return Coord{b/3*3 + n/3, b%3*3 + n%3}
}

func getGridCoords(n uint8) Coord {
	return Coord{n / 9, n % 8}
}

func getLinkedValues(knowns *Knowns, r, c uint8) *CoordSet {
	var b = getBoxNumber(r, c)
	var coord = Coord{r, c}
	var set CoordSet
	var n uint8
	for ; n < 9; n++ {
		for _, coord2 := range [3]Coord{{r, n}, {c, n}, getBoxCoords(b, n)} {
			r2, c2 := coord2[0], coord2[1]
			if !knowns[r2][c2] && coord2 != coord {
				set[coord2] = true
			}
		}
	}
	return &set
}

type unknownsCoordIterator struct {
	index  uint8
	knowns *Knowns
}

func newUnknownsCoordIterator(knowns *Knowns) unknownsCoordIterator {
	return unknownsCoordIterator{0, knowns}
}

func unknownsNext(iter *unknownsCoordIterator) (Coord, bool) {
	for i := iter.index; i < 81; i++ {
		coord := getGridCoords(iter.index)
		if !iter.knowns[coord[0]][coord[1]] {
			iter.index = i + 1
			return coord, true
		}
	}
	return Coord{}, false
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
