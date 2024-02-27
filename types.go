package main

type Coord [2]uint8

type Values [9][9]uint8
type Knowns [9][9]bool
type NPossibles [9][9]uint8
type Possibles [9][9][9]bool
type CoordSet map[Coord]bool

type Board struct {
	values                                             Values
	knowns                                             Knowns
	possibles                                          Possibles
	rowsNPossibles, columnsNPossibles, boxesNPossibles NPossibles
}
