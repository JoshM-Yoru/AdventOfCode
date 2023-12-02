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

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading line")
	}

    fmt.Println("Part One: ", gameCheckOne(&lines))

}

func gameCheckOne(file *[]string) int {
	lines := *file
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
