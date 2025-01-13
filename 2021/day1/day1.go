package main

import (
	"fmt"
	"strconv"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func part1(input *[]string) {

	prev := -1
	m := map[string]int{}
	for _, numStr := range *input {
		num, _ := strconv.Atoi(numStr)

		// Skip first iteration after assignement
		if prev == -1 {
			prev = num
			continue
		}

		// compare with prev
		if num-prev > 0 {
			m["up"]++
		} else {
			m["down"]++
		}

		prev = num

	}
	fmt.Printf("m: %v\n", m)
}

func part2(input *[]string) {

}

func main() {
	input := utils.ReadInput("./2021/day1/day1.input.txt")
	part1(&input)
	part2(&input)
}
