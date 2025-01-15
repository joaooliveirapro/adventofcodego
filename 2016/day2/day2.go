package main

import (
	"fmt"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func part1(input *[]string) {
	nums := make([]int, len(*input))
	ma := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	/*
		1 2 3
		4 5 6
		7 8 9
	*/
	x, y := 0, 0 // 5
	for i, line := range *input {
		for _, c_ := range line {
			c := string(c_)
			if c == "L" {
				x -= 1
				if x < 0 {
					x = 0
				}
			} else if c == "R" {
				x += 1
				if x > 2 {
					x = 2
				}
			} else if c == "U" {
				y -= 1
				if y < 0 {
					y = 0
				}
			} else if c == "D" {
				y += 1
				if y > 2 {
					y = 2
				}
			}
		}
		d := ma[y][x]
		nums[i] = d
	}
	fmt.Printf("nums: %v\n", nums)
}

func part2(input *[]string) {
	nums := make([]string, len(*input))
	ma := [][]string{
		{".", ".", "1", ".", "."}, // y = 0
		{".", "2", "3", "4", "."}, // y = 1
		{"5", "6", "7", "8", "9"}, // y = 2
		{".", "A", "B", "C", "."}, // y = 3
		{".", ".", "D", ".", "."}, // y = 4
	} // x=0, x=1, x=2, x=3, x=4
	/*
		. . 1 . .
		. 2 3 4 .
		5 6 7 8 9
		. A B C .
		. . D . .
	*/

	forbidden := map[string]bool{}
	f := [][]int{{0, 0}, {1, 0}, {3, 0}, {4, 0}, {0, 1}, {4, 1}, {0, 3}, {4, 3}, {0, 4}, {1, 4}, {3, 4}, {4, 4}}
	for _, a := range f {
		xx, yy := a[0], a[1]
		forbidden[fmt.Sprintf("%d,%d", xx, yy)] = true
	}

	x, y := 0, 2 // 5
	for i, line := range *input {
		for _, c_ := range line {
			oldX, oldY := x, y
			c := string(c_)
			if c == "L" {
				x -= 1
				if x < 0 {
					x = 0
				}
			} else if c == "R" {
				x += 1
				if x > 4 {
					x = 4
				}
			} else if c == "U" {
				y -= 1
				if y < 0 {
					y = 0
				}
			} else if c == "D" {
				y += 1
				if y > 4 {
					y = 4
				}
			}
			_, invalid := forbidden[fmt.Sprintf("%d,%d", x, y)]
			if invalid {
				x, y = oldX, oldY
			}

		}
		d := ma[y][x]
		nums[i] = d
	}
	fmt.Printf("nums: %v\n", nums)
}

func main() {
	input := utils.ReadInput("./2016/day2/day2.input.txt")
	part1(&input)
	part2(&input)
}
