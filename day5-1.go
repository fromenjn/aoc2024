package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func processUpdate1(update []int, rulesIsBefore map[int][]int, rulesIsAfter map[int][]int) int {
	for i := 0; i < len(update); i++ {
		cur := update[i]
		for j := i + 1; j < len(update); j++ {
			comp := update[j]
			if utils.ArrayContains(rulesIsBefore[comp], cur) {
				return 0
			} else if utils.ArrayContains(rulesIsAfter[cur], comp) {
				return 0
			}
		}
	}
	//fmt.Printf("OK: %v\n", update)
	return utils.GetMiddleElement(update)
}

func main() {
	filePath := "inputs/day5.txt"
	rulesIsBefore := make(map[int][]int)
	rulesIsAfter := make(map[int][]int)
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	count := 0
	for _, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Printf("Skipping invalid numbers in line: %s\n", line)
				continue
			}
			//rulesIsBefore[num1] = append(rulesIsBefore[num1], num2)
			//fmt.Printf("%d|%d\n", num1, num2)
			rulesIsAfter[num2] = append(rulesIsAfter[num2], num1)
		} else {
			if strings.Contains(line, ",") {
				// Process the value as a comma-separated list of integers
				update, err := utils.ProcessCommaSeparatedIntegers(line)
				if err != nil {
					fmt.Printf("Error processing line: %s\n", line)
					continue
				} else {
					//fmt.Printf("Processing update: %v\n", update)
					count += processUpdate1(update, rulesIsBefore, rulesIsAfter)
				}
			}
		}
	}

	fmt.Println(count)

}
