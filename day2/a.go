package main

import (
	"bufio"
	"fmt"
	"os"
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

  shapeMap := map[byte]int {
    'X': 1,
    'A': 1,
    'Y': 2,
    'B': 2,
    'Z': 3,
    'C': 3,
  }

  score := 0

  for fileScanner.Scan() {
    opponent := fileScanner.Text()[0]
    your := fileScanner.Text()[2] 
    score += shapeMap[your]
    if shapeMap[your] == shapeMap[opponent] {
      // draw
      score += 3
    } else if shapeMap[your] - shapeMap[opponent] == 1 {
      // win
      score += 6
    } else if shapeMap[your] == 1 && shapeMap[opponent] == 3 {
      // win
      score += 6
    } else {
      // lose
      score += 0
    }
  }
  fmt.Println(score)
}
