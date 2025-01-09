package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
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

	// save as image
	img := gridToImage(grid)
	// Save the image as a PNG file
	outputFile, err := os.Create("./2015/day6/output.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, img)
	if err != nil {
		panic(err)
	}

}

// Converts a 2D grid of 1s and 0s to an image
func gridToImage(grid [][]int) *image.Gray {
	height := len(grid)
	width := len(grid[0])

	// Create a new grayscale image
	img := image.NewGray(image.Rect(0, 0, width, height))

	// Set pixel colors based on the grid
	for y, row := range grid {
		for x, value := range row {
			if value == 1 {
				img.SetGray(x, y, color.Gray{Y: 0}) // Black for 1
			} else {
				img.SetGray(x, y, color.Gray{Y: 255}) // White for 0
			}
		}
	}

	return img
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

func turnon2(line string, grid *[][]int) {
	coords := getCoords(line)
	fromX, fromY := coords[0][0], coords[0][1]
	toX, toY := coords[1][0], coords[1][1]

	for y, row := range *grid {
		for x := range row {
			if x >= fromX && y >= fromY && x <= toX && y <= toY {
				(*grid)[y][x] += 1
			}
		}
	}
}

func turnoff2(line string, grid *[][]int) {
	coords := getCoords(line)
	fromX, fromY := coords[0][0], coords[0][1]
	toX, toY := coords[1][0], coords[1][1]

	for y, row := range *grid {
		for x := range row {
			if x >= fromX && y >= fromY && x <= toX && y <= toY {
				if (*grid)[y][x] > 0 {
					(*grid)[y][x] -= 1
				}
			}
		}
	}
}

func toggle2(line string, grid *[][]int) {
	coords := getCoords(line)
	fromX, fromY := coords[0][0], coords[0][1]
	toX, toY := coords[1][0], coords[1][1]

	for y, row := range *grid {
		for x := range row {
			if x >= fromX && y >= fromY && x <= toX && y <= toY {
				(*grid)[y][x] += 2
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

func gridToImageGrayscale(grid [][]int) *image.Gray {
	height := len(grid)
	width := len(grid[0])

	// Create a new grayscale image
	img := image.NewGray(image.Rect(0, 0, width, height))

	// Set pixel colors based on the grid
	for y, row := range grid {
		for x, value := range row {
			if value < 1 {
				img.SetGray(x, y, color.Gray{Y: 0}) // Black for 1
			} else {
				y_ := (float32(grid[y][x]) / float32(54)) * 255
				img.SetGray(x, y, color.Gray{Y: uint8(y_)}) // White for 0
			}
		}
	}

	return img
}
func part2(input *[]string) {
	// Make grid of slices
	grid := make([][]int, 1000) // outter slices
	for i := range grid {
		grid[i] = make([]int, 1000) // inner slices
	}

	// Read input lines
	for _, line := range *input {
		// turn on (1)
		if strings.Contains(line, "turn on") {
			turnon2(line, &grid)
		}
		// turn off (0)
		if strings.Contains(line, "turn off") {
			turnoff2(line, &grid)
		}
		// toggle
		if strings.Contains(line, "toggle") {
			toggle2(line, &grid)
		}
	}

	brightness := 0
	max := 0
	min := 0
	for y, line := range grid {
		for x := range line {
			if grid[y][x] >= max {
				max = grid[y][x]
			} else {
				min = grid[y][x]
			}
			brightness += grid[y][x]
		}
	}
	fmt.Printf("brightness: %v\n", brightness)
	fmt.Printf("max brightness: %v\n", max)
	fmt.Printf("min: %v\n", min)
	img := gridToImageGrayscale(grid)
	// Save the image as a PNG file
	outputFile, err := os.Create("./2015/day6/output2.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	err = png.Encode(outputFile, img)
	if err != nil {
		panic(err)
	}
}

func main() {
	input := utils.ReadInput("./2015/day6/day6.input.txt")
	part1(&input)
	part2(&input)
}
