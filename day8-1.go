package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
)

func main() {
	filePath := "inputs/day8.txt"
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	m, sizeX, sizeY := utils.LinesTo2DArray(lines)
	fmt.Printf("Size: %d, %d\n", sizeX, sizeY)
	antennas := make(map[rune][]utils.AntennaPosition)
	count := 0
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			if m[i][j] != '.' {
				ap := utils.AntennaPosition{X: i, Y: j}
				antennas[m[i][j]] = append(antennas[m[i][j]], ap)
			}
		}
	}
	utils.PrintRune2DArray(m)
	for k, v := range antennas {
		fmt.Printf("Antenna %c: %v\n", k, v)
		utils.MarkAntinodes(m, v, sizeX, sizeY)
	}
	utils.PrintRune2DArray(m)
	count = utils.CountAntinodes(m, sizeX, sizeY)
	fmt.Println(count)

}
