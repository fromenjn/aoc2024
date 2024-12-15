package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strings"
)

type RP struct {
	X int
	Y int
}

type RD struct {
	X int
	Y int
}

func MakeRDFromChar(c rune) RD {
	switch c {
	case '^':
		//println("UP")
		return RD{X: -1, Y: 0}
	case 'v':
		//println("DOWN")
		return RD{X: 1, Y: 0}
	case '<':
		//println("LEFT")
		return RD{X: 0, Y: -1}
	case '>':
		//println("RIGHT")
		return RD{X: 0, Y: 1}
	}
	return RD{X: 0, Y: 0}
}

func (r *RP) Move(m [][]rune, rd RD) bool {
	curChar := m[r.X][r.Y]
	nextPos := RP{X: r.X + rd.X, Y: r.Y + rd.Y}

	if m[nextPos.X][nextPos.Y] == '#' {
		return false
	} else if m[nextPos.X][nextPos.Y] == '.' {
		m[r.X][r.Y] = '.'
		m[nextPos.X][nextPos.Y] = curChar
		return true
	} else if m[nextPos.X][nextPos.Y] == 'O' {
		res := nextPos.Move(m, rd)
		if res {
			m[r.X][r.Y] = '.'
			m[nextPos.X][nextPos.Y] = curChar
		}
		return res
	} else if m[nextPos.X][nextPos.Y] == '[' {
		otherHalf := RP{X: nextPos.X, Y: nextPos.Y + 1}
		res1 := nextPos.Move(m, rd)
		res2 := true
		if rd.Y != 1 {
			res2 = otherHalf.Move(m, rd)
		}
		if res1 && res2 {
			m[r.X][r.Y] = '.'
			m[nextPos.X][nextPos.Y] = curChar
			return true
		}
		return false
	} else if m[nextPos.X][nextPos.Y] == ']' {
		otherHalf := RP{X: nextPos.X, Y: nextPos.Y - 1}
		res1 := nextPos.Move(m, rd)
		res2 := true
		if rd.Y != -1 {
			res2 = otherHalf.Move(m, rd)
		}
		if res1 && res2 {
			m[r.X][r.Y] = '.'
			m[nextPos.X][nextPos.Y] = curChar
			return true
		}
		return false
	}
	fmt.Printf("Error CanMove: %c (%d,%d)\n", m[nextPos.X][nextPos.Y], nextPos.X, nextPos.Y)
	return false
}

func CountScore(m [][]rune) int {
	score := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 'O' || m[i][j] == '[' {
				boxScore := 100*i + j
				score += boxScore
			}
		}
	}
	return score
}

func GetRPos(m [][]rune) RP {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == '@' {
				return RP{X: i, Y: j}
			}
		}
	}
	fmt.Printf("Error GetRPos could not find @\n")
	return RP{X: -1, Y: -1}
}

func MakeDoubleMap(lines []string) []string {
	newLines := make([]string, 0)
	for _, line := range lines {
		newLine := ""
		for _, c := range line {
			if c == '@' {
				newLine += "@."
			} else if c == 'O' {
				newLine += "[]"
			} else if c == '#' {
				newLine += "##"
			} else if c == '.' {
				newLine += ".."
			} else {
				fmt.Printf("Error MakeDoubleMap: %c\n", c)
			}

		}
		newLines = append(newLines, newLine)
	}
	return newLines

}

func main() {
	filePath := "inputs/day15-sample3.txt"

	lines, err := utils.ReadFile(filePath)
	lm := make([]string, 0)
	ld := make([]RD, 0)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			lm = append(lm, line)
		} else {
			for _, c := range line {
				ld = append(ld, MakeRDFromChar(c))
			}
		}
	}
	m, sizeX, sizeY := utils.LinesTo2DArray(lm)
	m2, _, _ := utils.LinesTo2DArray(MakeDoubleMap(lm))
	fmt.Printf("Size: %d, %d\n", sizeX, sizeY)
	utils.PrintRune2DArray(m)
	fmt.Printf("------------\n")
	utils.PrintRune2DArray(m2)
	fmt.Printf("------------\n")
	rp := GetRPos(m)
	rp2 := GetRPos(m2)
	for i := 0; i < len(ld); i++ {
		if rp.Move(m, ld[i]) {
			rp = RP{X: rp.X + ld[i].X, Y: rp.Y + ld[i].Y}
		}
		if rp2.Move(m2, ld[i]) {
			rp2 = RP{X: rp2.X + ld[i].X, Y: rp2.Y + ld[i].Y}
		}
		utils.PrintRune2DArray(m2)
		fmt.Printf("------------\n")
	}
	fmt.Println(CountScore(m))
	fmt.Println(CountScore(m2))
}