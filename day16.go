package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
)

type ReindeerPos struct {
	X int
	Y int
}

func GetReindeerPos(m [][]rune) ReindeerPos {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 'S' {
				return ReindeerPos{X: i, Y: j}
			}
		}
	}
	fmt.Printf("Error GetRPos could not find @\n")
	return ReindeerPos{X: -1, Y: -1}
}

func SolveMaze(m [][]rune, rp ReindeerPos, dx int, dy int) (int, bool) {
	minScore := 99999999
	hasWayOut := false
	prevChar := m[rp.X][rp.Y]
	bestDx := 0
	bestDy := 0
	if m[rp.X][rp.Y] == 'E' {
		return 0, true
	} else if m[rp.X][rp.Y] == '#' || m[rp.X][rp.Y] == '>' || m[rp.X][rp.Y] == '<' || m[rp.X][rp.Y] == '^' || m[rp.X][rp.Y] == 'V' || m[rp.X][rp.Y] == 'X' {
		return 0, false
	} else {
		m[rp.X][rp.Y] = 'X'
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if (i == 0 || j == 0) && !(i == 0 && j == 0) {
					supplementScore := 1000
					if (i == dx && j == dy) || (0 == dx && 0 == dy) {
						supplementScore = 1
					}
					score, stepHasSucceeded := SolveMaze(m, ReindeerPos{X: rp.X + i, Y: rp.Y + j}, i, j)
					score += supplementScore
					if stepHasSucceeded {
						hasWayOut = true
						if score < minScore {
							minScore = score
							bestDx = i
							bestDy = j
						}
					}
				}
			}
		}

	}
	if hasWayOut {
		if prevChar != 'S' {
			if bestDx == 1 {
				m[rp.X][rp.Y] = 'V'
			} else if bestDx == -1 {
				m[rp.X][rp.Y] = '^'
			} else if bestDy == 1 {
				m[rp.X][rp.Y] = '>'
			} else if bestDy == -1 {
				m[rp.X][rp.Y] = '<'
			}
		} else {
			m[rp.X][rp.Y] = 'S'
		}

	} else if prevChar != 'S' {
		m[rp.X][rp.Y] = '.'
	} else {
		m[rp.X][rp.Y] = 'S'
	}
	return minScore, hasWayOut
}

func main() {
	filePath := "inputs/day16-sample.txt"
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	m, _, _ := utils.LinesTo2DArray(lines)
	rp := GetReindeerPos(m)
	utils.PrintRune2DArray(m)
	fmt.Printf("Reindeer at %d, %d\n", rp.X, rp.Y)
	score, _ := SolveMaze(m, rp, 0, 0)
	utils.PrintRune2DArray(m)
	fmt.Printf("%d\n", score)

}
