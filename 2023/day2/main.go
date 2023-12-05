package main

import (
	"bufio"
	"fmt"
	"os"
	// "math"
	// "strings"
)

func main() {
	f, err := os.Open("testinput1.txt")
	if err != nil {
		fmt.Println("*** ERROR: ", err)
	}
	defer f.Close()

	part1(f)
	// part2(f)
}

func part1(f *os.File) {
	fmt.Println("Part 1")
	fmt.Println("======")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
	}
}

func part2(f *os.File) {
	fmt.Println("Part 2")
	fmt.Println("======")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
	}
}
