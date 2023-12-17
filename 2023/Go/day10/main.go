package main

import (
	"fmt"
	"os"
	"strings"
)

var PipeTypes map[string][][]int

func main() {

	PipeTypes = make(map[string][][]int)

	PipeTypes["|"] = [][]int{{0, -1}, {0, 1}}
	PipeTypes["-"] = [][]int{{-1, 0}, {1, 0}}
	PipeTypes["L"] = [][]int{{0, 1}, {1, 0}}
	PipeTypes["J"] = [][]int{{0, 1}, {-1, 0}}
	PipeTypes["7"] = [][]int{{0, -1}, {-1, 0}}
	PipeTypes["F"] = [][]int{{0, -1}, {1, 0}}

	file, err := os.ReadFile("../../inputs/day10.txt")
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(0)
	}

	content := strings.Split(strings.TrimSpace(string(file)), "\n")
	var start []int
	var currPos [][]int

	// steps := 0
	for i, line := range content {
		for j, character := range line {
			if string(character) == "S" {
				start = []int{j, i}
			}
		}
        fmt.Println(line)
	}

    fmt.Println("")

	currPos = append(currPos, start)

	directions := [][]int{
		{0, -1},
		{1, 0},
		{0, 1},
        {-1, 0},
	}

    foundEnd := false

	for !foundEnd {
		for _, d := range directions {
			currPosX := currPos[len(currPos)-1][0]
			currPosY := currPos[len(currPos)-1][1]
            newPosX := currPos[len(currPos)-1][0]+d[0]
            newPosY := currPos[len(currPos)-1][0]+d[1]
			if currPosX+d[0] > 0 && currPosX+d[0] < len(content[0]) && currPosY+d[1] >= 0 && currPosY+d[1] < len(content) {
                for _, p := range PipeTypes {
                    if string(content[currPosY+d[1]][currPosX+d[0]]) == p {

                    }
                }
				fmt.Println(string(content[currPosY+d[1]][currPosX+d[0]]))
			}
		}
        if start[0] != currPos[len(currPos)-1][0] && start[1] != currPos[len(currPos)-1][1] {
            foundEnd = true
        }
        break
	}

	fmt.Println(start)

}
