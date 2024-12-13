package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strings"
)

type ClawPos struct {
	X int64
	Y int64
}

type ClawButton struct {
	Label          string // A or B
	Pos            ClawPos
	Cost           int64 //
	RemainingShots int64
}

func (cb ClawButton) String() string {
	return fmt.Sprintf("Button %s: %d, %d", cb.Label, cb.Pos.X, cb.Pos.Y)
}

type ClawGame struct {
	ClawButtons      []*ClawButton
	PrizeRelativePos *ClawPos
	ClawCost         int64
}

func (cg *ClawGame) DeepCopy() *ClawGame {
	clawButtonsCopy := make([]*ClawButton, len(cg.ClawButtons))
	for i, button := range cg.ClawButtons {
		clawButtonsCopy[i] = &ClawButton{
			Label:          button.Label,
			Pos:            ClawPos{X: button.Pos.X, Y: button.Pos.Y},
			Cost:           button.Cost,
			RemainingShots: button.RemainingShots,
		}
	}
	var clawPosCopy *ClawPos
	if cg.PrizeRelativePos != nil {
		clawPosCopy = &ClawPos{
			X: cg.PrizeRelativePos.X,
			Y: cg.PrizeRelativePos.Y,
		}
	}
	return &ClawGame{
		ClawButtons:      clawButtonsCopy,
		PrizeRelativePos: clawPosCopy,
		ClawCost:         cg.ClawCost,
	}
}

func (cg ClawGame) String() string {
	for _, b := range cg.ClawButtons {
		fmt.Printf("%s\n", b)
	}
	return fmt.Sprintf("Prize: %d, %d", cg.PrizeRelativePos.X, cg.PrizeRelativePos.Y)
}

func MakeClawButton(label string, pos ClawPos) *ClawButton {
	if label == "A" {
		return &ClawButton{Label: label, Pos: pos, Cost: 3, RemainingShots: 100}
	} else if label == "B" {
		return &ClawButton{Label: label, Pos: pos, Cost: 1, RemainingShots: 100}
	}
	return nil
}

func MakeClawGame(ButtonA *ClawButton, ButtonB *ClawButton, PrizePos *ClawPos) *ClawGame {
	return &ClawGame{ClawButtons: []*ClawButton{ButtonA, ButtonB}, PrizeRelativePos: PrizePos, ClawCost: 0}
}

func MakeClawGamesFromStrings(input []string) []*ClawGame {
	var buttonA, buttonB *ClawButton
	var clawPos *ClawPos
	var clawGames []*ClawGame

	for _, line := range input {
		//fmt.Printf("Line: %s\n", line)
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Button A:") {
			var x, y int64
			fmt.Sscanf(line, "Button A: X+%d, Y+%d", &x, &y)
			buttonA = MakeClawButton("A", ClawPos{X: x, Y: y})
		} else if strings.HasPrefix(line, "Button B:") {
			var x, y int64
			fmt.Sscanf(line, "Button B: X+%d, Y+%d", &x, &y)
			buttonB = MakeClawButton("B", ClawPos{X: x, Y: y})
		} else if strings.HasPrefix(line, "Prize:") {
			var x, y int64
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &x, &y)
			clawPos = &ClawPos{X: x, Y: y}
		}

		if buttonA != nil && buttonB != nil && clawPos != nil {
			clawGames = append(clawGames, MakeClawGame(buttonA, buttonB, clawPos))
			buttonA = nil
			buttonB = nil
			clawPos = nil
		}
	}

	return clawGames
}

func recursiveClawGame(game *ClawGame, currentBestSolution int64) (int64, bool) {

	hasSolution := false
	if game.PrizeRelativePos.X < 0 || game.PrizeRelativePos.Y < 0 {
		return currentBestSolution, false // If we have passed the prize, stop
	}
	if game.ClawCost > currentBestSolution {
		return currentBestSolution, false // If there is already a better solution, stop
	}
	for _, button := range game.ClawButtons { //If one of the buttons is out of shots, stop
		if button.RemainingShots < 0 {
			return currentBestSolution, false
		}
	}
	if game.PrizeRelativePos.X == 0 && game.PrizeRelativePos.Y == 0 {
		currentBestSolution = game.ClawCost
		return currentBestSolution, true
	}
	for _, button := range game.ClawButtons {
		//Test button
		fmt.Printf("Testing button %s(%d shots left) for Prize in [%d,%d]\n", button.Label, button.RemainingShots, game.PrizeRelativePos.X, game.PrizeRelativePos.Y)
		button.RemainingShots--
		game.PrizeRelativePos.X -= button.Pos.X
		game.PrizeRelativePos.Y -= button.Pos.Y
		game.ClawCost += button.Cost
		res, subGameHasSolution := recursiveClawGame(game, currentBestSolution)
		if subGameHasSolution && res < currentBestSolution {
			hasSolution = true
			currentBestSolution = res
		}
		// Restore button
		button.RemainingShots++
		game.PrizeRelativePos.X += button.Pos.X
		game.PrizeRelativePos.Y += button.Pos.Y
		game.ClawCost -= button.Cost
	}
	return currentBestSolution, hasSolution
}

