package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input1.txt")
	if err != nil {
		fmt.Println("*** ERROR: ", err)
	}
	defer f.Close()

	// part1(f)
	part2(f)
}

func part1(f *os.File) {
	fmt.Println("Part 1")
	fmt.Println("======")

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		l, r := -1, -1
		for i := 0; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' && l == -1 {
				l = int(s[i] - '0')
			}

			if s[len(s)-1-i] >= '0' && s[len(s)-1-i] <= '9' && r == -1 {
				r = int(s[len(s)-1-i] - '0')
			}

			if l != -1 && r != -1 {
				break
			}
		}
		// fmt.Println(s)
		// fmt.Println(l, r)

		sum += l*10 + r
	}

	fmt.Println("Sum: ", sum)
}

func part2(f *os.File) {
	fmt.Println("Part 2")
	fmt.Println("======")

	scanner := bufio.NewScanner(f)

	sum := 0
	numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for scanner.Scan() {
		s := scanner.Text()
		li, l, ri, r := math.MaxInt32, 0, math.MinInt32, 0

		for j := 0; j < len(numbers); j++ {
			if t := strings.Index(s, numbers[j]); t != -1 && t < li {
				li = t
				l = j
			}

			if t := strings.LastIndex(s, numbers[j]); t != -1 && t > ri {
				ri = t
				r = j
			}
		}

		for i := 0; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' && i < li {
				li = i
				l = int(s[i] - '0')
			}

			if s[len(s)-1-i] >= '0' && s[len(s)-1-i] <= '9' && (len(s)-1-i) > ri {
				ri = len(s) - 1 - i
				r = int(s[len(s)-1-i] - '0')
			}
		}
		// fmt.Println(s)
		// fmt.Println(l, r)

		sum += l*10 + r
	}

	fmt.Println("Sum: ", sum)
}
