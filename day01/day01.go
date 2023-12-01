package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v\n", err)
	}
	defer input.Close()

	nums := regexp.MustCompile("[0-9]+")
	answer := 0
	count := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		count++
		lineNums := make([]rune, 0)
		in := scanner.Text()
		for _, c := range in {
			if nums.MatchString(strconv.QuoteRune(c)) {
				lineNums = append(lineNums, c)
			}
		}
		lineInt, _ := strconv.Atoi(fmt.Sprintf("%s%s", string(lineNums[0]), string(lineNums[len(lineNums)-1])))
		answer += lineInt
		fmt.Printf("(%d) Line value: %d, Total so far: %d\n", count, lineInt, answer)
	}

	fmt.Printf("Answer is: %d\n", answer)
}