func naiveClawGame(game *ClawGame, currentBestSolution int64) (int64, bool) {
	buttonA := game.ClawButtons[0]
	buttonB := game.ClawButtons[1]
	hasSolution := false
	for i := int64(0); i < buttonA.RemainingShots; i++ {
		cost := i * buttonA.Cost
		newRelativePos := ClawPos{X: game.PrizeRelativePos.X - i*buttonA.Pos.X, Y: game.PrizeRelativePos.Y - i*buttonA.Pos.Y}
		if newRelativePos.X%buttonB.Pos.X == 0 && newRelativePos.Y%buttonB.Pos.Y == 0 {
			testX := newRelativePos.X / buttonB.Pos.X
			testY := newRelativePos.Y / buttonB.Pos.Y
			if testX == testY && testX <= buttonB.RemainingShots {
				cost += testX * buttonB.Cost
				if cost < currentBestSolution {
					hasSolution = true
					currentBestSolution = cost
				}
			}
		}
	}
	return currentBestSolution, hasSolution
}

func naiveClawGameNoLimit(game *ClawGame, currentBestSolution int64) (int64, bool) {
	buttonA := game.ClawButtons[0]
	buttonB := game.ClawButtons[1]
	hasSolution := false
	stepAX := buttonA.Pos.X
	stepAY := buttonA.Pos.Y
	stepBX := buttonB.Pos.X
	stepBY := buttonB.Pos.Y
	maxStepsA := utils.Min(game.PrizeRelativePos.Y/stepAY+1, game.PrizeRelativePos.X/stepAX+1)
	maxStepsB := utils.Min(game.PrizeRelativePos.Y/stepBY+1, game.PrizeRelativePos.X/stepBX+1)
	remainingX := game.PrizeRelativePos.X
	remainingY := game.PrizeRelativePos.Y
	if maxStepsA < maxStepsB {
		for i := int64(0); i < maxStepsA; i++ {
			if remainingX%stepBX == 0 && remainingY%stepBY == 0 {
				multiplierX := remainingX / stepBX
				multiplierY := remainingY / stepBY
				// Both multipliers must match for a valid solution
				if multiplierX == multiplierY {
					cost := i*buttonA.Cost + multiplierX*buttonB.Cost
					if cost < currentBestSolution {
						hasSolution = true
						currentBestSolution = cost
					}
				}
			}
			remainingX -= stepAX
			remainingY -= stepAY
		}
	} else {
		for i := int64(0); i < maxStepsB; i++ {
			if remainingX%stepAX == 0 && remainingY%stepAY == 0 {
				multiplierX := remainingX / stepAX
				multiplierY := remainingY / stepAY
				// Both multipliers must match for a valid solution
				if multiplierX == multiplierY {
					cost := i*buttonB.Cost + multiplierX*buttonA.Cost
					if cost < currentBestSolution {
						hasSolution = true
						currentBestSolution = cost
					}
				}
			}
			remainingX -= stepBX
			remainingY -= stepBY
		}
	}
	return currentBestSolution, hasSolution
}

func main() {
	filePath := "inputs/day13.txt"
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	count := int64(0)
	ClawGames := MakeClawGamesFromStrings(lines)
	for _, game := range ClawGames {
		//Part2
		//game.PrizeRelativePos.X += 10000000000000
		//game.PrizeRelativePos.Y += 10000000000000

		fmt.Printf("%s\n", game)
		res, hasSolution := naiveClawGameNoLimit(game, 90000000000000)
		if hasSolution {
			fmt.Printf("Solution: %d\n-----\n", res)
			count += res
		} else {
			fmt.Printf("No solution\n")
		}
	}

	fmt.Println(count)
}
