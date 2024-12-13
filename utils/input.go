package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadFile reads the contents of a file and returns it as a string slice.
func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return lines, nil
}

// ReadCommandLine reads inputs from the command line as a string slice.
func ReadCommandLine() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter inputs (end with an empty line):")
	var inputs []string

	for {
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1] // Remove newline character
		if input == "" {
			break
		}
		inputs = append(inputs, input)
	}

	return inputs
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsDifferenceSafe(i int) bool {
	if Abs(i) >= 1 && Abs(i) <= 3 {
		return true
	}
	return false
}

func LinesTo2DArray(lines []string) ([][]rune, int, int) {
	var result [][]rune
	y := len(lines) // Number of rows (y-axis)
	x := 0          // Number of columns (x-axis)

	for _, line := range lines {
		row := []rune(line)
		result = append(result, row)

		// Update the maximum width (x-axis) if necessary
		if len(row) > x {
			x = len(row)
		}
	}

	return result, x, y
}

func ArrayContains(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func LineToIntArray(line string) ([]int, error) {
	parts := strings.Fields(line)
	result := make([]int, 0, len(parts))
	for _, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", part)
		}
		result = append(result, num)
	}
	return result, nil
}

func ProcessCommaSeparatedIntegers(line string) ([]int, error) {
	parts := strings.Split(line, ",")
	result := make([]int, 0, len(parts))
	for _, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", part)
		}
		result = append(result, num)
	}
	return result, nil
}

func GetMiddleElement(arr []int) int {
	if len(arr)%2 == 0 {
		return arr[len(arr)/2]
	} else {
		return arr[len(arr)/2]
	}
}

func PrintRune2DArray(arr [][]rune) {
	for _, row := range arr {
		fmt.Println(string(row))
	}
}

func HasNeighbor(arr [][]rune, neighbor rune, x, y int) bool {
	sizeY := len(arr[0])
	sizeX := len(arr)
	if x > 0 && arr[x-1][y] == neighbor {
		return true
	}
	if x < sizeX-1 && arr[x+1][y] == neighbor {
		return true
	}
	if y > 0 && arr[x][y-1] == neighbor {
		return true
	}
	if y < sizeY-1 && arr[x][y+1] == neighbor {
		return true
	}
	return false
}

func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func Ppcm(a, b int64) int64 {
	if a == 0 || b == 0 {
		return 0
	}
	return a * b / Gcd(a, b)
}

func EuclideanDivision(a, b int64) (quotient, remainder int64) {
	if b == 0 {
		panic("division by zero")
	}
	quotient = a / b
	remainder = a % b
	// Adjust to ensure the remainder is non-negative
	if remainder < 0 {
		if b > 0 {
			quotient--
			remainder += b
		} else {
			quotient++
			remainder -= b
		}
	}
	return
}
