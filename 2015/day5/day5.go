package main

import (
	"fmt"
	"strings"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func main() {

	input := utils.ReadInput("./2015/day5/day5.input.txt")

	nice := 0
	m := map[string]any{}
	for _, line := range input {

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
