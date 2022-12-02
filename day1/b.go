package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
  inputPath := os.Args[1]
  input, err  := os.Open(inputPath)
  defer input.Close()
  if err != nil {
    panic(err)
  }
  
  fileScanner := bufio.NewScanner(input)
  fileScanner.Split(bufio.ScanLines)

  var groups []int

  sum := 0
  for fileScanner.Scan() {
    if fileScanner.Text() == "" {
      groups = append(groups, sum)
      sum = 0
    } else {
      line, err := strconv.Atoi(fileScanner.Text())
      if err != nil {
        panic(err)
      }
      sum += line
    }
  }

  sort.Sort(sort.Reverse(sort.IntSlice(groups)))
  fmt.Printf("%v", groups[0]+groups[1]+groups[2])
}
