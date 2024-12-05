package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func isUpdateOrdered(update []int, rulesIsBefore map[int][]int, rulesIsAfter map[int][]int) bool {
	for i := 0; i < len(update); i++ {
		cur := update[i]
		for j := i + 1; j < len(update); j++ {
			comp := update[j]
			if utils.ArrayContains(rulesIsBefore[comp], cur) {
				return false
			} else if utils.ArrayContains(rulesIsAfter[cur], comp) {
				return false
			}
		}
	}
	return true
}

func orderUpdate(update []int, rulesIsBefore map[int][]int, rulesIsAfter map[int][]int) []int {
	for !isUpdateOrdered(update, rulesIsBefore, rulesIsAfter) {
		for i := 0; i < len(update); i++ {
			cur := update[i]
			for j := i + 1; j < len(update); j++ {
				comp := update[j]
				if utils.ArrayContains(rulesIsBefore[comp], cur) {
					update[i], update[j] = update[j], update[i]
				} else if utils.ArrayContains(rulesIsAfter[cur], comp) {
					update[i], update[j] = update[j], update[i]
				}
			}
		}
	}
	//fmt.Printf("Ordered update: %v\n", update)
	return update
}

func processUpdate2(update []int, rulesIsBefore map[int][]int, rulesIsAfter map[int][]int) int {
	for i := 0; i < len(update); i++ {
		cur := update[i]
		for j := i + 1; j < len(update); j++ {
			comp := update[j]
			if utils.ArrayContains(rulesIsBefore[comp], cur) {
				//fmt.Printf("Error in %v: %d is before %d\n", update, comp, cur)
				return utils.GetMiddleElement(orderUpdate(update, rulesIsBefore, rulesIsAfter))
			} else if utils.ArrayContains(rulesIsAfter[cur], comp) {
				//fmt.Printf("Error in %v: %d is after %d\n", update, comp, cur)
				return utils.GetMiddleElement(orderUpdate(update, rulesIsBefore, rulesIsAfter))
			}
		}
	}
	return 0
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
					count += processUpdate2(update, rulesIsBefore, rulesIsAfter)
				}
			}
		}
	}

	fmt.Println(count)

}
