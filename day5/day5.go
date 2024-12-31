package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadInput(filename string) []string {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Remove leading/trailing spaces
		if len(line) != 0 {                       // Skip empty lines
			lines = append(lines, line)
		}
	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	return lines
}

func parseOrderingRules(input *[]string) []string {
	ordRules := []string{}
	for _, line := range *input {
		if len(line) == 5 {
			ordRules = append(ordRules, line)
		}
	}
	return ordRules
}

func parseUpdates(input *[]string) []string {
	updates := []string{}
	for _, line := range *input {
		if len(line) != 5 {
			updates = append(updates, line)
		}
	}
	return updates
}

type Rule struct {
	num1 string
	num2 string
}

func getRulesToConsider(update *string, ordRules *[]string) []Rule {
	// REGEX approach
	// Check if both numbers on the ordRule are present in update
	rulesToConsider := []Rule{}
	for _, rule := range *ordRules {
		rg := regexp.MustCompile(`(\d+)\|(\d+)`)
		for _, match := range rg.FindAllStringSubmatch(rule, -1) {
			num1 := match[1]
			num2 := match[2]
			if strings.Contains(*update, num1) && strings.Contains(*update, num2) {
				rulesToConsider = append(rulesToConsider, Rule{num1: num1, num2: num2})
			}
		}
	}
	return rulesToConsider
}

func findValudUpdates(updates *[]string, ordRules *[]string) []string {
	validUpdates := []string{}
	for _, u := range *updates {
		isValid := true
		// find what rules are to be considered for this updated
		rules := getRulesToConsider(&u, ordRules)
		// check if ALL those rules are followed. If not, then update is NOT valid
		for _, rule := range rules {
			// check if the index of num1 is GREATER than index of num2 (comes AFTER than num2)
			// this invalidates the update
			if strings.Index(u, rule.num1) >= strings.Index(u, rule.num2) {
				isValid = false
				break // skip
			}
		}
		// If we get to this stage with isValid true, then update is valid
		if isValid {
			validUpdates = append(validUpdates, u)
		}
	}
	return validUpdates
}

func calculateMiddleSum(updates []string) int {
	middle := 0
	for _, vu := range updates {
		nums := strings.Split(vu, ",")
		temp := []int{}
		for _, nStr := range nums {
			n, _ := strconv.Atoi(nStr)
			temp = append(temp, n)
		}
		middleIndex := len(temp) / 2
		middle += temp[middleIndex]
	}
	return middle
}

func day5_part1() {
	input := ReadInput("./day5/day5.input.txt")
	ordRules := parseOrderingRules(&input)
	updates := parseUpdates(&input)
	validUpdates := findValudUpdates(&updates, &ordRules)

	middle := calculateMiddleSum(validUpdates)
	fmt.Printf("Middle=%d", middle)
}

func day5_part2() {

}

func main() {
	day5_part1()
	day5_part2()
}
