package main

import (
	"fmt"
	"strings"

	"github.com/joaooliveirapro/adventofcodego/utils"
)

func evalHand(myH, oppH string) int {
	t := 0
	if oppH == "A" { // Rock
		if myH == "X" { // Rock
			t += 3
		} else if myH == "Y" { // Paper
			t += 6
		} else { // Scissors
			t += 0
		}
	} else if oppH == "B" { // Paper
		if myH == "X" { // Rock
			t += 0
		} else if myH == "Y" { // Paper
			t += 3
		} else { // Scissors
			t += 6
		}
	} else { // Scissors
		if myH == "X" { // Rock
			t += 6
		} else if myH == "Y" { // Paper
			t += 0
		} else { // Scissors
			t += 3
		}
	}

	if myH == "X" { // Rock
		t += 1
	} else if myH == "Y" { // Paper
		t += 2
	} else { // Scissors
		t += 3
	}
	return t
}

func evalHand2(myH, oppH string) int {
	t := 0
	myHand := ""     // X-rock Y-paper Z-scissors
	if oppH == "A" { // Rock

		if myH == "X" { // Lose
			t += 0
			myHand = "Z"
		} else if myH == "Y" { // Draw
			t += 3
			myHand = "X"

		} else { // Win
			t += 6
			myHand = "Y"
		}

	} else if oppH == "B" { // Paper

		if myH == "X" { // Lose
			t += 0
			myHand = "X"
		} else if myH == "Y" { // Draw
			t += 3
			myHand = "Y"
		} else { // Win
			t += 6
			myHand = "Z"
		}

	} else { // Scissors
		if myH == "X" { // Lose
			t += 0
			myHand = "Y"
		} else if myH == "Y" { // Draw
			t += 3
			myHand = "Z"
		} else { // Win
			t += 6
			myHand = "X"
		}
	}

	if myHand == "X" { // Rock
		t += 1
	} else if myHand == "Y" { // Paper
		t += 2
	} else { // Scissors
		t += 3
	}
	return t
}

func part1(input *[]string) {
	t := 0
	for _, ch := range *input {
		f := strings.Fields(ch)
		myH, oppH := f[1], f[0]
		t += evalHand(myH, oppH)
	}
	fmt.Printf("t: %v\n", t)
}

func part2(input *[]string) {
	t := 0
	for _, ch := range *input {
		f := strings.Fields(ch)
		myH, oppH := f[1], f[0]
		t += evalHand2(myH, oppH)
	}
	fmt.Printf("t: %v\n", t)
}

func main() {
	input := utils.ReadInput("./2022/day2/day2.input.txt")
	part1(&input)
	part2(&input)
}
