package main

type Coord [2]uint8

type Values [9][9]uint8
type NPossibles [9][9]uint8
type Possibles [9][9][9]bool
type CoordSet map[Coord]bool
type ValueSet []uint8

type Board struct {
	values                                             Values
	possibles                                          Possibles
	rowsNPossibles, columnsNPossibles, boxesNPossibles NPossibles
	unknowns                                           CoordSet
}
