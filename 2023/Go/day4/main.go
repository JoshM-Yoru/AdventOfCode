package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {

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
    
}

func cardCheck(contents *[]string) int {
    lines := *contents
    sum := 0

    for i := 0; i < len(lines); i++ {
        line := lines[i]

        scoreTable := make(map[string]int)

        for j := 10; j < len(line); j++ {
            fmt.Print(string(line[j]))
        }

        fmt.Println()
    }

    return sum
}
