package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
)

type TrailPosition struct {
	X int
	Y int
}

type Trail struct {
	Start TrailPosition
	End   TrailPosition
}

//898 too high

func compareTrails(t1 Trail, t2 Trail) bool {
	if t1.Start.X == t2.Start.X && t1.Start.Y == t2.Start.Y && t1.End.X == t2.End.X && t1.End.Y == t2.End.Y {
		return true
	}
	return false
}

func recursiveTrail(m [][]rune, x int, y int, expected rune, trailHead TrailPosition, trails []Trail, part2 bool) (int, []Trail) {
	count := 0
	if x < 0 || x >= len(m) || y < 0 || y >= len(m[0]) {
		return 0, trails // Discard trails that go out of the map
	}
	currentRune := m[x][y] // check current rune if on the map
	if currentRune == expected {
		if expected == '9' { //terminal condition
			//fmt.Printf("Found 9 at %d, %d\n", x, y)
			trailEnd := TrailPosition{X: x, Y: y}
			trail := Trail{Start: trailHead, End: trailEnd}
			if !part2 {
				for _, t := range trails { // Discard trails that have a same start and end
					if compareTrails(t, trail) {
						return 0, trails
					}
				}
				trails = append(trails, trail)
			}
			return 1, trails
		} else {
			newExpected := expected + 1
			res1, trails := recursiveTrail(m, x+1, y, newExpected, trailHead, trails, part2)
			res2, trails := recursiveTrail(m, x, y+1, newExpected, trailHead, trails, part2)
			res3, trails := recursiveTrail(m, x-1, y, newExpected, trailHead, trails, part2)
			res4, trails := recursiveTrail(m, x, y-1, newExpected, trailHead, trails, part2)
			count += res1 + res2 + res3 + res4
			return count, trails
		}
	}
	return 0, trails
}

func main() {
	filePath := "inputs/day10.txt"
	part2 := true
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	m, sizeX, sizeY := utils.LinesTo2DArray(lines)
	trails := make([]Trail, 0)
	count := 0
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			//print(string(m[i][j]))
			charToFind := '0'
			if m[i][j] == charToFind {
				trailHead := TrailPosition{X: i, Y: j}
				newChar := charToFind + 1
				res1, trails := recursiveTrail(m, i+1, j, newChar, trailHead, trails, part2)
				res2, trails := recursiveTrail(m, i, j+1, newChar, trailHead, trails, part2)
				res3, trails := recursiveTrail(m, i-1, j, newChar, trailHead, trails, part2)
				res4, trails := recursiveTrail(m, i, j-1, newChar, trailHead, trails, part2)
				count += res1 + res2 + res3 + res4
			}
		}
	}

	fmt.Println(count)

}
