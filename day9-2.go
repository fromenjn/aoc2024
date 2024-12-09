package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
)

func main() {
	filePath := "inputs/day9.txt"
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	line := lines[0]
	diskUnits := utils.StringToDiskUnits(line)
	//utils.PrintDiskUnits(diskUnits)
	utils.DefragmentDisk(diskUnits)
	count := utils.PerformChecksum(diskUnits)
	//utils.PrintDiskUnits(diskUnits)
	fmt.Println(count)

}
