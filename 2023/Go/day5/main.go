package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("../../inputs/day5.txt")
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(0)
	}

	contents := string(file)

	contentsArr := strings.Split(contents, "\n\n")

	seedsArr := strings.Split(contentsArr[0], ":")[1]
	seeds := strings.Split(strings.TrimSpace(seedsArr), " ")

	// lowestLocation := -1
	lowestLocationPartTwo := -1

	// for _, seed := range seeds {
	// 	seedNum, err := strconv.Atoi(seed)
	// 	if err != nil {
	// 		fmt.Println("Seed is not a number: ", seed)
	// 		os.Exit(0)
	// 	}
	// 	location := walkAlmanac(&contentsArr, seedNum, -1)
	//
	// 	if lowestLocation == -1 || location < lowestLocation {
	// 		lowestLocation = location
	// 	}
	// }

	//Part Two
	for i := 0; i < len(seeds); i += 2 {
		seedNum, err := strconv.Atoi(seeds[i])
		if err != nil {
			fmt.Println("Seed is not a number: ", seeds[i])
			os.Exit(0)
		}
		seedRange, err := strconv.Atoi(seeds[i+1])
		if err != nil {
			fmt.Println("Seed Range is not a number: ", seeds[i+1])
			os.Exit(0)
		}
		// fmt.Println("seed: ", seedNum)
		// fmt.Println("seed range: ", seedRange)
		maxSeed := seedNum + seedRange - 1

		// for j := seedNum; j <= maxSeed; j++ {
		location := walkAlmanac(&contentsArr, seedNum, maxSeed)
		if lowestLocationPartTwo == -1 || location < lowestLocationPartTwo {
			lowestLocationPartTwo = location
		}
		// fmt.Println(j)
		// }
	}

	// fmt.Println("Part One: ", lowestLocation)
	fmt.Println("Part Two: ", lowestLocationPartTwo)

}

func walkAlmanac(contents *[]string, seed, maxSeed int) int {
	contentsArr := *contents
	currPath := seed
	// lowestSeed := -1

	// fmt.Println("Seed: ", currPath)
	for i := 1; i < len(contentsArr); i++ {
		line := strings.Split(strings.TrimSpace(strings.Split(contentsArr[i], ":")[1]), "\n")
		pathFound := false

		for j := 0; j < len(line); j++ {
			nums := strings.Split(line[j], " ")
			destination, err := strconv.Atoi(nums[0])
			if err != nil {
				fmt.Println("Not a number: ", destination)
				os.Exit(0)
			}
			source, err := strconv.Atoi(nums[1])
			if err != nil {
				fmt.Println("Not a number: ", source)
				os.Exit(0)
			}
			rng, err := strconv.Atoi(nums[2])
			if err != nil {
				fmt.Println("Not a number: ", rng)
				os.Exit(0)
			}
			// fmt.Println("Umm ", j, destination, source, rng)

			if maxSeed > source+rng-1 && seed <= source {
				currPath = source
				pathFound = true
			}
			fmt.Println("Seed: ", currPath)

			if currPath >= source && currPath < source+rng && !pathFound {
				currPath = destination + (currPath - source)
				pathFound = true
			}
			// fmt.Println("Current Path: ", currPath)
		}
	}

	return currPath
}

// Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
// Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
// Seed 55, soil 57, fertilizer 57, water 53, light 46, temperature 82, humidity 82, location 86.
// Seed 13, soil 13, fertilizer 52, water 41, light 34, temperature 34, humidity 35, location 35.

// seeds: 79 14 55 13

// seed-to-soil map:
// 50 98 2
// 52 50 48
//
// soil-to-fertilizer map:
// 0 15 37
// 37 52 2
// 39 0 15
//
// fertilizer-to-water map:
// 49 53 8
// 0 11 42
// 42 0 7
// 57 7 4
//
// water-to-light map:
// 88 18 7
// 18 25 70
//
// light-to-temperature map:
// 45 77 23
// 81 45 19
// 68 64 13
//
// temperature-to-humidity map:
// 0 69 1
// 1 0 69
//
// humidity-to-location map:
// 60 56 37
// 56 93 4
