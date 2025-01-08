package main

import (
	"fmt"
	"strings"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func part1(input *[]string) {
	nice := 0
	m := map[string]any{}
	for _, line := range *input {

		// Vowel rule
		vowelCount := 0
		for _, _v := range "aeiou" {
			v := string(_v)

			vowelCount += strings.Count(line, v)

		}

		// Repeated letters rule
		rule2 := false
		for x := range line {
			if x < len(line)-1 && line[x] == line[x+1] {
				rule2 = true
			}
		}

		// ab, cd, pq, or xy rule
		rule3 := true
		if strings.Contains(line, "ab") {
			rule3 = false
		}
		if strings.Contains(line, "cd") {
			rule3 = false
		}
		if strings.Contains(line, "pq") {
			rule3 = false
		}
		if strings.Contains(line, "xy") {
			rule3 = false
		}

		if vowelCount > 2 && rule2 && rule3 {
			nice++
			m[line] = "nice"
		} else {
			m[line] = fmt.Sprintf("%v %v %v", vowelCount, rule2, rule3)
		}

	}

	fmt.Printf("nice: %v\n", nice)

}

func part2(input *[]string) {
	nice := 0
	for _, line := range *input {
		rule1 := false
		// pair of any two letters
		for _, a := range "abcdefghijklmnopqrstuvwxyz" {
			for _, b := range "abcdefghijklmnopqrstuvwxyz" {
				pair := fmt.Sprintf("%s%s", string(a), string(b))
				if strings.Count(line, pair) > 1 {
					rule1 = true
				}
			}
		}

		rule2 := false
		// contains at least one letter which repeats with exactly one letter between them
		for x := range line {
			if x >= len(line)-2 {
				break
			}
			if line[x] == line[x+2] {
				rule2 = true
			}
		}

		if rule1 && rule2 {
			nice++
		}

	}
	fmt.Printf("nice: %v\n", nice)

}

func main() {
	input := utils.ReadInput("./2015/day5/day5.input.txt")
	part1(&input)
	part2(&input)
}
