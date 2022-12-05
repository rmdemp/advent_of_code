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
				score += scores.X + outcomes.draw
				fmt.Printf("AX %v\n", score)
			}
			if player == "Y" {
				score += scores.X
				fmt.Printf("AY %v\n", score)
			}
			if player == "Z" {
				score += scores.X + outcomes.win
				fmt.Printf("AZ %v\n", score)
			}
		}

		if opponent == "B" {
			if player == "X" {
				score += scores.Y + outcomes.win
				fmt.Printf("BX %v\n", score)
			}
			if player == "Y" {
				score += scores.Y + outcomes.draw
				fmt.Printf("BY %v\n", score)
			}
			if player == "Z" {
				score += scores.Y
				fmt.Printf("BZ %v\n", score)
			}
		}

		if opponent == "C" {
			if player == "X" {
				score += scores.Z
				fmt.Printf("CX %v\n", score)
			}
			if player == "Y" {
				score += scores.Z + outcomes.win
				fmt.Printf("CY %v\n", score)
			}
			if player == "Z" {
				score += scores.Z + outcomes.draw
				fmt.Printf("CZ %v\n", score)
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

/*

[A, X] = 1 rock
[B, Y] = 2 paper
[C, Z] = 3 scissors

lose = 0
draw = 3
win = 6

[A, X] > [C, Z] rock beats scissors
[B, Y] > [A, X] paper beats rock
[C, Z] > [B, Y] scissors beats paper

for each line

	split string
		string[0] = opponent
		string[1] = player (trim \n)

	if string[0] == A
		if string[2] == X
			A + draw
		if string[2] == Y
			A + lose
		if string[2] == Z
			A + win

	if string[0] == B
		if string[2] == X
			B + win
		if string[2] == Y
			B + draw
		if string[2] == Z
		 B + lose

	if string[0] == C
		if string[2] == X
		 C + lose
		if string[2] == Y
	   C + win
		if string[2] == Z
			C + draw
 
*/