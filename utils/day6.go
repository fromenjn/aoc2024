package utils

type GuardDirection struct {
	X         int    // X coordinate (row)
	Y         int    // Y coordinate (column)
	Character rune   // Character representation
	Direction string // Direction as a string
}
type GuardPosition struct {
	X              int // X coordinate
	Y              int // Y coordinate
	SavedDirection GuardDirection
}

func GuardChangeDirection(direction GuardDirection, directions []GuardDirection) GuardDirection {
	for i := 0; i < len(directions); i++ {
		if direction.Character == directions[i].Character {
			if i == len(directions)-1 {
				return directions[0]
			}
			return directions[i+1]
		}
	}
	return directions[0]
}

func GuardMove(m [][]rune, position GuardPosition, direction GuardDirection) GuardPosition {
	return GuardPosition{X: position.X + direction.X, Y: position.Y + direction.Y}
}

func IsGuardOut(m [][]rune, position GuardPosition, sizeX int, sizeY int) bool {
	if position.X < 0 || position.X >= sizeX || position.Y < 0 || position.Y >= sizeY {
		return true
	}
	return false
}

func GuardWalk(m [][]rune, sizeX int, sizeY int, position GuardPosition, direction GuardDirection, count int, directions []GuardDirection, walkedPositions []GuardPosition) int {
	if m[position.X][position.Y] == '.' {
		m[position.X][position.Y] = 'X'
		count++
	}
	walkedPositions = append(walkedPositions, GuardPosition{X: position.X, Y: position.Y, SavedDirection: direction})
	futurePosition := GuardMove(m, position, direction)
	if IsGuardOut(m, futurePosition, sizeX, sizeY) {
		return count
	}
	for m[futurePosition.X][futurePosition.Y] == '#' {
		direction = GuardChangeDirection(direction, directions)
		futurePosition = GuardMove(m, position, direction)
		if IsGuardOut(m, futurePosition, sizeX, sizeY) {
			return count
		}
	}
	return GuardWalk(m, sizeX, sizeY, futurePosition, direction, count, directions, walkedPositions)
}

func DoesGuardWalkInLoop(m [][]rune, sizeX int, sizeY int, position GuardPosition, direction GuardDirection, directions []GuardDirection, walkedPositions []GuardPosition) bool {
	if m[position.X][position.Y] == '.' {
		m[position.X][position.Y] = 'X'
	}

	//If the guard returns to the same position AND direction, then he is be stuck in a loop
	newWalkedPosition := GuardPosition{X: position.X, Y: position.Y}
	for w := 0; w < len(walkedPositions); w++ {
		if walkedPositions[w].X == newWalkedPosition.X && walkedPositions[w].Y == newWalkedPosition.Y {
			if walkedPositions[w].SavedDirection.Character == direction.Character {
				return true
			}
		}
	}
	walkedPositions = append(walkedPositions, GuardPosition{X: position.X, Y: position.Y, SavedDirection: direction})
	futurePosition := GuardMove(m, position, direction)
	if IsGuardOut(m, futurePosition, sizeX, sizeY) {
		return false
	}
	for m[futurePosition.X][futurePosition.Y] == '#' {
		direction = GuardChangeDirection(direction, directions)
		futurePosition = GuardMove(m, position, direction)
		if IsGuardOut(m, futurePosition, sizeX, sizeY) {
			return false
		}
	}
	return DoesGuardWalkInLoop(m, sizeX, sizeY, futurePosition, direction, directions, walkedPositions)
}
