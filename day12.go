package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
)

type FieldPosition struct {
	X int
	Y int
}

type BorderSide int

const (
	U BorderSide = iota
	D
	L
	R
)

func checkExtraPerimeter(m [][]rune, f FieldPosition, expected rune) int {
	res := 0
	testPos := []FieldPosition{FieldPosition{X: -1, Y: 0}, FieldPosition{X: 1, Y: 0}, FieldPosition{X: 0, Y: -1}, FieldPosition{X: 0, Y: 1}}
	for _, p := range testPos {
		x := f.X + p.X
		y := f.Y + p.Y
		if x < 0 {
			res += 1
		}
		if x >= len(m) {
			res += 1
		}
		if y < 0 {
			res += 1
		}
		if y >= len(m[0]) {
			res += 1
		}
		if !(x < 0 || x >= len(m) || y < 0 || y >= len(m[0])) {
			if m[x][y] != expected && m[x][y] != '%' {
				res += 1
			}
		}
	}
	return res
}

func replaceCharInMap(m [][]rune, old rune, new rune) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == old {
				m[i][j] = new
			}
		}
	}
}

func hasBorder(m [][]rune, x int, y int, direction BorderSide) bool {
	res := false
	if m[x][y] == '%' {
		if direction == U {
			if x == 0 {
				return true
			} else if m[x-1][y] != '%' {
				return true
			}
		} else if direction == D {
			if x == (len(m) - 1) {
				return true
			} else if m[x+1][y] != '%' {
				return true
			}
		} else if direction == L {
			if y == 0 {
				return true
			} else if m[x][y-1] != '%' {
				return true
			}
		} else if direction == R {
			if y == (len(m[x]) - 1) {
				return true
			} else if m[x][y+1] != '%' {
				return true
			}
		}
	}
	return res
}

func checkNumberOfSides(m [][]rune, fixed int, direction BorderSide) int {
	nBorders := 0
	hasSide := false
	if direction == U || direction == D {
		max := len(m[fixed])
		for i := 0; i < max; i++ {
			hasSide = false
			for i < max && hasBorder(m, fixed, i, U) {
				hasSide = true
				i++
			}
			if hasSide {
				nBorders += 1
			}
		}
		for i := 0; i < max; i++ {
			hasSide = false
			for i < max && hasBorder(m, fixed, i, D) {
				hasSide = true
				i++
			}
			if hasSide {
				nBorders += 1
			}
		}
	}
	if direction == L || direction == R {
		max := len(m)
		for i := 0; i < max; i++ {
			hasSide = false
			for i < max && hasBorder(m, i, fixed, L) {
				hasSide = true
				i++
			}
			if hasSide {
				nBorders += 1
			}
		}
		for i := 0; i < max; i++ {
			hasSide = false
			for i < max && hasBorder(m, i, fixed, R) {
				hasSide = true
				i++
			}
			if hasSide {
				nBorders += 1
			}
		}
	}
	return nBorders
}

func countNumberOfSides(m [][]rune) int {
	nCount := 0
	//for each row, check the  number of upper and lower sides
	for i := 0; i < len(m); i++ {
		nCount += checkNumberOfSides(m, i, U)
		nCount += checkNumberOfSides(m, i, D)
	}
	//for each column, check the  number of left and right sides
	for i := 0; i < len(m[0]); i++ {
		nCount += checkNumberOfSides(m, i, L)
		nCount += checkNumberOfSides(m, i, R)
	}
	return nCount
}

func recursivePrice(m [][]rune, f FieldPosition, expected rune) (int, int) {
	extraArea := 0
	extraPerimeter := 0
	testPos := []FieldPosition{FieldPosition{X: -1, Y: 0}, FieldPosition{X: 1, Y: 0}, FieldPosition{X: 0, Y: -1}, FieldPosition{X: 0, Y: 1}}
	if f.X < 0 || f.X >= len(m) || f.Y < 0 || f.Y >= len(m[0]) {
		return 0, 0 // Discard trails that go out of the map
	}
	if m[f.X][f.Y] == expected {
		extraArea = 1
		extraPerimeter = checkExtraPerimeter(m, f, expected)
		m[f.X][f.Y] = '%'
		for _, p := range testPos {
			neighbor := FieldPosition{X: f.X + p.X, Y: f.Y + p.Y}
			areaNeighbor, perimeterNeighbor := recursivePrice(m, neighbor, expected)
			extraArea += areaNeighbor
			extraPerimeter += perimeterNeighbor
		}
		return extraArea, extraPerimeter
	}
	return 0, 0
}

func main() {
	filePath := "inputs/day12-sample.txt"
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	m, sizeX, sizeY := utils.LinesTo2DArray(lines)
	totalCount := 0
	totalCountWithSides := 0
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			f := FieldPosition{X: i, Y: j}
			if m[i][j] != '*' {
				c := m[i][j]
				area, perimeter := recursivePrice(m, f, c)
				totalCount += area * perimeter
				nSides := countNumberOfSides(m)
				totalCountWithSides += area * nSides
				fmt.Printf("[%c]@%d,%d{Area: %d, Perimeter: %d, nSides: %d}\n", c, i, j, area, perimeter, nSides)
				replaceCharInMap(m, '%', '*')
			}
		}
	}
	fmt.Println(totalCount)
	fmt.Println(totalCountWithSides)

}
