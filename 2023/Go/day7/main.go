package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hand struct {
	Cards  string
	Points int
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

	for _, line := range contentsArr {

		parsedLine := strings.Split(strings.TrimSpace(line), " ")

		points, err := strconv.Atoi(parsedLine[1])
		if err != nil {
			fmt.Println("Not a number: ", parsedLine[1])
		}

		hands = append(hands, Hand{
			Cards:  parsedLine[0],
			Points: points,
            Strength: handParser(&parsedLine[0]),
		})
	}

    newHands := []Hand{}
    for _, v := range hands {
        newHands = handSort(newHands, v)
    }

    fmt.Println(hands)
    fmt.Println(newHands)

    total := 0
    for i, hand := range newHands {
        total = total + (i + 1) * hand.Points
    }

    fmt.Println(total)
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

func handSort(hands []Hand, newHand Hand) []Hand {
    left := 0
    right := len(hands) - 1

    for left <= right {
        mid := (left + right) / 2
        if hands[mid].Strength < newHand.Strength {
            left = mid + 1
        } else {
            right = mid - 1
        }
        // else if hands[mid].Strength == newHand.Strength {
        //
        // }
    }

    hands = append(hands[:left], append([]Hand{newHand}, hands[left:]...)...)
    copy(hands[left+1:], hands[left:])

    return hands
}
