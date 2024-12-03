package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//106784935 too high

func main() {
	filePath := "inputs/day3.txt"
	cleanPattern := `(?s)don't\([^)]*\).*?do\([^)]*\)` //`don't\([^)]*\).*?do\([^)]*\)`
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	reClean := regexp.MustCompile(cleanPattern)
	//filePath := "inputs/day2-sample.txt"
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	text := strings.Join(lines, "\n")
	cumulativeSum := 0
	cleanedText := reClean.ReplaceAllString(text, "")
	cleanedLines := strings.Split(cleanedText, "\n")

	for _, line := range cleanedLines {
		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			//fmt.Printf("Found: num1=%s, num2=%s\n", match[1], match[2])
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Error converting numbers:", err1, err2)
				continue
			}
			cumulativeSum += num1 * num2

		}

	}

	fmt.Println(cumulativeSum)

}
