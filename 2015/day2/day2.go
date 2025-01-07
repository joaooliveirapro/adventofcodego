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

func area(l, w, h int) (int, int) {
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

	// ribbon (part 2)
	ribbon := 2*m1 + 2*m2 + l*w*h

	return a, ribbon
}

func day2_part1() {
	input := utils.ReadInput("./2015/day2/day2.input.txt")

	t := 0
	ribbon := 0
	for _, dim := range input {
		s := strings.Split(dim, "x")
		l, _ := strconv.Atoi(string(s[0]))
		w, _ := strconv.Atoi(string(s[1]))
		h, _ := strconv.Atoi(string(s[2]))
		t_, ribbon_ := area(l, w, h)
		t += t_
		ribbon += ribbon_
	}
	fmt.Printf("t: %v r: %v\n", t, ribbon)
}

func main() {
	day2_part1()
	day2_part2()
}
