package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
)

func main() {
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
	walkedPositions := make([]utils.GuardPosition, 0)
	count := 0
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			for d := 0; d < len(directions); d++ {
				if m[i][j] == directions[d].Character {
					startDirection = directions[d]
					//position = utils.GuardPosition{X: i, Y: j}
					startPosition = utils.GuardPosition{X: i, Y: j}
					fmt.Printf("Found %c (%s) at %d, %d\n", m[i][j], startDirection.Direction, i, j)
					m[i][j] = '.'
					count = utils.GuardWalk(m, sizeX, sizeY, startPosition, startDirection, count, directions, walkedPositions)
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
				if utils.DoesGuardWalkInLoop(newMap, sizeX, sizeY, startPosition, startDirection, directions, walkedPositions) {
					countObstacles++
				}
			}
		}
	}

	utils.PrintRune2DArray(m)
	fmt.Printf("Count original: %d\n", count)
	fmt.Printf("Count possible obstacles: %d\n", countObstacles)

}
