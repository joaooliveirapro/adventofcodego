package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input file
	data, _ := os.ReadFile("./day9/day9.input.txt")

	// Parse the input into an array of integers
	lengths := make([]int, len(data))
	for i, ch := range data {
		num, _ := strconv.Atoi(string(ch))
		lengths[i] = num
	}

	// Create the grid
	grid := []int{}
	for i, num := range lengths {
		for j := 0; j < num; j++ {
			if i%2 == 0 {
				grid = append(grid, i/2)
			} else {
				grid = append(grid, -1)
			}
		}
	}

	// Compact the grid
	for {
		if len(grid) > 0 && grid[len(grid)-1] == -1 {
			grid = grid[:len(grid)-1]
		} else {
			index := -1
			for idx, val := range grid {
				if val == -1 {
					index = idx
					break
				}
			}
			if index == -1 {
				break
			}
			grid[index] = grid[len(grid)-1]
			grid = grid[:len(grid)-1]
		}
	}

	// Calculate the answer
	answer := 0
	for i, num := range grid {
		answer += i * num
	}

	fmt.Println(answer)
}
