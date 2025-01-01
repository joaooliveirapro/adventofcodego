package main

import (
	"bufio"
	"fmt"
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

func getGuardPositionAndDirection(input *[]string) (int, int, string) {
	for y, line := range *input {
		for x, char := range line {
			c := string(char)
			if c == "^" || c == ">" || c == "v" || c == "<" {
				return x, y, c
			}
		}
	}
	return -1, -1, ""
}

func getCharAt(input *[]string, x, y int) string {
	// Check if y is within bounds of the input
	if y >= 0 && y < len(*input) {
		line := (*input)[y]
		// Check if x is within bounds of the line
		if x >= 0 && x < len(line) {
			return string(line[x])
		}
	}
	return "" // Return empty string if out of bounds
}

func replaceCharAt(input *[]string, x, y int, r string) []string {
	t := make([]string, len(*input)) // Preallocate slice for efficiency
	for yy, line := range *input {
		if yy == y {
			var builder strings.Builder
			builder.WriteString(line[:x])   // Add everything before the target character
			builder.WriteString(r)          // Replace target character
			builder.WriteString(line[x+1:]) // Add everything after the target character
			t[yy] = builder.String()
		} else {
			t[yy] = line // Copy the unchanged line
		}
	}
	return t
}

func MoveGuard(input *[]string) ([]string, bool) {
	x, y, dir := getGuardPositionAndDirection(input)
	newX, newY, newDir := -1, -1, dir // assume dir doesn't change
	outofbounds := false
	if x != -1 && y != -1 {
		// Use ^(up) >(right) v(down) <(left) to denote direction of guard
		c := ""
		if dir == "^" {
			c = getCharAt(input, x, y-1)
			newX, newY = x, y-1
		} else if dir == ">" {
			c = getCharAt(input, x+1, y)
			newX, newY = x+1, y
		} else if dir == "v" {
			c = getCharAt(input, x, y+1)
			newX, newY = x, y+1
		} else if dir == "<" {
			c = getCharAt(input, x-1, y)
			newX, newY = x-1, y
		}
		dirs := "^>v<"
		// obstacle; change dir
		if c == "#" {
			newDirInd := (strings.Index(dirs, dir) + 1) % len(dirs)
			newDir = string(dirs[newDirInd])
			newX, newY = x, y // Don't move guard to obstacle.
		}
		if c == "" {
			outofbounds = true
		}
		// replace guard with "."
		newInput := replaceCharAt(input, x, y, "X")
		// Move guard
		newInput = replaceCharAt(&newInput, newX, newY, newDir)
		return newInput, outofbounds
	}
	return []string{}, outofbounds
}

func day6_part1() {
	input := ReadInput("./day6/day6.input.txt")
	var done bool
	for {
		input, done = MoveGuard(&input)
		if done {
			x := 0
			for _, line := range input {
				x += strings.Count(line, "X")
			}
			fmt.Printf("Guard exited screen after visiting %d positions\n", x)
			break
		}
		var s strings.Builder
		for _, line := range input {
			for _, char := range line {
				charStr := string(char)
				if charStr == "." {
					s.WriteString(charStr)
				} else if charStr != "." && charStr != "#" {
					s.WriteString(fmt.Sprintf("\033[41m%s\033[0m", charStr))
				} else {
					s.WriteString(fmt.Sprintf("\033[33m%s\033[0m", charStr))
				}
			}
			s.WriteString("\n")
		}
		fmt.Print("\033[H\033[2J")
		fmt.Print(s.String())
		// time.Sleep(250 * time.Millisecond) (optional)
	}

}

func day6_part2() {

}

func main() {
	day6_part1()
	day6_part2()
}
