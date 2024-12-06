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
	position := utils.GuardPosition{X: 0, Y: 0}
	direction := utils.GuardDirection{}
	walkedPositions := make([]utils.GuardPosition, 0)
	count := 0
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			for d := 0; d < len(directions); d++ {
				if m[i][j] == directions[d].Character {
					direction = directions[d]
					position.X = i
					position.Y = j
					fmt.Printf("Found %c (%s) at %d, %d\n", m[i][j], direction.Direction, i, j)
					m[i][j] = '.'
					count = utils.GuardWalk(m, sizeX, sizeY, position, direction, count, directions, walkedPositions)
				}
			}
		}
	}
	utils.PrintRune2DArray(m)
	fmt.Println(count)

}
