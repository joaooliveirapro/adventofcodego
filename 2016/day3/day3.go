package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func part1(input *[]string) {
	valid := 0
	for _, line := range *input {
		n := strings.Fields(line)
		a, _ := strconv.Atoi(n[0])
		b, _ := strconv.Atoi(n[1])
		c, _ := strconv.Atoi(n[2])

		if a+b > c && a+c > b && b+c > a {
			valid++
		}
	}
	fmt.Printf("valid: %v\n", valid)
}

func part2(input *[]string) {

}

func main() {
	input := utils.ReadInput("./2016/day3/day3.input.txt")
	part1(&input)
	part2(&input)
}
