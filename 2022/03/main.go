package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("simple.txt")
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

  priorities := make([]string, 0)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatalf("%v", err)
		}

    if err == io.EOF {
      return
    }

    trimLine := strings.Trim(line, "\n")
		firstHalf := line[:len(trimLine) / 2]
		secondHalf := line[len(trimLine) / 2:]
    trimSecondHalf := secondHalf[:len(secondHalf) - 1] // why does this happen?

    fmt.Printf("%v\n", firstHalf)
    for _, b := range strings.Split(firstHalf, "") {
      fmt.Printf("%v contains %v %v\n", trimSecondHalf, b, strings.Contains(trimSecondHalf, b))
      if strings.Contains(trimSecondHalf, b) {
        priorities = append(priorities, b)
      }
    }
    fmt.Printf("priorities %v\n\n", priorities)    
	}
}