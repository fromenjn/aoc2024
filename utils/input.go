package utils

import (
	"bufio"
	"fmt"
	"os"
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
