package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func ReadInput(filename string) [][]string {
	// Open the file
	file, err := os.Open(filename)
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

func getLeft(input *[][]string, x, y int) (string, [][]int) {
	// x, y - 0-based index
	/* Since we're looking for a 4-word letter only
	coordinates from x = 4 are valid. */
	coords := [][]int{}
	if x > 2 {
		temp := strings.Builder{}
		for i := 0; i < 4; i++ {
			temp.Write([]byte((*input)[y][x-i]))
			coords = append(coords, []int{x - i, y})
			if temp.Len() == 4 {
				break
			}
		}
		return temp.String(), coords
	}
	return "", coords
}

func getRight(input *[][]string, x, y int) (string, [][]int) {
	/* Since we're looking for a 4-word letter only
	coordinates that go UP to x - len(row) - 4 are valid. */
	coords := [][]int{}
	if len((*input)[0])-x >= 4 {
		temp := strings.Builder{}
		for i := 0; i < 4; i++ {
			temp.Write([]byte((*input)[y][x+i]))
			coords = append(coords, []int{x + i, y})
			if temp.Len() == 4 {
				break
			}
		}
		return temp.String(), coords
	}
	return "", coords
}

func getTop(input *[][]string, x, y int) (string, [][]int) {
	/*There must be at least 4 letters starting from y (0,1,2,3)*/
	coords := [][]int{}
	if y > 2 {
		temp := strings.Builder{}
		for i := 0; i < 4; i++ {
			temp.Write([]byte((*input)[y-i][x]))
			coords = append(coords, []int{x, y - i})
			if temp.Len() == 4 {
				break
			}
		}
		return temp.String(), coords
	}
	return "", coords
}

func getBottom(input *[][]string, x, y int) (string, [][]int) {
	coords := [][]int{}
	if len((*input))-y > 3 {
		temp := strings.Builder{}
		for i := 0; i < 4; i++ {
			temp.Write([]byte((*input)[y+i][x]))
			coords = append(coords, []int{x, y + i})
			if temp.Len() == 4 {
				break
			}
		}
		return temp.String(), coords
	}
	return "", coords
}

func getTopLeft(input *[][]string, x, y int) (string, [][]int) {
	coords := [][]int{}
	if x > 2 && y > 2 {
		temp := strings.Builder{}
		for i := 0; i < 4; i++ {
			temp.Write([]byte((*input)[y-i][x-i]))
			coords = append(coords, []int{x - i, y - i})
			if temp.Len() == 4 {
				break
			}
		}
		return temp.String(), coords
	}
	return "", coords
}

func getTopRight(input *[][]string, x, y int) (string, [][]int) {
	coords := [][]int{}
	if y > 2 && len((*input)[0])-x > 3 {
		temp := strings.Builder{}
		for i := 0; i < 4; i++ {
			temp.Write([]byte((*input)[y-i][x+i]))
			coords = append(coords, []int{x + i, y - i})
			if temp.Len() == 4 {
				break
			}
		}
		return temp.String(), coords
	}
	return "", coords
}

func getBottomRight(input *[][]string, x, y int) (string, [][]int) {
	coords := [][]int{}
	if len((*input))-y > 3 && len((*input)[0])-x > 3 {
		temp := strings.Builder{}
		for i := 0; i < 4; i++ {
			temp.Write([]byte((*input)[y+i][x+i]))
			coords = append(coords, []int{x + i, y + i})
			if temp.Len() == 4 {
				break
			}
		}
		return temp.String(), coords
	}
	return "", coords
}

func getBottomLeft(input *[][]string, x, y int) (string, [][]int) {
	coords := [][]int{}
	if x > 2 && len((*input))-y > 3 {
		temp := strings.Builder{}
		for i := 0; i < 4; i++ {
			temp.Write([]byte((*input)[y+i][x-i]))
			coords = append(coords, []int{x - i, y + i})
			if temp.Len() == 4 {
				break
			}
		}
		return temp.String(), coords
	}
	return "", coords
}

func findXmasCount(input *[][]string) (int, map[string][][]int) {
	track := 0
	o := map[string][][]int{}
	for y, row := range *input {
		for x := range row {
			// Check left
			toLeft, coords := getLeft(input, x, y)
			if toLeft == "XMAS" {
				track++
				// for debugging
				o[fmt.Sprintf("%d,%d", x, y)] = coords
			}
			// check top
			toTop, coords := getTop(input, x, y)
			if toTop == "XMAS" {
				track++
				// for debugging
				o[fmt.Sprintf("%d,%d", x, y)] = coords
			}
			// check right
			toRight, coords := getRight(input, x, y)
			if toRight == "XMAS" {
				track++
				// for debugging
				o[fmt.Sprintf("%d,%d", x, y)] = coords
			}
			// check bottom
			toBottom, coords := getBottom(input, x, y)
			if toBottom == "XMAS" {
				track++
				// for debugging
				o[fmt.Sprintf("%d,%d", x, y)] = coords
			}
			// check top-left
			toTopLeft, coords := getTopLeft(input, x, y)
			if toTopLeft == "XMAS" {
				track++
				// for debugging
				o[fmt.Sprintf("%d,%d", x, y)] = coords
			}
			// check top-right
			toTopRight, coords := getTopRight(input, x, y)
			if toTopRight == "XMAS" {
				track++
				// for debugging
				o[fmt.Sprintf("%d,%d", x, y)] = coords
			}
			// check bottom-right
			toBottomRight, coords := getBottomRight(input, x, y)
			if toBottomRight == "XMAS" {
				track++
				// for debugging
				o[fmt.Sprintf("%d,%d", x, y)] = coords
			}
			// check bottom-left
			toBottomLeft, coords := getBottomLeft(input, x, y)
			if toBottomLeft == "XMAS" {
				track++
				// for debugging
				o[fmt.Sprintf("%d,%d", x, y)] = coords
			}
		}
	}
	return track, o
}

func day4_part1() {
	/*
		Approach:
		Treat the input as a matrix n * m
		and check all 8 compass neighbours for the word
		XMAS and SAMX
		track all of the ones found to avoid duplicates
	*/
	input := ReadInput("./day4/day4.input.txt")
	track, _ := findXmasCount(&input)
	fmt.Printf("track: %v\n", track)
}

func day4_part2() {

}

func generateRandom(n int) {
	ls := []string{"X", "M", "A", "S"}
	max := big.NewInt(4)
	lines := []string{}
	txt := strings.Builder{}
	for {
		x, _ := rand.Int(rand.Reader, max)
		txt.Write([]byte(ls[x.Int64()]))
		if txt.Len() == n {
			lines = append(lines, txt.String())
			txt.Reset()
		}
		if len(lines) == n {
			break
		}
	}
	file, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}

func printWithColors(input *[][]string, o map[string][][]int) {
	allCoordsToPain := map[string]any{}
	// FLatten all coords
	for y, line := range *input {
		for x := range line {
			if coords, exists := o[fmt.Sprintf("%d,%d", x, y)]; exists {
				allCoordsToPain[fmt.Sprintf("%d,%d", x, y)] = 0
				for _, c := range coords {
					allCoordsToPain[fmt.Sprintf("%d,%d", c[0], c[1])] = 0
				}
			}
		}
	}

	// Draw
	colorcode := 31
	for y, line := range *input {
		for x, char := range line {
			if _, exists := allCoordsToPain[fmt.Sprintf("%d,%d", x, y)]; exists {
				fmt.Printf("\033[%dm%s\033[0m", colorcode, char)
				colorcode++
				if colorcode > 33 {
					colorcode = 31
				}
			} else {
				fmt.Print(char)
			}
		}
		fmt.Println()
	}
}

func main() {
	day4_part1()
	day4_part2()

	// Colourful touch
	// for range 2 {
	// 	generateRandom(100)
	// 	input := ReadInput("new.txt")
	// 	track, o := findXmasCount(&input)
	// 	printWithColors(&input, o)
	// 	fmt.Println(track)
	// 	fmt.Println()
	// 	fmt.Println()
	// }
}
