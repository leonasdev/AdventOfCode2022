package main

import (
	"bufio"
  "fmt"
	"os"
	"strconv"
)

func FindMaxAndReturnSum(groups [][]int) int {
  max := 0
  for _, group := range groups {
    sum := 0
    for _, num := range group {
      sum += num
    }
    if sum >= max {
      max = sum
    }
  }

  return max
}

func main() {
  inputPath := os.Args[1]
  input, err  := os.Open(inputPath)
  defer input.Close()
  if err != nil {
    panic(err)
  }
  
  fileScanner := bufio.NewScanner(input)
  fileScanner.Split(bufio.ScanLines)

  var groups [][]int
  var group []int

  for fileScanner.Scan() {
    if fileScanner.Text() == "" {
      groups = append(groups, group)
      group = nil
    } else {
      line, err := strconv.Atoi(fileScanner.Text())
      if err != nil {
        panic(err)
      }
      group = append(group, line)
    }
  }

  fmt.Println(FindMaxAndReturnSum(groups))
}
