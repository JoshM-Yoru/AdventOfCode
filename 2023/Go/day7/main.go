package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hand struct {
	Cards    string
	Points   int
	Strength int
}

var CardStrength map[string]int

func main() {

	file, err := os.ReadFile("../../inputs/day7.txt")
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(0)
	}

	contents := string(file)

	contentsArr := strings.Split(strings.TrimSpace(contents), "\n")

	CardStrength = make(map[string]int)

	CardStrength["A"] = 14
	CardStrength["K"] = 13
	CardStrength["Q"] = 12
	CardStrength["J"] = 11
	CardStrength["T"] = 10
	CardStrength["9"] = 9
	CardStrength["8"] = 8
	CardStrength["7"] = 7
	CardStrength["6"] = 6
	CardStrength["5"] = 5
	CardStrength["4"] = 4
	CardStrength["3"] = 3
	CardStrength["2"] = 2

	hands := []Hand{}
	partTwoHands := []Hand{}

	for _, line := range contentsArr {

		parsedLine := strings.Split(strings.TrimSpace(line), " ")

		points, err := strconv.Atoi(parsedLine[1])
		if err != nil {
			fmt.Println("Not a number: ", parsedLine[1])
		}

		hands = append(hands, Hand{
			Cards:    parsedLine[0],
			Points:   points,
			Strength: handParser(&parsedLine[0]),
		})

		partTwoHands = append(partTwoHands, Hand{
			Cards:    parsedLine[0],
			Points:   points,
			Strength: handParserPartTwo(&parsedLine[0]),
		})
	}

	newHands := []Hand{}
	for _, v := range hands {
		newHands = handSort(newHands, v)
	}

	newHandsPartTwo := []Hand{}
	for _, v := range partTwoHands {
		newHandsPartTwo = handSort(newHandsPartTwo, v)
	}

	// fmt.Println(hands)
	// fmt.Println(newHands)
	fmt.Println(newHandsPartTwo)

	total := 0
	for i, hand := range newHands {
		total = total + (i+1)*hand.Points
	}

	totalPartTwo := 0
	for i, handPartTwo := range newHandsPartTwo {
		totalPartTwo = totalPartTwo + (i+1)*handPartTwo.Points
	}

	fmt.Println("Part One: ", total)
	fmt.Println("Part Two: ", totalPartTwo)
}

func handParser(hand *string) int {

	mappedHand := make(map[string]int)

	for _, card := range *hand {
		mappedHand[string(card)]++
	}

	if len(mappedHand) == 1 {
		return 6
	}

	if len(mappedHand) == 2 {
		for _, v := range mappedHand {
			if v == 1 || v == 4 {
				return 5
			} else {
				return 4
			}
		}
	}

	if len(mappedHand) == 3 {
		for _, v := range mappedHand {
			if v == 3 {
				return 3
			}
		}
		return 2
	}

	if len(mappedHand) == 4 {
		return 1
	}

	return 0
}

func handParserPartTwo(hand *string) int {

	mappedHand := make(map[string]int)

	for _, card := range *hand {
		mappedHand[string(card)]++
	}
	// fmt.Println("Before: ", mappedHand)
	// checkForJoker := false

	if mappedHand["J"] > 0 && len(mappedHand) > 1 {
		// checkForJoker = true
		var mostFreq string
		highest := 0
		for k, v := range mappedHand {
			if v > highest && k != "J" {
				mostFreq = k
				highest = v
				// fmt.Println("K V Highest: ", k, v, highest)
				// fmt.Println("")
			}
		}
		mappedHand[mostFreq] = mappedHand[mostFreq] + mappedHand["J"]
		delete(mappedHand, "J")
	}

	// if checkForJoker {
		// fmt.Println("After: ", mappedHand)
	// }
	// fmt.Println("")

	if len(mappedHand) == 1 {
		return 6
	}

	if len(mappedHand) == 2 {
		for _, v := range mappedHand {
			if v == 1 || v == 4 {
				return 5
			} else {
				return 4
			}
		}
	}

	if len(mappedHand) == 3 {
		for _, v := range mappedHand {
			if v == 3 {
				return 3
			}
		}
		return 2
	}

	if len(mappedHand) == 4 {
		return 1
	}

	return 0
}

// func handParserPartTwo(hand *string) int {
//
// 	mappedHand := make(map[string]int)
//
// 	for _, card := range *hand {
// 		mappedHand[string(card)]++
// 	}
//
// 	if len(mappedHand) == 1 {
// 		return 6
// 	}
//
// 	if len(mappedHand) == 2 {
// 		for _, v := range mappedHand {
// 			if v == 1 || v == 4 {
//                 if mappedHand["J"] == 4 {
//                     return 5
//                 }
// 				return 5 + mappedHand["J"]
// 			} else {
//                 if mappedHand["J"] == 3 {
//                     return 6
//                 }
// 				return 4 + mappedHand["J"]
// 			}
// 		}
// 	}
//
// 	if len(mappedHand) == 3 {
// 		for _, v := range mappedHand {
// 			if v == 3 {
// 				if mappedHand["J"] == 1 {
// 					return 5
// 				}
// 				if mappedHand["J"] == 2 {
// 					return 6
// 				}
// 				return 3
// 			}
// 		}
//
// 		if mappedHand["J"] == 2 {
// 			return 5
// 		}
// 		return 2
// 	}
//
// 	if len(mappedHand) == 4 {
// 		return 1 + mappedHand["J"]
// 	}
//
// 	return 0 + mappedHand["J"]
// }

func handSort(hands []Hand, newHand Hand) []Hand {
	idx := 0

Outer:
	for idx < len(hands) {
		if hands[idx].Strength == newHand.Strength {
		Inner:
			for j := 0; j < len(hands[idx].Cards); j++ {
				if CardStrength[string(hands[idx].Cards[j])] > CardStrength[string(newHand.Cards[j])] {
					// fmt.Println("Break on cards check ", idx, j)
					break Outer
				} else if CardStrength[string(hands[idx].Cards[j])] < CardStrength[string(newHand.Cards[j])] {
					break Inner
				}
			}
		} else if hands[idx].Strength > newHand.Strength {
			// fmt.Println("Break on strength check ", idx)
			break Outer
		}
		idx++
	}

	hands = append(hands[:idx], append([]Hand{newHand}, hands[idx:]...)...)
	// fmt.Println(hands)

	return hands
}
