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
