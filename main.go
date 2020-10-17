package main

import (
	"log"
)

//Grid :
//[
//["A","B","C","E"],
//["S","F","C","S"],
//["A","D","E","E"]
//]
//
//
//
//“ABCCEDX” : false
//“ECCEDAS” : true
//“DEES” : true
//
//Write a function which takes in a string and this grid and return true if the word can be formed from this grid. Up down left right. Only visit cell once

var grid = [][]string{
	{"A", "B", "C", "E"},
	{"S", "F", "C", "S"},
	{"A", "D", "E", "E"},
}

type position struct {
	i, j int
}

var pMap = make(map[position]string)

// getValue provide a cache layer to make a cell is visited twice
func getValue(i, j int) string {
	p := position{
		i: i,
		j: j,
	}
	v, ok := pMap[p]

	if ok {
		//log.Printf("chached %v: %v", p, v)
		return v
	}

	v = grid[i][j]

	pMap[p] = v

	return v
}
func main() {

	log.Printf("------------ “ABCCEDX” : false | %v", canForm("ABCCEDX"))
	log.Printf("------------ “ECCEDAS” : true | %v", canForm("ECCEDAS"))
	log.Printf("------------ “DEES” : true | %v", canForm("DEES"))
}

// canForm will try to frame the target by matching the non diagonal neighbour letters
func canForm(target string) bool {

	// pull all possible start positions for the first letter in target
	for _, position := range GetPositionsByChar(string(target[0])) {
		if ok := tryAPosition(position[0], position[1], target); ok {
			return true
		}
	}

	return false
}

// tryAPosition check wheather target can be framed for grid position i,j
func tryAPosition(i, j int, target string) bool {
	log.Printf("------------------------------------------checking for %v with %v's grid position: %v, %v", target, string(target[0]), i, j)
	var ok bool
	for li, l := range target {

		if li == 0 {
			continue
		}
		//log.Printf("-------------looking for ------- %v", string(l))
		in, jn, okn := isANeighbour(i, j, string(l))
		if okn {
			i = in
			j = jn
			ok = okn
			continue
		}
		return false
	}
	return ok
}

// isANeighbour determines char next can found in non diagonal positions from position i,j
func isANeighbour(i, j int, next string) (int, int, bool) {

	if i != 0 && j != 0 {
		// check left neighbour
		ni, nj := i, j-1

		if nj >= 0 {
			nv := getValue(ni, nj)

			if next == nv {
				log.Printf("left | %v | %v,%v > %v", next, ni, nj, nv)
				return ni, nj, true
			}
		}
	}

	// check up
	ni, nj := i-1, j
	if i == 0 {
		ni = len(grid) - 1
	}
	nv := getValue(ni, nj)

	if next == nv {
		log.Printf("  up | %v | %v,%v > %v", next, ni, nj, nv)
		return ni, nj, true
	}

	// check down
	ni, nj = i+1, j
	if i == len(grid)-1 {
		ni = 0
	}
	nv = getValue(ni, nj)

	if next == nv {
		log.Printf("down | %v | %v,%v > %v", next, ni, nj, nv)
		return ni, nj, true
	}

	if j != len(grid[len(grid)-1])-1 {
		// check right
		if len(grid[i]) > j+1 {
			ni, nj := i, j+1
			nv := getValue(ni, nj)

			if next == nv {
				log.Printf("right | %v | %v,%v > %v", next, ni, nj, nv)
				return ni, nj, true
			}
		}
	}

	return 0, 0, false
}

// GetPositionsByChar return all possible positions of a char s the grid
func GetPositionsByChar(s string) [][]int {

	var positions [][]int

	for i, set := range grid {
		for j, l := range set {
			pMap[position{
				i: i,
				j: j,
			}] = l
			if l == s {
				positions = append(positions, []int{i, j})

			}
		}
	}

	return positions
}
