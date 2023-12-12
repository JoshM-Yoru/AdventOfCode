package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("../../inputs/day6.txt")
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(0)
	}

	contents := string(file)

	contentsArr := strings.Split(strings.TrimSpace(contents), "\n")
	measurements := [][]int{}
    partTwoMeasurements := []string{}

	for _, line := range contentsArr {
		measures := []int{}
		number := 0
        stringMeasure := ""

		trimmedLine := strings.TrimSpace(line)

		for _, character := range trimmedLine {
			num, err := strconv.Atoi(string(character))
			if err != nil {
				if number > 0 {
					measures = append(measures, number)
					number = 0
				}
			}
            if num >= 0 && err == nil {
                stringMeasure += string(character)
            }
			if number > 0 {
				number = number*10 + num
			} else {
				number = num
			}
            // fmt.Println(stringMeasure)
		}
		if number > 0 {
			measures = append(measures, number)
			measurements = append(measurements, measures)
            partTwoMeasurements = append(partTwoMeasurements, stringMeasure)
		}
	}
	fmt.Println(measurements)
    fmt.Println("Part 2 Measurements:", partTwoMeasurements)

	fmt.Println(waysToWin(&measurements))
	fmt.Println(waysToWinPartTwo(&partTwoMeasurements))
}

func waysToWin(measurements *[][]int) int {
	winsPerRace := []int{}

	races := *measurements

	for i := 0; i < len(races[0]); i++ {
		raceTime := races[0][i]
		raceDist := races[1][i]
		var waysToTry int64
		waysToTry = int64(raceTime) / 2

		wins := 0

		for j := 1; j <= int(waysToTry); j++ {
			if (raceTime-j)*j > raceDist {
				wins += 2
			}
		}
        if raceTime&1 == 0 {
            wins--
        }
		winsPerRace = append(winsPerRace, wins)
	}

    total := 1
    for _, v := range winsPerRace {
        total *= v
    }

    return total
}

func waysToWinPartTwo(measurements *[]string) int {

	race := *measurements

		raceTime, err := strconv.Atoi(race[0])
        if err != nil {
            fmt.Println("Not a number: ", race[0])
        }
		raceDist, err := strconv.Atoi(race[1])
        if err != nil {
            fmt.Println("Not a number: ", race[1])
        }
		var waysToTry int64
		waysToTry = int64(raceTime) / 2

		wins := 0

		for j := 1; j <= int(waysToTry); j++ {
			if (raceTime-j)*j > raceDist {
				wins += 2
			}
		}
        if raceTime&1 == 0 {
            wins--
        }

    return wins
}
