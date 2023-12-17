package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("../../inputs/day9.txt")
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(0)
	}

	content := strings.Split(strings.TrimSpace(string(file)), "\n")

	var readings [][]int
	for _, line := range content {
		lineArr := strings.Split(line, " ")
		var nums []int
		for _, v := range lineArr {
			num, err := strconv.Atoi(string(v))
			if err != nil {
				fmt.Println("Not a Number: ", string(v))
			}
			nums = append(nums, num)
		}
		readings = append(readings, nums)
	}

    partOneTotal := 0
    for _, line := range readings {
        partOneTotal += sequencingAlgo(line)
    }

    fmt.Println("Part One: ", partOneTotal)

    partTwoTotal := 0
    for _, line := range readings {
        partTwoTotal += sequencingHistory(line)
    }
    fmt.Println("Part Two: ", partTwoTotal)

}

func sequencingAlgo(nums []int) int {

	nextNum := 0

	lowestRow := nums

	for !zeroedArray(lowestRow) {
		var lowerRow []int
		nextNum += lowestRow[len(lowestRow)-1]
		for i := 1; i < len(lowestRow); i++ {
			lowerRow = append(lowerRow, lowestRow[i]-lowestRow[i-1])
		}
		lowestRow = lowerRow
	}

	return nextNum
}

func sequencingHistory(nums []int) int {

	lowestRow := nums

    var totalArray [][]int
    totalArray = append(totalArray, nums)
	for !zeroedArray(lowestRow) {
		var lowerRow []int
		for i := 1; i < len(lowestRow); i++ {
			lowerRow = append(lowerRow, lowestRow[i]-lowestRow[i-1])
		}
		lowestRow = lowerRow
        totalArray = append(totalArray, lowerRow)
	}

    historicalNum := 0
    for i := len(totalArray) - 1; i >= 0 ; i-- {
        historicalNum = totalArray[i][0] - historicalNum
    }

	return historicalNum
}

func zeroedArray(nums []int) bool {
	for _, v := range nums {
		if v != 0 {
			return false
		}
	}
	return true
}
