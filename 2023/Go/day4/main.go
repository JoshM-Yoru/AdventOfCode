package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	file, err := os.Open("../../inputs/day4.txt")
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

	fmt.Println("Part One: ", cardCheck(&contents))

}

func cardCheck(contents *[]string) int {
	lines := *contents
	sum := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		scoreTable := make(map[string]int)
		checkMyCard := false
		numString := ""

		for j := 8; j < len(line); j++ {
			character := line[j]
			if character == 124 {
				checkMyCard = true
			}
			if unicode.IsDigit(rune(character)) {
				numString += string(character)
			} else if len(numString) > 0 && !unicode.IsDigit(rune(character)) && !checkMyCard {
				scoreTable[numString] = 0
				numString = ""
			}

			if checkMyCard && len(numString) > 0 && !unicode.IsDigit(rune(character)) {
				if _, ok := scoreTable[numString]; ok {
					scoreTable[numString] = 1
				}
                numString = ""
			}
		}

		if len(numString) > 0 && checkMyCard {
			if _, ok := scoreTable[numString]; ok {
				scoreTable[numString] = 1
			}
		}

		numString = ""

		matchPoints := 0
		for _, v := range scoreTable {
			if matchPoints >= 1 && v >= 1 {
				matchPoints *= 2
            } else if v >= 1  {
				matchPoints++
			} 
		}

		sum += matchPoints
	}

	return sum
}
