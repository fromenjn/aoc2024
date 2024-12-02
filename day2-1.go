package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func isSafe(report []int) bool {
	isIncreasing := false
	if len(report) < 2 {
		return false
	}
	firstDiff := report[1] - report[0]
	if firstDiff > 0 {
		isIncreasing = true
	}
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if utils.IsDifferenceSafe(diff) {
			if isIncreasing && diff < 0 {
				return false
			}
			if !isIncreasing && diff > 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func main() {

	filePath := "inputs/day2.txt"
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	reports := [][]int{}
	for _, line := range lines {
		report := []int{}
		numbers := strings.Fields(line)
		for i := 0; i < len(numbers); i++ {
			num, err := strconv.Atoi(numbers[i])
			if err != nil {
				fmt.Printf("Skipping line with invalid numbers: %s\n", line)
				continue
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	cumulativeRes := 0
	for i := 0; i < len(reports); i++ {
		if isSafe(reports[i]) {
			cumulativeRes += 1
		}
	}
	fmt.Println(cumulativeRes)

}
