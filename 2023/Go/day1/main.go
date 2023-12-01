package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
    file, err := os.Open("../../inputs/day1.txt")
    if err != nil {
        fmt.Println("Error opening file")
        os.Exit(0)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    var lines []string

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading line")
    }

    partOne(&lines)
    partTwo(&lines)
}

func partOne(file *[]string) {
    lines := *file
    sum := 0

    for i := 0; i < len(lines); i++ {
        var firstNumber string
        var lastNumber string
        foundFirstNumber := false
        foundAnotherNumber := false

        for j := 0; j < len(lines[i]); j++ {
            character := lines[i][j]
            if unicode.IsDigit(rune(character)) && !foundFirstNumber {
                firstNumber = string(character)
                foundFirstNumber = true
            } else if unicode.IsDigit(rune(character)) && foundFirstNumber {
                lastNumber = string(character)
                foundAnotherNumber = true
            }
        }

        if !foundAnotherNumber {
            lastNumber = firstNumber
        }

        value, err := strconv.Atoi(firstNumber + lastNumber)
        if err != nil {
            fmt.Println("error with value")
            fmt.Println(firstNumber)
            fmt.Println(lastNumber)
        }

        sum += value
    }

    fmt.Println(sum)
}

func partTwo(file *[]string) {
    lines := *file
    sum := 0

    for i := 0; i < len(lines); i++ {
        var firstNumber string
        var lastNumber string
        foundFirstNumber := false
        foundAnotherNumber := false

        for j := 0; j < len(lines[i]); j++ {
            character := lines[i][j]
            if unicode.IsDigit(rune(character)) && !foundFirstNumber {
                firstNumber = string(character)
                foundFirstNumber = true
            } else if unicode.IsDigit(rune(character)) && foundFirstNumber {
                lastNumber = string(character)
                foundAnotherNumber = true
            } 
        }

        if !foundAnotherNumber {
            lastNumber = firstNumber
        }

        value, err := strconv.Atoi(firstNumber + lastNumber)
        if err != nil {
            fmt.Println("error with value")
            fmt.Println(firstNumber)
            fmt.Println(lastNumber)
        }

        sum += value
    }

    fmt.Println(sum)
}
