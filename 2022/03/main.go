package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

  priorities := ""
  uniqueString := ""
  var prioritiesTotal rune

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatalf("%v", err)
		}

    firstHalf := line[:len(strings.Trim(line, "")) / 2]
    secondHalf := strings.Trim(line[len(strings.Trim(line, "")) / 2:], "\n")
    
    for _, r := range firstHalf {
      if strings.Contains(secondHalf, string(r)) {
        priorities += string(r)
      }
    }

    if err == io.EOF {
      break
    }
  }

  for _, r := range priorities {
    if !strings.Contains(uniqueString, string(r)) {
      uniqueString += string(r)
    }
  }

  for _, r := range uniqueString {
    if unicode.IsLower(r) {
      prioritiesTotal += r - 96
    } else if unicode.IsUpper(r) { 
      prioritiesTotal += r - 38
    }
  }

  fmt.Printf("uniqueString %v\n", uniqueString)
  fmt.Printf("prioritiesTotal %v\n", prioritiesTotal)
}