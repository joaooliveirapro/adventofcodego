package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadInput(filename string) []string {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Remove leading/trailing spaces
		if len(line) != 0 {                       // Skip empty lines
			lines = append(lines, line)
		}
	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	return lines
}
