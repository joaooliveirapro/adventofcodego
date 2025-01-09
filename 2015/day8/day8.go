package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func part1(input *[]string) {
	codelen := 0
	memlen := 0
	for _, line := range *input {
		codelen += len(line)
		line1 := regexp.MustCompile(`\\x[a-fA-F0-9]{2}`).ReplaceAllString(line, "#")
		line2 := strings.ReplaceAll(line1, "\\\\", "#")
		line3 := strings.ReplaceAll(line2, "\\\"", "#")
		memlen += len(line3) - 2
		fmt.Printf("len(line): %v %v len(line3) %v\n", len(line), line, len(line3)-2)
	}
	fmt.Printf("Diff: %v | (%v - %v)\n", codelen-memlen-2, codelen, memlen)
}

func part2(input *[]string) {

}

func main() {
	input := utils.ReadInput("./2015/day8/day8.input.txt")
	part1(&input)
}
