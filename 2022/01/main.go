package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

	sum := int(0)
	totals := []int{}

	for {
		line, err := rd.ReadString('\n')

		parsedStr := strings.Split(line, "\n")[0]
		stoi, _ := strconv.ParseInt(parsedStr, 10, 32)
		sum += int(stoi)

		if parsedStr == "" {
			totals = append(totals, sum)
			sum = 0
		}

		if err == io.EOF {
			// https://www.codekru.com/go/how-to-sort-an-array-in-golang#sorting_an_int_array_in_descending_order
			totals = append(totals, sum)
			arrSlice := totals[:]
			sort.Sort(sort.Reverse(sort.IntSlice(arrSlice)))
			fmt.Printf("top three total %v\n", totals[0] + totals[1] + totals[2])
			return
		}

		if err != nil && err != io.EOF {
			log.Fatalf("err %v", err)
		}
	}
}