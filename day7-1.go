package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func checkEquation1(operands []int, expected int, o utils.EquationOperator) int {
	if len(operands) == 1 {
		if operands[0] == expected {
			return expected
		}
		return -1
	}
	newLen := len(operands) - 1
	newOperands := make([]int, newLen)
	if o == utils.Plus {
		newOperands[0] = operands[0] + operands[1]
	} else if o == utils.Multiply {
		newOperands[0] = operands[0] * operands[1]
	}
	for i := 1; i < newLen; i++ {
		if i < len(operands)-1 {
			newOperands[i] = operands[i+1]
		}
	}
	res := -1
	res = checkEquation1(newOperands, expected, utils.Plus)
	if res == -1 {
		res = checkEquation1(newOperands, expected, utils.Multiply)
	}
	return res
}

func main() {
	filePath := "inputs/day7.txt"
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	count := 0
	operands := []int{}
	for _, line := range lines {
		parts := strings.Split(line, ":")

		if len(parts) == 2 {
			expected, err1 := strconv.Atoi(parts[0])
			if err1 != nil {
				fmt.Printf("Skipping invalid number in line: %s\n", line)
				continue
			}
			operands, _ = utils.LineToIntArray(parts[1])
			//fmt.Printf("Operands: %v\n", operands)

			res := checkEquation1(operands, expected, utils.Plus)
			if res == -1 { // 2+2 == 2*2, do not count the result twice
				res = checkEquation1(operands, expected, utils.Multiply)
			}
			if res >= 0 {
				fmt.Printf("OK: %d\n", res)
				count += res
			}
		} else {
			fmt.Printf("Skipping line with invalid format: %s\n", line)
		}
	}
	fmt.Println(count)
}
