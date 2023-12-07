package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type card struct {
	hand    string
	pattern patternType
	bid     int
	rank    int
}

type patternType int

const (
	fiveOfAKind patternType = iota
	fourOfAKind
	fullHouse
	threeOfAKind
	twoPair
	onePair
	highCard
)

func main() {
	f, err := os.Open("input.txt")
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
	cards := make([]card, 0)
	totalHands := 0
	for scanner.Scan() {
		s := scanner.Text()
		details := strings.Split(s, " ")
		bid, _ := strconv.Atoi(details[1])
		cards = append(cards, card{hand: details[0], pattern: getPattern1(details[0]), bid: bid})
		totalHands++
	}

	sort.Slice(cards, func(i, j int) bool {
		return compareFunc1(cards[i].hand, cards[j].hand)
	})

	res := 0
	for i, v := range cards {
		res += v.bid * (len(cards) - i)
		// fmt.Println(v.hand, v.bid, len(cards)-i)
	}

	fmt.Println(res)
}

func part2(f *os.File) {
	fmt.Println("Part 2")
	fmt.Println("======")

	scanner := bufio.NewScanner(f)
	cards := make([]card, 0)
	totalHands := 0
	for scanner.Scan() {
		s := scanner.Text()
		details := strings.Split(s, " ")
		bid, _ := strconv.Atoi(details[1])
		cards = append(cards, card{hand: details[0], pattern: getPattern2(details[0]), bid: bid})
		totalHands++
	}

	sort.Slice(cards, func(i, j int) bool {
		return compareFunc2(cards[i].hand, cards[j].hand)
	})

	res := int64(0)
	for i, v := range cards {
		res += int64(v.bid * (len(cards) - i))
		// fmt.Println(v.hand, v.bid, v.pattern.getInt(), len(cards)-i)
	}

	fmt.Println(res)
}

func getPattern1(hand string) patternType {
	m := map[byte]int{}
	for _, c := range hand {
		m[byte(c)]++
	}

	switch len(m) {
	case 1:
		return fiveOfAKind
	case 2:
		for _, v := range m {
			if v == 4 {
				return fourOfAKind
			}
		}
		return fullHouse
	case 3:
		for _, v := range m {
			if v == 3 {
				return threeOfAKind
			}
		}
		return twoPair
	case 4:
		return onePair
	}

	return highCard
}

func (p patternType) getInt() int {
	switch p {
	case fiveOfAKind:
		return 7
	case fourOfAKind:
		return 6
	case fullHouse:
		return 5
	case threeOfAKind:
		return 4
	case twoPair:
		return 3
	case onePair:
		return 2
	}

	return 1
}

func compareFunc1(s1, s2 string) bool {
	if getPattern1(s1).getInt() > getPattern1(s2).getInt() {
		return true
	} else if getPattern1(s1).getInt() < getPattern1(s2).getInt() {
		return false
	}

	cardRank := []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(cardRank); j++ {
			if s1[i] == cardRank[j] && s2[i] != cardRank[j] {
				return true
			}
			if s2[i] == cardRank[j] && s1[i] != cardRank[j] {
				return false
			}
		}
	}

	return true
}

func getPattern2(hand string) patternType {
	m := map[byte]int{}
	for _, c := range hand {
		m[byte(c)]++
	}

	switch len(m) {
	case 1:
		return fiveOfAKind
	case 2:
		if v, ok := m['J']; ok && (v == 1 || v == 2 || v == 3 || v == 4) {
			return fiveOfAKind
		}
		for _, v := range m {
			if v == 4 {
				return fourOfAKind
			}
		}
		return fullHouse
	case 3:
		j := 0
		if v, ok := m['J']; ok {
			j = v
		}
		if j == 3 || j == 2 {
			return fourOfAKind
		}
		for k, v := range m {
			if v == 3 && j == 2 && k != 'J' {
				return fiveOfAKind
			}
			if v == 3 && j == 1 && k != 'J' {
				return fourOfAKind
			}
			if v == 3 {
				return threeOfAKind
			}
		}
		if j == 1 {
			return fullHouse
		}
		return twoPair
	case 4:
		j := 0
		if v, ok := m['J']; ok {
			j = v
		}
		if j == 2 || j == 1 {
			return threeOfAKind
		}
		return onePair
	}

	if v, ok := m['J']; ok && v == 1 {
		return onePair
	}

	return highCard
}

func compareFunc2(s1, s2 string) bool {
	if getPattern2(s1).getInt() > getPattern2(s2).getInt() {
		return true
	} else if getPattern2(s1).getInt() < getPattern2(s2).getInt() {
		return false
	}

	cardRank := []byte{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(cardRank); j++ {
			if s1[i] == cardRank[j] && s2[i] != cardRank[j] {
				return true
			}
			if s2[i] == cardRank[j] && s1[i] != cardRank[j] {
				return false
			}
		}
	}

	return true
}
