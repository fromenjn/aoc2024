package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strings"
)

type RobotCoords struct {
	X int
	Y int
}

type Robot struct {
	Speed    RobotCoords
	Position RobotCoords
}

type RobotMap struct {
	sizeX  int
	sizeY  int
	Robots []*Robot
}

func (m *RobotMap) RobotMove(r *Robot) {
	r.Position.X += r.Speed.X
	r.Position.Y += r.Speed.Y
	r.Position.X = (r.Position.X + m.sizeX) % m.sizeX
	r.Position.Y = (r.Position.Y + m.sizeY) % m.sizeY
}

func (m *RobotMap) PlayTurn() {
	for _, r := range m.Robots {
		m.RobotMove(r)
	}
}

func (m *RobotMap) CountSafety() int {
	parts := int(2)
	sizeX := m.sizeX / parts
	sizeY := m.sizeY / parts
	addX, addY := int(0), int(0)
	if m.sizeX%parts != 0 {
		addX += 1
	}
	if m.sizeY%parts != 0 {
		addY += 1
	}
	scores := make([]int, 0)
	for i := int(0); i < parts; i++ {
		for j := int(0); j < parts; j++ {
			safetyScore := int(0)
			startX := i*sizeX + i*addX
			startY := j*sizeY + j*addY
			endX := startX + sizeX
			endY := startY + sizeY
			for _, r := range m.Robots {
				if r.Position.X >= startX && r.Position.X < endX && r.Position.Y >= startY && r.Position.Y < endY {
					//fmt.Printf("Robot at (%d, %d) is in quadrant (%d, %d)\n", r.Position.X, r.Position.Y, i, j)
					safetyScore += 1
				}
			}
			fmt.Printf("Quadrant (%d, %d) has safetyScore %d      | boundaries[(%d,%d), (%d,%d)]\n", i, j, safetyScore, startX, startY, endX, endY)
			scores = append(scores, safetyScore)
		}
	}
	totalScore := int(1)
	for _, s := range scores {
		totalScore *= s
	}
	return totalScore
}

func (m *RobotMap) Print() {
	matrix := make([][]rune, m.sizeX)
	for i := range matrix {
		matrix[i] = make([]rune, m.sizeY)
		for j := range matrix[i] {
			matrix[i][j] = '.' // Fill with '.'
		}
	}
	for _, r := range m.Robots {
		c := matrix[r.Position.X][r.Position.Y]
		if c == '.' {
			matrix[r.Position.X][r.Position.Y] = '1'
		} else if c >= '1' && c <= '8' {
			matrix[r.Position.X][r.Position.Y] = c + 1
		} else {
			fmt.Printf("Error when trying to udapte character %c at %d, %d\n", c, r.Position.X, r.Position.Y)
		}

	}

	for j := 0; j < m.sizeY; j++ {
		for i := 0; i < m.sizeX; i++ {
			c := matrix[i][j]
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func MakeRobotMap(Robots []*Robot, sizeX int, sizeY int) *RobotMap {
	return &RobotMap{Robots: Robots, sizeX: sizeX, sizeY: sizeY}
}

func MakeRobotMapFromStrings(input []string, sizeX int, sizeY int) *RobotMap {
	Robots := make([]*Robot, 0)
	for _, line := range input {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "p=") {
			var px, py, vx, vy int
			fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
			r := &Robot{Speed: RobotCoords{X: vx, Y: vy}, Position: RobotCoords{X: px, Y: py}}
			Robots = append(Robots, r)
		}
	}
	return MakeRobotMap(Robots, sizeX, sizeY)
}

// 11,7
//101,103

func main() {
	filePath := "inputs/day14.txt"
	sizeX := int(101)
	sizeY := int(103)
	//filePath := "inputs/day14-sample.txt"
	//sizeX := int(11)
	//sizeY := int(7)
	nTurns := int(100)
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	count := int(0)
	m := MakeRobotMapFromStrings(lines, sizeX, sizeY)
	m.Print()
	for i := int(0); i < nTurns; i++ {
		m.PlayTurn()
	}
	fmt.Printf("-----After %d turns:-----\n", nTurns)
	m.Print()
	count = m.CountSafety()

	fmt.Println(count)
}
