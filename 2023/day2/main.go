package main

import (
	"bufio"
	"fmt"
	"os"

	// "math"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input2.txt")
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
	breakOutOfGame := false
	for scanner.Scan() {
		s := scanner.Text()

		breakOutOfGame = false
		games := strings.Split(s, ":")
		gameid, _ := strconv.Atoi(strings.TrimSpace(games[0][5:]))
		turns := strings.Split(games[1], ";")
		for _, turn := range turns {
			if breakOutOfGame {
				break
			}

			turn = strings.TrimSpace(turn)
			if turn == "" {
				continue
			}
			// fmt.Println(turn)
			bags := strings.Split(turn, ",")
			for _, bag := range bags {
				if breakOutOfGame {
					break
				}

				bag = strings.TrimSpace(bag)
				if bag == "" {
					continue
				}
				// fmt.Println(bag)

				numcolor := strings.Split(bag, " ")
				num, _ := strconv.Atoi(strings.TrimSpace(numcolor[0]))
				color := strings.TrimSpace(numcolor[1])

				// fmt.Println(num, "*", color)

				if color == "red" && num > 12 {
					breakOutOfGame = true
				}
				if color == "green" && num > 13 {
					breakOutOfGame = true
				}
				if color == "blue" && num > 14 {
					breakOutOfGame = true
				}
			}
		}

		// fmt.Println(s, gameid)
		if !breakOutOfGame {
			// fmt.Println("Summing", gameid)
			sum += gameid
		}
	}

	fmt.Println("Sum:", sum)
}

func part2(f *os.File) {
	fmt.Println("Part 2")
	fmt.Println("======")

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()

		games := strings.Split(s, ":")
		turns := strings.Split(games[1], ";")
		red, green, blue := 0, 0, 0
		for _, turn := range turns {
			turn = strings.TrimSpace(turn)
			if turn == "" {
				continue
			}

			bags := strings.Split(turn, ",")
			for _, bag := range bags {
				bag = strings.TrimSpace(bag)
				if bag == "" {
					continue
				}

				numcolor := strings.Split(bag, " ")
				num, _ := strconv.Atoi(strings.TrimSpace(numcolor[0]))
				color := strings.TrimSpace(numcolor[1])

				if color == "red" && num > red {
					red = num
				}
				if color == "green" && num > green {
					green = num
				}
				if color == "blue" && num > blue {
					blue = num
				}
			}
		}

		// fmt.Println(s, red, green, blue, red*green*blue)
		sum += red * green * blue
	}

	fmt.Println("Sum:", sum)
}
