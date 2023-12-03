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

	//nums := regexp.MustCompile("[0-9]+")
	nums := regexp.MustCompile("zero|one|two|three|four|five|six|seven|eight|nine|[0-9]+")
	answer := 0
	count := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		count++
		lineNums := make([]string, 0)
		in := scanner.Text()
		for _, c := range in {
			if nums.MatchString(strconv.QuoteRune(c)) {
				lineNums = append(lineNums, c)
			}
		}
		//lineInt, _ := strconv.Atoi(fmt.Sprintf("%s%s", string(lineNums[0]), string(lineNums[len(lineNums)-1])))
		//answer += lineInt
		//fmt.Printf("(%d) Line value: %d, Total so far: %d\n", count, lineInt, answer)
		fmt.Println(lineNums)
	}

	fmt.Printf("Pt. 1 answer is: %d\n", answer)
}

func partTwo(in []rune) int {
	return 0
}
