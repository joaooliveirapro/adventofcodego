package main

import (
	"fmt"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func part1(input *[]string) {
	// Starts always with one present delivered at first house (0,0)

	seen := map[string]int{"0,0": 1}

	xPos, yPos := 0, 0
	for _, line := range *input {
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

func part2(input *[]string) {

	santa := map[string]int{"0,0": 1}
	robot := map[string]int{"0,0": 1}

	santaXpos, santaYpos := 0, 0
	robotXpos, robotYpos := 0, 0

	for _, line := range *input {
		for x, _ch := range line {
			ch := string(_ch)

			who := "santa"
			if x%2 == 0 {
				// santa goes first so he's set by default
			} else {
				// if odd, it's robot's turn
				who = "robot"
			}

			if ch == "<" {
				if who == "santa" {
					santaXpos--
				} else {
					robotXpos--
				}
			} else if ch == "^" {
				if who == "santa" {
					santaYpos++
				} else {
					robotYpos++
				}
			} else if ch == ">" {
				if who == "santa" {
					santaXpos++
				} else {
					robotXpos++
				}
			} else if ch == "v" {
				if who == "santa" {
					santaYpos--
				} else {
					robotYpos--
				}
			}

			if who == "santa" {
				santa[fmt.Sprintf("%d,%d", santaXpos, santaYpos)] += 1
			} else {
				robot[fmt.Sprintf("%d,%d", robotXpos, robotYpos)] += 1
			}
		}
	}

	all := map[string]int{}
	for a, b := range santa {
		all[a] = b
	}
	for a, b := range robot {
		all[a] += b
	}
	fmt.Printf("all: %v\n", len(all))

}

func main() {

	input := utils.ReadInput("./2015/day3/day3.input.txt")
	part1(&input)
	part2(&input)
}
