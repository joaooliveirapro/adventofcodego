package main

import (
	"fmt"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func main() {

	input := utils.ReadInput("./2015/day3/day3.input.txt")

	// Starts always with one present delivered at first house (0,0)

	seen := map[string]int{"0,0": 1}

	xPos, yPos := 0, 0
	for _, line := range input {
		for _, _ch := range line {
			ch := string(_ch)
			if ch == "<" {
				xPos--
			} else if ch == "^" {
				yPos++
			} else if ch == ">" {
				xPos++
			} else if ch == "v" {
				yPos--
			}
			seen[fmt.Sprintf("%d,%d", xPos, yPos)] += 1
		}
	}

	fmt.Printf("seen: %v\n", len(seen))
}
