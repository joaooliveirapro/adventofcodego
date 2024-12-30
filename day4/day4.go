package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput() [][]string {
	// Open the file
	file, err := os.Open("./day4/day4.input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var lines [][]string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Remove leading/trailing spaces
		if len(line) != 0 {                       // Skip empty lines
			lines = append(lines, strings.Split(line, ""))
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	return lines
}

func getLeft(input *[][]string, x, y int) string {
	// x, y - 0-based index
	/* Since we're looking for a 4-word letter only
	coordinates from x = 4 are valid. */
	if x > 2 {
		temp := strings.Builder{}
		for i := 0; i < 4; i++ {
			temp.Write([]byte((*input)[y][x-i]))
			if temp.Len() == 4 {
				break
			}
		}
		return temp.String()
	}
	return ""
}

func getRight(input *[][]string, x, y int) string {
	/* Since we're looking for a 4-word letter only
	coordinates that go UP to x - len(row) - 4 are valid. */
	if len((*input)[0])-x > 4 {
		temp := strings.Builder{}
		for i := 0; i < 4; i++ {
			temp.Write([]byte((*input)[y][x+i]))
			if temp.Len() == 4 {
				break
			}
		}
		return temp.String()
	}
	return ""
}

func getTop(input *[][]string, x, y int) string {
	/*There must be at least 4 letters starting from y (0,1,2,3)*/
	if y > 2 {
		temp := strings.Builder{}
		for i := 0; i < 4; i++ {
			temp.Write([]byte((*input)[y-i][x]))
			if temp.Len() == 4 {
				break
			}
		}
		return temp.String()
	}
	return ""
}

func getBottom(input *[][]string, x, y int) string {
	if len((*input))-y > 3 {
		temp := strings.Builder{}
		for i := 0; i < 4; i++ {
			temp.Write([]byte((*input)[y+i][x]))
			if temp.Len() == 4 {
				break
			}
		}
		return temp.String()
	}
	return ""
}

func day4_part1() {
	/*
		Approach:
		Treat the input as a matrix n * m
		and check all 8 compass neighbours for the word
		XMAS and SAMX
		track all of the ones found to avoid duplicates
	*/

	input := ReadInput()
	track := 0

	o := map[string]any{}

	for y, row := range input {
		for x, _ := range row {
			//Check left
			toLeft := getLeft(&input, x, y)
			if toLeft == "XMAS" || toLeft == "SAMX" {
				track++
				// for debugging
				o[fmt.Sprintf("%d,%d", x, y)] = "toLeft"
			}
			// check top
			toTop := getTop(&input, x, y)
			if toTop == "XMAS" || toTop == "SAMX" {
				track++
				// for debugging
				o[fmt.Sprintf("%d,%d", x, y)] = "toTop"
			}
			// check right
			toRight := getRight(&input, x, y)
			if toRight == "XMAS" || toRight == "SAMX" {
				track++
				// for debugging
				o[fmt.Sprintf("%d,%d", x, y)] = "toRight"
			}
			// check bottom
			toBottom := getBottom(&input, x, y)
			if toBottom == "XMAS" || toBottom == "SAMX" {
				track++
				// for debugging
				o[fmt.Sprintf("%d,%d", x, y)] = "toBottom"
			}

			// check top-left
			// check top-right
			// check bottom-right
			// check bottom-left

		}
	}
	fmt.Printf("track: %v\n", track)

	for k, v := range o {
		fmt.Printf("%-8s => %s\n", k, v)
	}

}

func day4_part2() {

}

func main() {
	day4_part1()
	day4_part2()
}
