package utils

type AntennaPosition struct {
	X int // X coordinate
	Y int // Y coordinate
}

func MarkAntinode(m [][]rune, sizeX int, sizeY int, x int, y int) {
	if x < 0 || x >= sizeX || y < 0 || y >= sizeY {
		return
	}
	m[x][y] = '#'
}

func CountAntinodes(m [][]rune, sizeX int, sizeY int) int {
	count := 0
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			if m[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

func MarkAntinodes(m [][]rune, antennas []AntennaPosition, sizeX int, sizeY int) {
	for i := 0; i < len(antennas)-1; i++ {
		antenna1 := antennas[i]
		for j := i + 1; j < len(antennas); j++ {
			antenna2 := antennas[j]
			diffX := antenna2.X - antenna1.X
			diffY := antenna2.Y - antenna1.Y
			MarkAntinode(m, sizeX, sizeY, antenna2.X+diffX, antenna2.Y+diffY)
			MarkAntinode(m, sizeX, sizeY, antenna1.X-diffX, antenna1.Y-diffY)
		}
	}
}

func MarkHarmonicsAntinodes(m [][]rune, antennas []AntennaPosition, sizeX int, sizeY int) {
	for i := 0; i < len(antennas)-1; i++ {
		antenna1 := antennas[i]
		for j := i + 1; j < len(antennas); j++ {
			antenna2 := antennas[j]
			diffX := antenna2.X - antenna1.X
			diffY := antenna2.Y - antenna1.Y
			MarkAntinode(m, sizeX, sizeY, antenna2.X, antenna2.Y)
			MarkAntinode(m, sizeX, sizeY, antenna1.X, antenna1.Y)
			for k := 1; k < 100; k++ { //not very nice, but it works
				MarkAntinode(m, sizeX, sizeY, antenna2.X+diffX*k, antenna2.Y+diffY*k)
				MarkAntinode(m, sizeX, sizeY, antenna1.X-diffX*k, antenna1.Y-diffY*k)
			}
		}
	}
}
