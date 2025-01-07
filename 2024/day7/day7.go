package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func ParseInput(input *[]string) [][]int {
	// Parse input into [][]int
	allNums := [][]int{}
	for _, line := range *input {
		line_ := strings.Replace(line, ":", "", -1)
		numsStr := strings.Fields(line_)
		nums := []int{}
		for _, n := range numsStr {
			nInt, _ := strconv.Atoi(n)
			nums = append(nums, nInt)
		}
		allNums = append(allNums, nums)
	}
	return allNums
}

func generateExpressions(numbers []int) []string {
	if len(numbers) < 2 {
		return nil // At least two numbers are needed to insert operators
	}
	var result []string
	var helper func(expr string, index int)
	helper = func(expr string, index int) {
		if index == len(numbers)-1 {
			result = append(result, expr)
			return
		}
		helper(expr+" * "+fmt.Sprint(numbers[index+1]), index+1)
		helper(expr+" + "+fmt.Sprint(numbers[index+1]), index+1)
		helper(expr+" || "+fmt.Sprint(numbers[index+1]), index+1)
	}
	// Start with the first number
	helper(fmt.Sprint(numbers[0]), 0)
	return result
}

// evaluateExpression evaluates an expression with an arbitrary number of "*" and "+" operators.
func evaluateExpression(expr string) (int, error) {
	tokens := strings.Fields(expr)
	total := 0
	op := ""
	for i, t := range tokens {
		// If it's a number (odd positions)
		if i%2 == 0 {
			// check previous index for the operand
			// unless it's the very first token. In that case op is +
			if i == 0 {
				op = "+"
			} else {
				op = tokens[i-1]
			}

			n, _ := strconv.Atoi(t)
			if op == "+" {
				total += n
			} else if op == "*" {
				total *= n
			} else { // Op = || (Concatenate prev + next num)
				nextD, _ := strconv.Atoi(tokens[i-2])
				if i > 3 {
					nextD = total
				}
				newConcDig, _ := strconv.Atoi(fmt.Sprintf("%d%d", nextD, n))
				total += newConcDig - nextD
			}

		}
	}
	return total, nil
}

func day7_part1() {
	input := utils.ReadInput("./day7/day7.input.txt")
	nums := ParseInput(&input)
	sols := map[string]int{}
	for _, line := range nums {
		total := line[0] // First number is the total
		ns := line[1:]
		for _, p := range generateExpressions(ns) {
			t, _ := evaluateExpression(p)
			if t == total {
				sols[p] = total
				break
			}
		}
	}

	total := 0
	file, err := os.Create("sols.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write content to the file

	for k, v := range sols {
		total += v
		_, err = file.WriteString(fmt.Sprintf("%v=%v\n", v, k))
		if err != nil {
			panic(err)
		}

	}
	fmt.Println("total=", total)
}

func day7_part2() {
	// included in day7_part1()
}

func main() {
	day7_part1()
	day7_part2()
}
