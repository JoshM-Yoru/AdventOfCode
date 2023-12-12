package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AlmanacEntry struct {
	sourceLo      int
	sourceHi      int
	destinationLo int
	destinationHi int
}

type PathRange struct {
	Lo int
	Hi int
}

func main() {
	file, err := os.ReadFile("../../inputs/day5.txt")
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(0)
	}

	contents := string(file)

	contentsArr := strings.Split(contents, "\n\n")

	seedsContents := strings.Split(contentsArr[0], ":")[1]
	seeds := strings.Split(strings.TrimSpace(seedsContents), " ")

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
	seedsArray := []PathRange{}

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
		maxSeed := seedNum + seedRange - 1

		seedsArray = append(seedsArray, PathRange{
			Lo: seedNum,
			Hi: maxSeed,
		})
	}
	almanac := getAlmanac(&contentsArr)
	lowestLocationPartTwo = walkAlmanacPartTwo(&almanac, seedsArray)
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

			// if maxSeed > source+rng-1 && seed <= source {
			// 	currPath = source
			//              // lowestSeed = source
			//              maxSeed = source+rng-1
			// 	pathFound = true
			// }
			// fmt.Println("Seed: ", currPath)

			if currPath >= source && currPath < source+rng && !pathFound {
				currPath = destination + (currPath - source)
				pathFound = true
			}
			// fmt.Println("Current Path: ", currPath)
		}
	}

	return currPath
}

func getAlmanac(contents *[]string) [][]AlmanacEntry {
	contentsArr := *contents
	almanac := [][]AlmanacEntry{}

	for i := 1; i < len(contentsArr); i++ {
		line := strings.Split(strings.TrimSpace(strings.Split(contentsArr[i], ":")[1]), "\n")
		subArray := []AlmanacEntry{}

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
			maxSource := source + rng - 1
			maxDestination := destination + rng - 1

			entry := AlmanacEntry{
				sourceLo:      source,
				sourceHi:      maxSource,
				destinationLo: destination,
				destinationHi: maxDestination,
			}

			subArray = append(subArray, entry)

		}
		almanac = append(almanac, subArray)
	}

	// fmt.Println(almanac)

	return almanac
}

// func walkAlmanacPartTwo(almanac *[][]AlmanacEntry, seeds []PathRange) int {
// 	alm := *almanac
//
//     // fmt.Println(seeds)
//     // fmt.Println()
//     //
//     // fmt.Println(alm[0])
//     // fmt.Println(alm[1])
//     // fmt.Println(alm[2])
//     // fmt.Println(alm[3])
//     // fmt.Println(alm[4])
//     // fmt.Println(alm[5])
//     // fmt.Println(alm[6])
//     //
// 	for _, section := range alm {
// 		newRanges := []PathRange{}
// 		for len(seeds) > 0 {
// 			seedRange := pop(&seeds)
//
//             // fmt.Print(i)
//             // fmt.Println(seeds)
// 			for _, almEntry := range section {
// 				newPath := PathRange{
// 					Lo: max(seedRange.Lo, almEntry.sourceLo),
// 					Hi: min(seedRange.Hi, almEntry.sourceHi),
// 				}
//
// 				if newPath.Lo < newPath.Hi {
// 					newRanges = append(newRanges, PathRange{
// 						Lo: newPath.Lo - almEntry.sourceLo + almEntry.destinationLo,
// 						Hi: newPath.Hi - newPath.Lo + almEntry.destinationLo,
// 					})
//
// 					if newPath.Lo > seedRange.Lo {
// 						seeds = append(seeds, PathRange{
// 							Lo: seedRange.Lo,
// 							Hi: newPath.Lo,
// 						})
// 					}
// 					if seedRange.Hi > newPath.Hi {
// 						seeds = append(seeds, PathRange{
// 							Lo: newPath.Hi,
// 							Hi: seedRange.Hi,
// 						})
// 					}
// 					break
// 				}
// 			}
//             // fmt.Println(seeds)
//             if seedRange.Lo < seedRange.Hi {
//                 newRanges = append(newRanges, seedRange)
//             }
// 		}
// 		seeds = newRanges
// 	}
// 	// fmt.Println("Seeds: ", seeds)
//
// 	min := seeds[0].Lo
// 	for _, v := range seeds {
// 		if v.Lo < min {
// 			min = v.Lo
// 		}
// 	}
//
// 	return min
// }

func walkAlmanacPartTwo(almanac *[][]AlmanacEntry, seeds []PathRange) int {
	alm := *almanac

for _, section := range alm {
    newRanges := []PathRange{}
    for len(seeds) > 0 {
        seedRange := pop(&seeds)

        fmt.Println("Processing seed range:", seedRange)

        for _, almEntry := range section {
            newPath := PathRange{
                Lo: max(seedRange.Lo, almEntry.sourceLo),
                Hi: min(seedRange.Hi, almEntry.sourceHi),
            }

            fmt.Println("New path:", newPath)

            if newPath.Lo < newPath.Hi {
                newDestLo := newPath.Lo - almEntry.sourceLo + almEntry.destinationLo
                newDestHi := newPath.Hi - newPath.Lo + almEntry.destinationHi

                fmt.Printf("Adding to newRanges: Lo=%d, Hi=%d\n", newDestLo, newDestHi)

                newRanges = append(newRanges, PathRange{
                    Lo: newDestLo,
                    Hi: newDestHi,
                })

                if newPath.Lo > seedRange.Lo {
                    seeds = append(seeds, PathRange{
                        Lo: seedRange.Lo,
                        Hi: newPath.Lo,
                    })
                }
                if seedRange.Hi > newPath.Hi {
                    seeds = append(seeds, PathRange{
                        Lo: newPath.Hi,
                        Hi: seedRange.Hi,
                    })
                }
                break
            }
        }

        if seedRange.Lo < seedRange.Hi {
            newRanges = append(newRanges, seedRange)
        }
    }
    seeds = newRanges
}

// Print seeds at the end of the function
// fmt.Println("Final seeds:", seeds)

	min := seeds[0].Lo
	for _, v := range seeds {
		if v.Lo < min {
            fmt.Println(v.Lo)
            fmt.Println(v.Hi)
            fmt.Println()
			min = v.Lo
		}
	}

	return min
}

func pop(arr *[]PathRange) PathRange {
	f := len(*arr)
	rv := (*arr)[f-1]
	*arr = (*arr)[:f-1]
	return rv
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
