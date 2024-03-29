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
	fmt.Println("Part Two: ", cardCheckPartTwo(&contents))

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
			} else if v >= 1 {
				matchPoints++
			}
		}

		sum += matchPoints
	}

	return sum
}

func cardCheckPartTwo(contents *[]string) int {
	lines := *contents
	sum := 0
	cardCopies := make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		if cardCopies[i] == 0 {
			cardCopies[i] = 1
		}
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

		matches := 0
		for _, v := range scoreTable {
			if v >= 1 {
				matches++
			}
		}

		for k := i + 1; k < i+matches+1; k++ {
			if cardCopies[k] == 0 {
				cardCopies[k] = 1
			}
            cardCopies[k] += cardCopies[i]
		}

	}

	for i := 0; i < len(cardCopies); i++ {
		sum += cardCopies[i]
	}

	return sum
}
