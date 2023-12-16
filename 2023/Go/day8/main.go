package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Branches struct {
	L string
	R string
}

func main() {
	file, err := os.ReadFile("../../inputs/day8.txt")
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(0)
	}

	contents := strings.Split(strings.TrimSpace(string(file)), "\n\n")
	instructions := contents[0]
	network := strings.Split(strings.TrimSpace(contents[1]), "\n")

	networkMap := make(map[string]Branches)
	var startingPostions []string
	for _, line := range network {
		lineSplit := strings.Split(line, "=")
		key := strings.TrimSpace(lineSplit[0])
		if string(key[2]) == "A" {
			startingPostions = append(startingPostions, key)
		}
		branches := strings.Split(lineSplit[1], ",")
		var trimmed Branches
		for i, v := range branches {
			if i == 0 {
				trimmed.L = strings.TrimSpace(v)
				trimmed.L = strings.Trim(trimmed.L, "(")
			} else {
				trimmed.R = strings.TrimSpace(v)
				trimmed.R = strings.Trim(trimmed.R, ")")
			}
		}
		networkMap[key] = trimmed
	}

	currPos := "AAA"
	count := 0
	idx := 0
	for idx < len(instructions) {
		if currPos == "ZZZ" {
			break
		}
		count++
		if string(instructions[idx]) == "L" {
			currPos = networkMap[currPos].L
		} else {
			currPos = networkMap[currPos].R
		}
		if idx+1 == len(instructions) {
			idx = 0
		} else {
			idx++
		}
	}

	fmt.Println("Part One:", count)

	currPositions := startingPostions
	partTwoCount := 0
	partTwoIdx := 0
	intervals := make([]int, len(startingPostions))
	for partTwoIdx < len(instructions) {
		for i, currPos := range currPositions {
			if string(currPos[2]) == "Z" {
				intervals[i] = partTwoCount
			}
		}
		foundIntervals := true
		for _, v := range intervals {
			if v == 0 {
				foundIntervals = false
			}
		}
		if foundIntervals {
			break
		}
		partTwoCount++
		var tempPositions []string
		for _, currPos := range currPositions {
			if string(instructions[partTwoIdx]) == "L" {
				tempPositions = append(tempPositions, networkMap[currPos].L)
			} else {
				tempPositions = append(tempPositions, networkMap[currPos].R)
			}
		}
		currPositions = tempPositions

		if partTwoIdx+1 == len(instructions) {
			partTwoIdx = 0
		} else {
			partTwoIdx++
		}
	}

	fmt.Println("Part Two:", lcmOfArray(intervals))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

func lcmOfArray(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}
