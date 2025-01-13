package main

import (
	"fmt"
	"strconv"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func part1(input *[]string) {
	mostCalories := 0
	elfCounter := 0
	elfCals := 0
	for _, line := range *input {
		cals, _ := strconv.Atoi(line)
		elfCals += cals
		if line == "" {
			if elfCals > mostCalories {
				mostCalories = elfCals
			}
			elfCounter += 1
			elfCals = 0
		}
	}
	fmt.Printf("mostCalories: %v\n", mostCalories)

}

func part2(input *[]string) {

}

func main() {
	input := utils.ReadInputNoSkipEmpty("./2022/day1/day1.input.txt")
	part1(&input)
}
