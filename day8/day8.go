package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func getReflectedCoords(a []int, b []int) [2]int {
	ax, ay := a[0], a[1]
	bx, by := b[0], b[1]
	rx := (bx - ax) + bx
	ry := (by - ay) + by
	if rx < 0 || rx > 49 || ry < 0 || ry > 49 {
		return [2]int{-1, -1}
	}
	return [2]int{rx, ry}
}

func getSimilarAntennaCoords(input *[]string, ch string, cx, cy int) [][]int {
	coords := [][]int{}
	for y, line := range *input {
		for x, char := range line {
			if x == cx && y == cy {
				continue // skip the ch itself
			}
			if string(char) == ch {
				coords = append(coords, []int{x, y})
			}
		}
	}
	return coords
}

func day8_part1() {
	input := utils.ReadInput("./day8/day8.input.txt")

	newTextGrid := [][]string{}
	for _, line := range input {
		temp := []string{}
		for _, char := range line {
			temp = append(temp, string(char))
		}
		newTextGrid = append(newTextGrid, temp)
	}

	for y, line := range input {
		for x, char := range line {
			ch := string(char)
			// Skip "." as these don't matter
			if ch != "." {
				// Get all similar antenna's coords
				similarAntenas := getSimilarAntennaCoords(&input, ch, x, y)

				fmt.Println(similarAntenas)
				// Find similar antenas coord
				for _, antennaCoord := range similarAntenas {
					ax, ay := antennaCoord[0], antennaCoord[1]

					// Find the reflected coords
					r := getReflectedCoords([]int{x, y}, []int{ax, ay})

					rx, ry := r[0], r[1]

					if rx == -1 || ry == -1 {
						// skip out of bounds
						continue
					}
					newTextGrid[ry][rx] = "#"
				}
			}
		}
	}

	var newTxt strings.Builder
	for _, line := range newTextGrid {
		for _, char := range line {
			newTxt.WriteString(char)
		}
		newTxt.WriteString("\n")
	}

	file, err := os.Create("./day8/sols8.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(newTxt.String())
}

func main() {
	day8_part1()
}
