package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {

	filePath := "inputs/day1.txt"
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	list1 := []int{}
	list2 := []int{}
	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			fmt.Printf("Skipping invalid line: %s\n", line)
			continue
		}

		// Convert the numbers to integers
		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			fmt.Printf("Skipping line with invalid numbers: %s\n", line)
			continue
		}

		// Append to the respective lists
		list1 = append(list1, num1)
		list2 = append(list2, num2)

	}
	sort.Ints(list1)
	sort.Ints(list2)

	// Iterate over both lists and calculate the sum of differences
	intMap := make(map[int]int)

	for i := 0; i < len(list2); i++ {
		if res, ok := intMap[list2[i]]; ok {
			intMap[list2[i]] = res + 1
		} else {
			intMap[list2[i]] = 1
		}
	}

	cumulativeRes := 0
	for i := 0; i < len(list1); i++ {
		if res, ok := intMap[list1[i]]; ok {
			//fmt.Print(list1[i], " appears ", res, " times\n")
			cumulativeRes += list1[i] * res
		}
	}
	fmt.Println(cumulativeRes)

}
