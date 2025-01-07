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

func findValidUpdates(updates *[]string, ordRules *[]string) []string {
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

func fixUpdate(update []int, rules [][2]int) string {
	// Build graph and in-degree map
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	for _, num := range update {
		inDegree[num] = 0
	}

	for _, rule := range rules {
		from, to := rule[0], rule[1]
		graph[from] = append(graph[from], to)
		inDegree[to]++
	}

	// Topological sort using Kahn's algorithm
	var sortedOrder []int
	queue := []int{}

	// Enqueue nodes with in-degree 0
	for num, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, num)
		}
	}

	// Process the queue
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sortedOrder = append(sortedOrder, current)

		// Decrease in-degree of neighbors
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	sortedOrderStr := strings.Builder{}
	for p, i := range sortedOrder {
		comma := ","
		if p == len(sortedOrder)-1 {
			comma = ""
		}
		sortedOrderStr.Write([]byte(fmt.Sprintf("%d%s", i, comma)))
	}
	return sortedOrderStr.String()
}

func day5_part1() {
	input := ReadInput("./day5/day5.input.txt")
	ordRules := parseOrderingRules(&input)
	updates := parseUpdates(&input)
	validUpdates := findValidUpdates(&updates, &ordRules)
	middle := calculateMiddleSum(validUpdates)
	fmt.Printf("Middle=%d", middle)
}

func day5_part2() {
	input := ReadInput("./day5/day5.input.txt")
	ordRules := parseOrderingRules(&input)
	updates := parseUpdates(&input)
	validUpdates := findValidUpdates(&updates, &ordRules)
	invalidUpdates := []string{}
	// Find which updates are invalid
	for _, u := range updates {
		isValid := false
		for _, vu := range validUpdates {
			if u == vu {
				isValid = true
				break
			}
		}
		if !isValid {
			invalidUpdates = append(invalidUpdates, u)
		}
	}
	fixedUpdates := []string{}
	// Find what rules they break
	for _, iu := range invalidUpdates {
		rulesToConsider := getRulesToConsider(&iu, &ordRules)

		update := []int{}
		for _, char := range strings.Split(iu, ",") {
			n, _ := strconv.Atoi(char)
			update = append(update, n)
		}

		rules := [][2]int{}
		for _, r := range rulesToConsider {
			n1, _ := strconv.Atoi(r.num1)
			n2, _ := strconv.Atoi(r.num2)
			rules = append(rules, [2]int{n1, n2})
		}

		fixedup := fixUpdate(update, rules)
		fixedUpdates = append(fixedUpdates, fixedup)
	}

	fmt.Println(calculateMiddleSum(fixedUpdates))
}

func main() {
	day5_part1()
	day5_part2()
}
