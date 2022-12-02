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
    'A': 1,
    'B': 2,
    'C': 3,
    'X': 0,
    'Y': 3,
    'Z': 6,
  }

  score := 0

  for fileScanner.Scan() {
    opponent := fileScanner.Text()[0]
    round := fileScanner.Text()[2] 
    score += shapeMap[round]
    switch round {
    case 'X':
      // lose
      if opponent == 'A' {
        score += shapeMap['C']
      } else if opponent == 'B' {
        score += shapeMap['A']
      } else {
        score += shapeMap['B']
      }
      break
    case 'Y':
      // draw
      if opponent == 'A' {
        score += shapeMap['A']
      } else if opponent == 'B' {
        score += shapeMap['B']
      } else {
        score += shapeMap['C']
      }
      break
    case 'Z':
      // win
      if opponent == 'A' {
        score += shapeMap['B']
      } else if opponent == 'B' {
        score += shapeMap['C']
      } else {
        score += shapeMap['A']
      }
      break
    }
  }
  fmt.Println(score)
}
