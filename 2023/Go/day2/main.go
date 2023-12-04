package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../../inputs/day2.txt")
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(0)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var contents []string

	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading line")
	}

	fmt.Println("Part One: ", gameCheckOne(&contents))
	fmt.Println("Part Two: ", gameCheckTwo(&contents))

}

func gameCheckOne(contents *[]string) int {
	lines := *contents
	sum := 0

	table := make(map[string]int)
	table["red"] = 12
	table["green"] = 13
	table["blue"] = 14

	for i := 0; i < len(lines); i++ {
		cleanedLine := strings.ReplaceAll(lines[i], ",", "")
		cleanedLine = strings.ReplaceAll(cleanedLine, ";", "")
		line := strings.Split(cleanedLine, " ")

		possible := true

	Line:
		for j := 2; j < len(line); j += 2 {
			num, err := strconv.Atoi(line[j])
			if err != nil {
				fmt.Println("Not a number: ", lines[i], line[j])
			}

			if num > table[line[j+1]] {
				possible = false
				break Line
			}
		}

		if possible {
			sum = sum + i + 1
		}
	}

	return sum
}

func gameCheckTwo(file *[]string) int {
	lines := *file
	sum := 0

	table := make(map[string]int)
	table["red"] = 0
	table["green"] = 0
	table["blue"] = 0

	for i := 0; i < len(lines); i++ {
		cleanedLine := strings.ReplaceAll(lines[i], ",", "")
		cleanedLine = strings.ReplaceAll(cleanedLine, ";", "")
		line := strings.Split(cleanedLine, " ")

		var power int

		for j := 2; j < len(line); j += 2 {
			num, err := strconv.Atoi(line[j])
			if err != nil {
				fmt.Println("Not a number: ", lines[i], line[j])
			}

			if num > table[line[j+1]] {
				table[line[j+1]] = num
			}
		}

		power = table["red"] * table["green"] * table["blue"]

		sum += power

		table["red"] = 0
		table["green"] = 0
		table["blue"] = 0
	}

	return sum
}
