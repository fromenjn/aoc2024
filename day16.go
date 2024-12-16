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

func SolveMaze2(m [][]rune, rp ReindeerPos, dx int, dy int, currentScore int, memo map[ReindeerPos]int) (int, bool) {
	hasWayOut := false
	if m[rp.X][rp.Y] == '#' {
		return 0, false
	}
	if m[rp.X][rp.Y] == 'E' {
		oldScore, ok := memo[rp]
		if !ok || currentScore < oldScore {
			memo[rp] = currentScore
		}
		fmt.Printf("oldScore: %d, currentScore:%d // Found exit at %d, %d\n", oldScore, currentScore, rp.X, rp.Y)
		return currentScore, true
	} else if m[rp.X][rp.Y] == 'X' {
		oldScore, ok := memo[rp]
		if !ok || currentScore < oldScore {
			memo[rp] = currentScore
		} else {
			return oldScore, false
		}
	}
	m[rp.X][rp.Y] = 'X'
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (i == 0 || j == 0) && !(i == 0 && j == 0) {
				newPos := ReindeerPos{X: rp.X + i, Y: rp.Y + j}
				addendum := 1001
				if i == dx && j == dy {
					addendum = 1
				}
				_, stepHasSucceeded := SolveMaze2(m, newPos, i, j, currentScore+addendum, memo)
				if stepHasSucceeded {
					hasWayOut = true
				}
			}
		}

	}

	return 0, hasWayOut
}

// 397328 too high

func SolveMaze(m [][]rune, rp ReindeerPos, dx int, dy int, memo map[ReindeerPos]int) (int, bool) {
	minScore := 99999999
	hasWayOut := false
	prevChar := m[rp.X][rp.Y]
	bestDx := 0
	bestDy := 0
	//memoizedPos, ok := memo[rp]
	//if !ok {

	if m[rp.X][rp.Y] == 'E' {
		fmt.Printf("Found exit at %d, %d\n", rp.X, rp.Y)
		//memo[rp] = 0
		return 0, true
	} else if m[rp.X][rp.Y] == '#' || m[rp.X][rp.Y] == 'X' {
		return 0, false
	} else {
		m[rp.X][rp.Y] = 'X'
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if (i == 0 || j == 0) && !(i == 0 && j == 0) {
					stepHasSucceeded := false
					score := -1
					supplementScore := 1001
					if i == dx && j == dy {
						supplementScore = 1
					}
					newPos := ReindeerPos{X: rp.X + i, Y: rp.Y + j}
					score, stepHasSucceeded = SolveMaze(m, newPos, i, j, memo)
					score += supplementScore
					if stepHasSucceeded {
						hasWayOut = true
						//fmt.Printf("Found way out with Score %d\n", score)
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
			m[rp.X][rp.Y] = prevChar
		}

	} else {
		m[rp.X][rp.Y] = '.'
	}

	//memo[rp] = minScore Can t make this work
	//if !hasWayOut {
	//	memo[rp] = -1 * minScore
	//}

	/*} else {
		minScore = memoizedPos
		hasWayOut = true
		if minScore < 0 {
			minScore = -1 * minScore
			hasWayOut = false
		}
	}*/

	return minScore, hasWayOut
}

func main() {
	//filePath := "inputs/day16-sample11048.txt"
	//filePath := "inputs/day16-sample7036.txt"
	filePath := "inputs/day16.txt"
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	m, _, _ := utils.LinesTo2DArray(lines)
	memo := make(map[ReindeerPos]int)
	rp := GetReindeerPos(m)
	utils.PrintRune2DArray(m)
	fmt.Printf("Reindeer at %d, %d\n", rp.X, rp.Y)
	score, _ := SolveMaze2(m, rp, 0, 1, 0, memo)
	utils.PrintRune2DArray(m)
	fmt.Printf("%d\n", score)

}
