package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	filePath := "inputs/day6.txt"
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	m, sizeX, sizeY := utils.LinesTo2DArray(lines)
	fmt.Printf("Size: %d, %d\n", sizeX, sizeY)
	right := utils.GuardDirection{
		X:         0,
		Y:         1,
		Character: '>',
		Direction: "Right",
	}
	left := utils.GuardDirection{
		X:         0,
		Y:         -1,
		Character: '<',
		Direction: "Left",
	}
	up := utils.GuardDirection{
		X:         -1,
		Y:         0,
		Character: '^',
		Direction: "Up",
	}
	down := utils.GuardDirection{
		X:         1,
		Y:         0,
		Character: 'V',
		Direction: "Down",
	}
	directions := []utils.GuardDirection{right, down, left, up}
	//position := utils.GuardPosition{X: 0, Y: 0}
	startDirection := utils.GuardDirection{}
	startPosition := utils.GuardPosition{}
	startHasBeenFound := false
	count := 0
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			for d := 0; d < len(directions); d++ {
				if !startHasBeenFound && m[i][j] == directions[d].Character {
					startHasBeenFound = true
					startDirection = directions[d]
					//position = utils.GuardPosition{X: i, Y: j}
					startPosition = utils.GuardPosition{X: i, Y: j}
					fmt.Printf("Found %c (%s) at %d, %d\n", m[i][j], startDirection.Direction, i, j)
					m[i][j] = '.'
					count = utils.GuardWalk(m, sizeX, sizeY, startPosition, startDirection, count, directions)
				}
			}
		}
	}

	countObstacles := 0

	//For all possible positions
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			// As we have only one shot, the only possible obstacle is ON the walked path
			// But not on the starting position
			if m[i][j] == 'X' && !(i == startPosition.X && j == startPosition.Y) {
				newMap, _, _ := utils.LinesTo2DArray(lines)
				newMap[i][j] = '#'
				newMap[startPosition.X][startPosition.Y] = '.'
				if utils.DoesGuardWalkInLoop(newMap, sizeX, sizeY, startPosition, startDirection, directions) {
					countObstacles++
				}
				newMap[i][j] = '.'
			}
		}
	}

	utils.PrintRune2DArray(m)
	fmt.Printf("Count original: %d\n", count)
	fmt.Printf("Count possible obstacles: %d\n", countObstacles)
}
