package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func part1(input *[]string) {
	// Make grid of slices
	grid := make([][]int, 1000) // outter slices
	for i := range grid {
		grid[i] = make([]int, 1000) // inner slices
	}

	// Read input lines
	for _, line := range *input {
		// turn on (1)
		if strings.Contains(line, "turn on") {
			turnon(line, &grid)
		}
		// turn off (0)
		if strings.Contains(line, "turn off") {
			turnoff(line, &grid)
		}
		// toggle
		if strings.Contains(line, "toggle") {
			toggle(line, &grid)
		}
	}
	// lit lights
	lit := 0
	for y, line := range grid {
		for x := range line {
			lit += grid[y][x]
			// fmt.Print(grid[y][x])
		}
		// fmt.Println()
	}
	fmt.Printf("lit: %v\n", lit)
}

func turnon(line string, grid *[][]int) {
	coords := getCoords(line)
	fromX, fromY := coords[0][0], coords[0][1]
	toX, toY := coords[1][0], coords[1][1]

	for y, row := range *grid {
		for x := range row {
			if x >= fromX && y >= fromY && x <= toX && y <= toY {
				(*grid)[y][x] = 1
			}
		}
	}
}

func turnoff(line string, grid *[][]int) {
	coords := getCoords(line)
	fromX, fromY := coords[0][0], coords[0][1]
	toX, toY := coords[1][0], coords[1][1]

	for y, row := range *grid {
		for x := range row {
			if x >= fromX && y >= fromY && x <= toX && y <= toY {
				(*grid)[y][x] = 0
			}
		}
	}
}

func toggle(line string, grid *[][]int) {
	coords := getCoords(line)
	fromX, fromY := coords[0][0], coords[0][1]
	toX, toY := coords[1][0], coords[1][1]

	for y, row := range *grid {
		for x := range row {
			if x >= fromX && y >= fromY && x <= toX && y <= toY {
				if (*grid)[y][x] == 1 {
					(*grid)[y][x] = 0
				} else {
					(*grid)[y][x] = 1
				}
			}
		}
	}
}

func getCoords(line string) [][]int {
	// returns [[x1 y1], [x2, y2]]
	a := [][]int{}
	r := regexp.MustCompile(`(\d+\,\d+)`)
	for _, m := range r.FindAllString(line, -1) {
		n := strings.Split(m, ",")
		n1, _ := strconv.Atoi(n[0])
		n2, _ := strconv.Atoi(n[1])
		a = append(a, []int{n1, n2})
	}
	return a
}

func part2(input *[]string) {

}

func main() {
	input := utils.ReadInput("./2015/day6/day6.input.txt")
	part1(&input)
}
