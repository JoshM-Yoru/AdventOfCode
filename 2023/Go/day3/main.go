package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("../../inputs/day3.txt")
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

	fmt.Println("Part One: ", partNumberCheck(&contents))
	fmt.Println("Part Two: ", gearFinder(&contents))

}

func partNumberCheck(contents *[]string) int {
	lines := *contents
	sum := 0

	for i := 0; i < len(lines); i++ {
		numberString := ""
		var numberStart int
		foundNumberStart := false
		var numberEnd int
		line := lines[i]

		for j := 0; j < len(line); j++ {

			character := line[j]

			if unicode.IsDigit(rune(character)) && !foundNumberStart {
				numberString += string(character)
				foundNumberStart = true
				numberStart = j
			} else if unicode.IsDigit(rune(character)) {
				numberString += string(character)
			} else if (len(numberString) > 0) && !unicode.IsDigit(rune(character)) {
				numberEnd = j - 1
				if borderWalk(&lines, i, numberStart, numberEnd) {
					num, err := strconv.Atoi(numberString)
					if err != nil {
						fmt.Println("Not a number")
						os.Exit(0)
					}
					sum += num
				}
				numberStart = 0
				numberEnd = 0
				foundNumberStart = false
				numberString = ""
			}

			if (len(numberString) > 0) && j == len(line)-1 {
				numberEnd = j - 1
				if borderWalk(&lines, i, numberStart, numberEnd) {
					num, err := strconv.Atoi(numberString)
					if err != nil {
						fmt.Println("Not a number")
						os.Exit(0)
					}
					sum += num
				}
				numberStart = 0
				numberEnd = 0
				foundNumberStart = false
				numberString = ""
			}

		}
	}

	return sum
}

func borderWalk(contents *[]string, line, start, end int) bool {
	lines := *contents
	var borderStart int
	if start-1 >= 0 {
		borderStart = start - 1
	} else {
		borderStart = start
	}

	var borderEnd int
	if end+1 < len(lines[line]) {
		borderEnd = end + 1
	} else {
		borderEnd = end
	}

	var topBorder int
	if line-1 >= 0 {
		topBorder = line - 1
	} else {
		topBorder = -1
	}

	var bottomBorder int
	if line+1 < len(lines) {
		bottomBorder = line + 1
	} else {
		bottomBorder = -1
	}

	if symbolCheck(lines[line][borderStart]) || symbolCheck(lines[line][borderEnd]) {
		return true
	}
	for i := borderStart; i <= borderEnd; i++ {
		if topBorder != -1 && symbolCheck(lines[topBorder][i]) {
			return true
		}
		if bottomBorder != -1 && symbolCheck(lines[bottomBorder][i]) {
			return true
		}
	}

	return false
}

func symbolCheck(character byte) bool {
	if int(character) >= 33 && int(character) <= 47 && int(character) != 46 {
		return true
	}

	if int(character) >= 58 && int(character) <= 64 {
		return true
	}

	if int(character) >= 91 && int(character) <= 96 {
		return true
	}

	if int(character) >= 123 && int(character) <= 126 {
		return true
	}

	return false
}

// Part Two

type Gear struct {
	x int
	y int
}

type Part struct {
	ratio  int
	amount int
}

func gearFinder(contents *[]string) int {
	lines := *contents
	sum := 0

	gears := make(map[Gear][]int)

	for i := 0; i < len(lines); i++ {
		numberString := ""
		var numberStart int
		foundNumberStart := false
		var numberEnd int
		line := lines[i]

		for j := 0; j < len(line); j++ {

			character := line[j]

			if unicode.IsDigit(rune(character)) && !foundNumberStart {
				numberString += string(character)
				foundNumberStart = true
				numberStart = j
			} else if unicode.IsDigit(rune(character)) {
				numberString += string(character)
			} else if (len(numberString) > 0) && !unicode.IsDigit(rune(character)) {
				numberEnd = j - 1
				num, err := strconv.Atoi(numberString)
				if err != nil {
					fmt.Println("Not a number")
					os.Exit(0)
				}
				borderWalkPartTwo(&lines, i, numberStart, numberEnd, num, gears)
				numberStart = 0
				numberEnd = 0
				foundNumberStart = false
				numberString = ""
			}

			if (len(numberString) > 0) && j == len(line)-1 {
				numberEnd = j - 1
				num, err := strconv.Atoi(numberString)
				if err != nil {
					fmt.Println("Not a number")
					os.Exit(0)
				}
				borderWalkPartTwo(&lines, i, numberStart, numberEnd, num, gears)
				numberStart = 0
				numberEnd = 0
				foundNumberStart = false
				numberString = ""
			}

		}
	}

	for _, v := range gears {
		if v[1] > 1 {
			sum += v[0]
		}
	}

	return sum
}

func borderWalkPartTwo(contents *[]string, line, start, end, currentNum int, gears map[Gear][]int) {
	lines := *contents
	var borderStart int
	if start-1 >= 0 {
		borderStart = start - 1
	} else {
		borderStart = start
	}

	var borderEnd int
	if end+1 < len(lines[line]) {
		borderEnd = end + 1
	} else {
		borderEnd = end
	}

	var topBorder int
	if line-1 >= 0 {
		topBorder = line - 1
	} else {
		topBorder = -1
	}

	var bottomBorder int
	if line+1 < len(lines) {
		bottomBorder = line + 1
	} else {
		bottomBorder = -1
	}

	if gearCheck(lines[line][borderStart]) {
		if val, ok := gears[Gear{
			x: borderStart,
			y: line,
		}]; !ok {
			addKeyValue(borderStart, line, currentNum, gears)
		} else {
			val[0] *= currentNum
			val[1]++
		}
	}
	if gearCheck(lines[line][borderEnd]) {
		if val, ok := gears[Gear{
			x: borderEnd,
			y: line,
		}]; !ok {
			addKeyValue(borderEnd, line, currentNum, gears)
		} else {
			val[0] *= currentNum
			val[1]++
		}
	}

	for i := borderStart; i <= borderEnd; i++ {
		if topBorder != -1 && gearCheck(lines[topBorder][i]) {
			if val, ok := gears[Gear{
				x: i,
				y: topBorder,
			}]; !ok {
				addKeyValue(i, topBorder, currentNum, gears)
			} else {
				val[0] *= currentNum
				val[1]++
			}
		}
		if bottomBorder != -1 && gearCheck(lines[bottomBorder][i]) {
			if val, ok := gears[Gear{
				x: i,
				y: bottomBorder,
			}]; !ok {
				addKeyValue(i, bottomBorder, currentNum, gears)
			} else {
				val[0] *= currentNum
				val[1]++
			}
		}
	}

}

func addKeyValue(posX, posY, num int, gears map[Gear][]int) {
	gears[Gear{
		x: posX,
		y: posY,
	}] = []int{num, 1}
}

func gearCheck(character byte) bool {
	if int(character) == 42 {
		return true
	}

	return false
}
