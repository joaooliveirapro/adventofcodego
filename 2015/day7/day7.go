package main

import (
	"fmt"
	"strings"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func part1(input *[]string) {
	m := map[string]int{}
	for _, line := range *input {
		// op can be OR, NOT, AND, RSHIFT, LSHIFT and MOVE
		op := parseInstruction(line)
		// Control flow
		if op == "OR" {
			applyOR(line, &m)
		} else if op == "NOT" {
			applyNOT(line, &m)
		} else if op == "AND" {
			applyAND(line, &m)
		} else if op == "RSHIFT" {
			applyRSHIFT(line, &m)
		} else if op == "LSHIFT" {
			applyLSHIFT(line, &m)
		} else { // MOVE
			applyMOVE(line, &m)
		}
	}
	fmt.Printf("m[\"a\"]: %v\n", m["a"])
}

func applyOR(line string, m *map[string]int) {
	// kg OR kf -> kh
	p := strings.Fields(line)
	left, right, out := p[0], p[2], p[4]
	// ensure it's defined first
	if leftV, ok := (*m)[left]; ok {
		if rightV, ok := (*m)[right]; ok {
			// Because maps are initialized to 0,
			// no need to do existence checking
			(*m)[out] += leftV | rightV
		}
	}
}

func applyNOT(line string, m *map[string]int) {}

func applyAND(line string, m *map[string]int) {}

func applyRSHIFT(line string, m *map[string]int) {}

func applyLSHIFT(line string, m *map[string]int) {}

func applyMOVE(line string, m *map[string]int) {}

func parseInstruction(line string) string {
	if strings.Contains(line, "OR") {
		return "OR"
	} else if strings.Contains(line, "NOT") {
		return "NOT"
	} else if strings.Contains(line, "AND") {
		return "AND"
	} else if strings.Contains(line, "RSHIFT") {
		return "RSHIFT"
	} else if strings.Contains(line, "LSHIFT") {
		return "LSHIFT"
	} else {
		return "MOVE"
	}
}

func part2(input *[]string) {

}

func main() {
	input := utils.ReadInput("./2015/day7/day7.input.txt")
	part1(&input)
}
