package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func AreAllIncreasing(nums []int) bool {
	n := nums[0] // Starts with first number
	i := 1       // Index
	for {
		if i == len(nums) {
			break
		}
		if nums[i] <= n {
			return false
		}
		n = nums[i]
		i++
	}
	return true
}

func AreAllDecreasing(nums []int) bool {
	n := nums[0] // Starts with first number
	i := 1       // Index
	for {
		if i == len(nums) {
			break
		}
		if nums[i] >= n {
			return false
		}
		n = nums[i]
		i++
	}
	return true
}

func DiffersByAtLeastOneOrAtMostThree(nums []int) bool {
	index := 0
	for {
		if index == len(nums)-1 { // -1 to avoid OutOfBounds error
			break
		}
		// Calculate absolute difference between element[index] and the following element
		diff := math.Abs(float64(nums[index]) - float64(nums[index+1]))

		// Exit conditions
		if diff < 1 || diff > 3 {
			return false
		}
		// Increment index to check following elements
		index++
	}
	return true
}

func IsSafeReport(nums []int) bool {
	return (AreAllDecreasing(nums) || AreAllIncreasing(nums)) && DiffersByAtLeastOneOrAtMostThree(nums)
}

func ReadInput(filename string) [][]int {
	txt, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	allLines := [][]int{}
	for _, i := range strings.Split(string(txt), "\n") {
		lineNums := strings.Fields(i)
		lineNumsInt := []int{}
		for _, n := range lineNums {
			nu, _ := strconv.Atoi(n)
			lineNumsInt = append(lineNumsInt, nu)
		}
		allLines = append(allLines, lineNumsInt)
	}
	return allLines
}

func day2_part1() {
	safeReports := 0
	input := ReadInput("./day2/day2.input.txt")
	for _, nums := range input {
		if IsSafeReport(nums) {
			safeReports++
			// fmt.Printf("nums: %v\n", nums)
		}
	}
	fmt.Printf("There are %d safe reports.\n", safeReports)
}

func day2_part2() {
	safeReports := 0
	input := ReadInput("./day2/day2.input.txt")
	for _, nums := range input {
		if IsSafeReport(nums) {
			safeReports++
			// fmt.Printf("nums: %v\n", nums)
		} else {
			// Check by removing each level
			for i := 0; i < len(nums); i++ {
				// i is the index to exclude
				numsWithoutLevel := []int{}
				for x := 0; x < len(nums); x++ {
					if x == i {
						continue
					}
					numsWithoutLevel = append(numsWithoutLevel, nums[x])
				}
				if IsSafeReport(numsWithoutLevel) {
					safeReports++
					break
				}
			}
		}
	}
	fmt.Printf("There are %d safe reports.\n", safeReports)

}

func main() {
	day2_part1()
	day2_part2()
}
