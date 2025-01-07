package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func day3_part1() {
	// Read input data
	txt, err := os.ReadFile("./day3/day3.input.txt")
	if err != nil {
		panic(err.Error())
	}

	// Find all mul(X,X)
	// Multiply and add to total
	total := 0
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for _, matches := range r.FindAllStringSubmatch(string(txt), -1) {
		d1, _ := strconv.Atoi(matches[1])
		d2, _ := strconv.Atoi(matches[2])
		total += d1 * d2
	}

	fmt.Printf("Total=%v\n", total)
}

func day3_part2() {
	// Ready input data
	txt, err := os.ReadFile("./day3/day3.input.txt")
	if err != nil {
		panic(err.Error())
	}

	// My approach is to re-write the input and remove everything that is
	// between a don't() and a do() since this is not meant to be executed.
	r := regexp.MustCompile(`(?s)don't\(\).*?do\(\)`)
	newTxt := r.ReplaceAll(txt, []byte(""))

	// Multipy and add
	total := 0
	r2 := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for _, matches := range r2.FindAllStringSubmatch(string(newTxt), -1) {
		d1, _ := strconv.Atoi(matches[1])
		d2, _ := strconv.Atoi(matches[2])
		total += d1 * d2
	}
	fmt.Printf("Total=%v\n", total)
}

func main() {
	day3_part1()
	day3_part2()
}
