package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    inputPath := os.Args[1]
    input, err := os.Open(inputPath)
    defer input.Close()
    if err != nil {
        panic(err)
    }

    fileScanner := bufio.NewScanner(input)
    fileScanner.Split(bufio.ScanLines)

    sum := 0

    for fileScanner.Scan() {
        rucksack := fileScanner.Text()
        compartment_left := make(map[rune]bool)
        compartment_right := make(map[rune]bool)

        for _, c := range rucksack[:len(rucksack) / 2] {
            compartment_left[c] = true
        }
        for _, c := range rucksack[len(rucksack) / 2:] {
            compartment_right[c] = true
        }

        for key := range compartment_left {
            if !compartment_right[key] {
                continue
            }

            if key >= 'a' && key <= 'z'{
                sum += int(key) - 'a' + 1
            } else if key >= 'A' && key <= 'Z' {
                sum += int(key) - 'A' + 1 + 26
            }
        }
    }

    fmt.Println(sum)
}
