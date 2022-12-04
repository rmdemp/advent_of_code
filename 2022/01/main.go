package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	sum := int64(0)
	highest := int64(0)

	for {
		line, err := rd.ReadString('\n')

		parsedStr := strings.Split(line, "\n")[0]
		stoi, _ := strconv.ParseInt(parsedStr, 10, 64)
		sum += stoi

		if parsedStr == "" {
			if sum > highest {
				highest = sum
			}
			sum = 0
		}

		if err == io.EOF {
			if sum > highest {
				highest = sum
			}
			fmt.Printf("highest %v", highest)
			return
		}

		if err != nil && err != io.EOF {
			log.Fatalf("err %v", err)
		}
	}
}