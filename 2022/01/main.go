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
	/*
	file, err := os.Open("simple.txt")
    if err != nil {
			log.Fatalf("open file error: %v", err)
			return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)

		sum := int64(0)
		highest := int64(0)

    for sc.Scan() {
			stoi, _ := strconv.ParseInt(sc.Text(), 10, 64)
			sum += stoi
			fmt.Printf("Text() %v %T, stoi %v, sum %v\n", sc.Text(), sc.Text(), stoi, sum)

			if sc.Text() == "" { 
				if sum > highest {
					fmt.Printf("prev highest %v\n", highest)
					highest = sum
					// sum = 0
					fmt.Printf("current highest %v\n\n", highest)
				}

				fmt.Printf("sum %v\n", sum)
				sum = 0
				fmt.Printf("\n")
			}
	}

	fmt.Printf("final highest %v\n\n", highest)

	if err := sc.Err(); err != nil {
			log.Fatalf("scan file error: %v", err)
			return
	}
	*/

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
			// line, err := rd.ReadString('\n')
			// val := strings.Split(line, "\n")[0]
			// stoi, _ := strconv.ParseInt(val, 10, 64)
			// sum += stoi
			// fmt.Printf("line %v %T, stoi %v %T, sum %v %T\n", line, line, stoi, stoi, sum, sum)

			// if err != nil {
			// 	fmt.Printf("err != nil check\n")
			// 	if err == io.EOF {
			// 		fmt.Printf("io.EOF check\n")
			// 		if sum > highest {
			// 			fmt.Printf("EOF sum > highest check\n")
			// 			fmt.Printf("EOF prev highest %v\n", highest)
			// 			highest = sum
			// 			// sum = 0
			// 			fmt.Printf("EOF current highest %v\n\n", highest)
			// 		}
			// 		fmt.Printf("final highest %v", highest)
			// 		return
			// 	}
			// }

			// if line == "" {
			// 	fmt.Printf("line == '' check\n")
			// 	if sum > highest {
			// 		fmt.Printf("are we even getting here?\n")
			// 		fmt.Printf("line prev highest %v\n", highest)
			// 		highest = sum
			// 		// sum = 0
			// 		fmt.Printf("line current highest %v\n\n", highest)
			// 	}

			// 	fmt.Printf("line == '' sum %v\n", sum)
			// 	sum = 0
			// }

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