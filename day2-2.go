package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func removeElement(report []int, i int) []int {
	// Create a new slice that does not share the same underlying array
	newReport := make([]int, 0, len(report)-1)     // Preallocate space for better performance
	newReport = append(newReport, report[:i]...)   // Append elements before index i
	newReport = append(newReport, report[i+1:]...) // Append elements after index i
	return newReport
}

func isReportSafe(report []int, isIncreasing bool, tolerance int) bool {
	problemDampener := 0
	if len(report) < 2 {
		fmt.Print("Report is too short")
		return false
	}
	hasIssue := false
	for i := 1; i < len(report); i++ {
		hasIssue = false
		diff := report[i] - report[i-1]
		if utils.IsDifferenceSafe(diff) {
			if isIncreasing && diff <= 0 {
				hasIssue = true
				problemDampener += 1
			} else if !isIncreasing && diff >= 0 {
				hasIssue = true
				problemDampener += 1
			}
		} else {
			hasIssue = true
			problemDampener += 1
		}
		if problemDampener > tolerance {
			return false
		} else if hasIssue {
			newReport := removeElement(report, i)
			newReport2 := removeElement(report, i-1)
			if isReportSafe(newReport, isIncreasing, tolerance-1) {
				return true
			} else if isReportSafe(newReport2, isIncreasing, tolerance-1) {
				return true
			} else {
				return false
			}
		}

	}
	return true
}

func main() {
	filePath := "inputs/day2.txt"
	//filePath := "inputs/day2-sample.txt"
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
	tolerance := 1
	for i := 0; i < len(reports); i++ {
		if isReportSafe(reports[i], true, tolerance) {
			fmt.Print(reports[i], " is Safe\n")
			cumulativeRes += 1
		} else if isReportSafe(reports[i], false, tolerance) {
			fmt.Print(reports[i], " is Safe\n")
			cumulativeRes += 1
		} else {
			fmt.Print(reports[i], " is Unsafe\n")
		}
	}
	fmt.Println(cumulativeRes)

}
