package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortz(a []int) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
}

func ReadInput(filename string) ([]int, []int) {
	txt, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	left := []int{}
	right := []int{}
	for _, i := range strings.Split(string(txt), "\n") {
		lineNums := strings.Fields(i)
		l, _ := strconv.Atoi(lineNums[0])
		r, _ := strconv.Atoi(lineNums[1])
		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func Set(nums []int) []int {
	temp := map[int]bool{}
	for _, i := range nums {
		temp[i] = true
	}
	unique := []int{}
	for k, _ := range temp {
		unique = append(unique, k)
	}
	return unique
}

func CountIn(n int, side []int) int {
	c := 0
	for _, k := range side {
		if k == n {
			c++
		}
	}
	return c
}

func day1_part2() {
	left, right := ReadInput("./day1/day1.input.txt")
	i := 0
	similarity := 0
	for {
		if i == len(left) {
			break
		}

		leftNum := left[i]
		rightCount := CountIn(leftNum, right)

		similarity += leftNum * rightCount
		i++
	}
	fmt.Println(similarity)
}

func day1_part1() {
	// Ready input
	left, right := ReadInput("./day1/day1.input.txt")

	// Sort each list first
	sortz(left)
	sortz(right)

	// loop through index o to len(left) - 1
	diff := 0
	i := 0
	for {
		// Exit once all diff are computed
		if i > len(left)-1 {
			break
		}
		// Add to diff the abs difference between
		// left[i] and right[i]
		diff += int(math.Abs(float64(left[i]) - float64(right[i])))
		i++
	}
	fmt.Println(diff)
}

func main() {
	day1_part1()
	day1_part2()
}
