package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func day2_part2() {

}

func min(nums []int) int {
	m := nums[0]
	for _, b := range nums {
		if b < m {
			m = b
		}
	}
	return m
}

func area(l, w, h int) int {
	// main
	a := 2*l*w + 2*w*h + 2*h*l
	// Min side
	t := []int{l, w, h}
	m1 := min(t)

	t2 := []int{}
	skipped := false
	for x := range t {
		if t[x] == m1 {
			if skipped {
				t2 = append(t2, t[x])
			} else {
				skipped = true
				continue
			}

		} else {
			t2 = append(t2, t[x])
		}
	}

	m2 := min(t2)

	// slack
	a += m1 * m2
	// fmt.Printf("%v Area=%d slack=%d, m1=%d,m2=%d\n", t2, a, m1*m2, m1, m2)
	return a
}

func day2_part1() {
	input := utils.ReadInput("./2015/day2/day2.input.txt")

	t := 0
	for _, dim := range input {
		s := strings.Split(dim, "x")
		l, _ := strconv.Atoi(string(s[0]))
		w, _ := strconv.Atoi(string(s[1]))
		h, _ := strconv.Atoi(string(s[2]))
		t += area(l, w, h)
	}
	fmt.Printf("t: %v\n", t)
}

func main() {
	day2_part1()
	day2_part2()

}
