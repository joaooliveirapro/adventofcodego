package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func part1(input *[]string) {
	x, y := 0, 0
	facing := "n" // nesw
	for _, line := range *input {
		fs := strings.Fields(line)
		for _, ins := range fs {
			ins_ := strings.Replace(ins, ",", "", 1)
			dir, numS := ins_[:1], ins_[1:]
			num, _ := strconv.Atoi(numS)

			if facing == "n" {
				if dir == "L" {
					x -= num
					facing = "w"
				} else {
					x += num
					facing = "e"
				}
			} else if facing == "e" {
				if dir == "L" {
					y += num
					facing = "n"
				} else {
					y -= num
					facing = "s"
				}
			} else if facing == "s" {
				if dir == "L" {
					x += num
					facing = "e"
				} else {
					x -= num
					facing = "w"
				}
			} else if facing == "w" {
				if dir == "L" {
					y -= num
					facing = "s"
				} else {
					y += num
					facing = "n"
				}
			}
			// fmt.Printf("x,y: %v,%v\n", x, y)
		}
	}
	fmt.Printf("You're %v blocks away\n", x+y)
}

func part2(input *[]string) {

}

func main() {
	input := utils.ReadInput("./2016/day1/day1.input.txt")
	part1(&input)
	part2(&input)
}
