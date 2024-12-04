package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
)

type Direction int

const (
	U Direction = iota
	D
	L
	R
	UL
	UR
	DL
	DR
)

// 106784935 too high
func recursiveXmas(m [][]rune, x int, y int, sizeX int, sizeY int, word []rune, d Direction) int {
	currentRune := word[0]
	nextWord := word[1:]
	count := 0
	if x < 0 || x >= sizeX || y < 0 || y >= sizeY {
		return 0
	}
	if m[x][y] == currentRune {
		if len(nextWord) == 0 {
			return 1
		} else {
			if d == R {
				count += recursiveXmas(m, x+1, y, sizeX, sizeY, nextWord, d)
			} else if d == L {
				count += recursiveXmas(m, x-1, y, sizeX, sizeY, nextWord, d)
			} else if d == D {
				count += recursiveXmas(m, x, y+1, sizeX, sizeY, nextWord, d)
			} else if d == U {
				count += recursiveXmas(m, x, y-1, sizeX, sizeY, nextWord, d)
			} else if d == DR {
				count += recursiveXmas(m, x+1, y+1, sizeX, sizeY, nextWord, d)
			} else if d == DL {
				count += recursiveXmas(m, x-1, y-1, sizeX, sizeY, nextWord, d)
			} else if d == UR {
				count += recursiveXmas(m, x+1, y-1, sizeX, sizeY, nextWord, d)
			} else if d == UL {
				count += recursiveXmas(m, x-1, y+1, sizeX, sizeY, nextWord, d)
			}
		}
		return count
	}
	return 0
}

func main() {
	filePath := "inputs/day4.txt"

	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	m, sizeX, sizeY := utils.LinesTo2DArray(lines)
	count := 0
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			//print(string(m[i][j]))
			if m[i][j] == 'X' {
				count += recursiveXmas(m, i, j, sizeX, sizeY, []rune("XMAS"), U)
				count += recursiveXmas(m, i, j, sizeX, sizeY, []rune("XMAS"), D)
				count += recursiveXmas(m, i, j, sizeX, sizeY, []rune("XMAS"), L)
				count += recursiveXmas(m, i, j, sizeX, sizeY, []rune("XMAS"), R)
				count += recursiveXmas(m, i, j, sizeX, sizeY, []rune("XMAS"), UR)
				count += recursiveXmas(m, i, j, sizeX, sizeY, []rune("XMAS"), UL)
				count += recursiveXmas(m, i, j, sizeX, sizeY, []rune("XMAS"), DR)
				count += recursiveXmas(m, i, j, sizeX, sizeY, []rune("XMAS"), DL)
			}
		}
	}

	fmt.Println(count)

}
