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
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	type Outcomes struct {
		draw, win int
	}

	type Scores struct {
		X, Y, Z int
	}

	outcomes := Outcomes{
		draw: 3,
		win: 6,
	}

	scores := Scores {
		X: 1,
		Y: 2,
		Z: 3,
	}

	score := 0

	for {
		line, err := rd.ReadString('\n')
		splitStr := strings.Split(line, " ")

		opponent := splitStr[0]
		player := strings.Split(splitStr[1], "\n")[0]
			
		if opponent == "A" {
			if player == "X" {
				score += scores.Z
			}
			if player == "Y" {
				score += scores.X + outcomes.draw
			}
			if player == "Z" {
				score += scores.Y + outcomes.win
			}
		}

		if opponent == "B" {
			if player == "X" {
				score += scores.X
			}
			if player == "Y" {
				score += scores.Y + outcomes.draw
			}
			if player == "Z" {
				score += scores.Z + outcomes.win
			}
		}

		if opponent == "C" {
			if player == "X" {
				score += scores.Y
			}
			if player == "Y" {
				score += scores.Z + outcomes.draw
			}
			if player == "Z" {
				score += scores.X + outcomes.win
			}
		}
 
		if err == io.EOF {
			fmt.Printf("final score %v\n", score)
			return
		}

		if err != nil && err != io.EOF {
			log.Fatalf("err %v", err)
		}
	}
}
