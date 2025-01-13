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

func popMax(a []int) (int, []int) {
	aa := []int{}
	max := 0
	xx := 0
	for x := range a {
		if a[x] > max {
			max = a[x]
			xx = x
		}
	}
	for x := range a {
		if x == xx {
			continue
		}
		aa = append(aa, a[x])
	}
	return max, aa
}

func sum(a []int) int {
	t := 0
	for x := range a {
		t += a[x]
	}
	return t
}

func part2(input *[]string) {
	allCals := []int{}

	elfCounter := 0
	elfCals := 0
	for _, line := range *input {
		cals, _ := strconv.Atoi(line)
		elfCals += cals
		if line == "" {
			allCals = append(allCals, elfCals)
			elfCounter += 1
			elfCals = 0
		}
	}
	top3 := []int{}
	for range 3 {
		n, newAll := popMax(allCals)
		allCals = newAll
		top3 = append(top3, n)
	}

	fmt.Printf("top3: %v\n", top3)
	fmt.Printf("sum(top3): %v\n", sum(top3))
}

func main() {
	input := utils.ReadInputNoSkipEmpty("./2022/day1/day1.input.txt")
	part1(&input)
	part2(&input)
}
