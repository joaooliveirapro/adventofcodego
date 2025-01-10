package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func getNums(line string) int {
	n1 := -1
	n2 := -1
	for x := range line {
		c := string(line[x])
		if strings.Contains("0123456789", c) {
			n, _ := strconv.Atoi(c)
			n1 = n //
			break
		}
	}

	for x := len(line) - 1; x >= 0; x-- {
		c := string(line[x])
		if strings.Contains("0123456789", c) {
			n, _ := strconv.Atoi(c)
			n2 = n //
			break
		}
	}

	n3, _ := strconv.Atoi(fmt.Sprintf("%v%v", n1, n2))
	return n3
}

func part1(input *[]string) {
	sum := 0
	for _, line := range *input {
		sum += getNums(line)
	}
	fmt.Printf("sum: %v\n", sum)
}

func part2(input *[]string) {

}

func main() {
	input := utils.ReadInput("./2023/day1/day1.input.txt")
	part1(&input)
}
