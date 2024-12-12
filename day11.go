package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
)

type StoneTuple struct {
	Stone int
	NRec  int
}

func noRuleApply(stones []int, i int) bool {
	res := true
	numStr := fmt.Sprintf("%d", stones[i])
	if stones[i] == 0 {
		res = false
	} else if len(numStr)%2 == 0 {
		res = false
	}
	return res
}

func noRuleApplyOnStone(stone int) bool {
	res := true
	numStr := fmt.Sprintf("%d", stone)
	if stone == 0 {
		res = false
	} else if len(numStr)%2 == 0 {
		res = false
	}
	return res
}

func ruleZero(stones []int, i int) []int {
	if stones[i] == 0 {
		stones[i] = 1
	}
	return stones
}

func ruleEvenOnStone(stone int) []int {
	//fmt.Printf("Checking ruleEven on %d\n", stones[i])
	numStr := fmt.Sprintf("%d", stone)
	stones := []int{}
	lenNum := len(numStr)
	if lenNum%2 == 0 {
		firstHalfStr := numStr[:lenNum/2]
		secondHalfStr := numStr[lenNum/2:]
		firstHalf, err1 := strconv.Atoi(firstHalfStr)
		if err1 != nil {
			fmt.Printf("Error converting %s to int", firstHalfStr)
		}
		secondHalf, err2 := strconv.Atoi(secondHalfStr)
		if err2 != nil {
			fmt.Printf("Error converting %s to int", firstHalfStr)
		}
		//fmt.Printf("Splitting %d into %d and %d\n", stones[i], firstHalf, secondHalf)
		stones = append(stones, firstHalf)
		stones = append(stones, secondHalf)
	} else {
		stones = append(stones, stone)
	}
	return stones
}

func ruleEven(stones []int, i int) ([]int, int) {
	//fmt.Printf("Checking ruleEven on %d\n", stones[i])
	numStr := fmt.Sprintf("%d", stones[i])
	lenNum := len(numStr)
	if lenNum%2 == 0 {
		firstHalfStr := numStr[:lenNum/2]
		secondHalfStr := numStr[lenNum/2:]
		firstHalf, err1 := strconv.Atoi(firstHalfStr)
		if err1 != nil {
			fmt.Printf("Error converting %s to int", firstHalfStr)
		}
		secondHalf, err2 := strconv.Atoi(secondHalfStr)
		if err2 != nil {
			fmt.Printf("Error converting %s to int", firstHalfStr)
		}
		//fmt.Printf("Splitting %d into %d and %d\n", stones[i], firstHalf, secondHalf)
		stones = append(stones[:i], append([]int{firstHalf, secondHalf}, stones[i+1:]...)...)
		i += 1
	}
	return stones, i
}

func applyRuleOnSingleStone(stone int) []int {
	stones := []int{}
	if noRuleApplyOnStone(stone) {
		stones = append(stones, stone*2024)
	} else if stone == 0 {
		stones = append(stones, 1)
	} else {
		stones = ruleEvenOnStone(stone)
	}
	return stones
}

func recursiveEleven(stones []int, nRec int, memo map[StoneTuple]int) int {
	res := 0
	if nRec == 0 {
		return len(stones)
	} else {
		leftMost := stones[0]
		newArray := []int{}
		st := StoneTuple{Stone: leftMost, NRec: nRec}
		if memo[st] != 0 {
			res += memo[st]
		} else {
			newArray = applyRuleOnSingleStone(leftMost)
			resStone := recursiveEleven(newArray, nRec-1, memo)
			memo[st] = resStone
			res += resStone
		}
		if len(stones) > 1 {
			remainder := stones[1:]
			res += recursiveEleven(remainder, nRec, memo)
		}
	}
	return res
}

func main() {
	filePath := "inputs/day11.txt"
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	memo := make(map[StoneTuple]int)
	count := 0
	for _, line := range lines {
		stones, _ := utils.LineToIntArray(line)
		count += recursiveEleven(stones, 75, memo)
		/*fmt.Printf("%v\n", stones)
		for nBlinks := 0; nBlinks < 25; nBlinks++ {
			for i := 0; i < len(stones); i++ {
				if noRuleApply(stones, i) {
					stones[i] *= 2024
				} else if stones[i] == 0 {
					stones = ruleZero(stones, i)
				} else {
					if i < len(stones) {
						stones, i = ruleEven(stones, i)
					}
				}

			}
			//fmt.Printf("%v\n", stones)
		}
		count = len(stones)*/
	}

	fmt.Println(count)
}
