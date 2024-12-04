package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
)

// 106784935 too high
func xmas(m [][]rune, x int, y int, sizeX int, sizeY int) int {
	if x < 0 || x >= sizeX || y < 0 || y >= sizeY {
		return 0
	}
	if (m[x+1][y+1] == 'M' && m[x-1][y-1] == 'S') || (m[x+1][y+1] == 'S' && m[x-1][y-1] == 'M') {
		if (m[x-1][y+1] == 'M' && m[x+1][y-1] == 'S') || (m[x-1][y+1] == 'S' && m[x+1][y-1] == 'M') {
			return 1
		}
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
	for i := 1; i < sizeX-1; i++ {
		for j := 1; j < sizeY-1; j++ {
			if m[i][j] == 'A' {
				count += xmas(m, i, j, sizeX, sizeY)
			}
		}
	}

	fmt.Println(count)

}
