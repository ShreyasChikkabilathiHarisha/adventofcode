package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

type travel struct {
	time        int
	distance    int
	waysToWin   int
	wayMin      int
	wayMax      int
	timeStr     string
	distanceStr string
}

func part1(f *os.File) {
	fmt.Println("Part 1")
	fmt.Println("======")

	travelGuide := make([]travel, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		// fmt.Println(s)

		row := strings.Split(s, ":")
		i := 0
		if row[0] == "Time" {
			l := strings.Split(row[1], " ")
			for _, v := range l {
				if v == "" {
					continue
				}
				n, _ := strconv.Atoi(v)
				travelGuide = append(travelGuide, travel{time: n})
			}
		} else if row[0] == "Distance" {
			l := strings.Split(row[1], " ")
			for _, v := range l {
				if v == "" {
					continue
				}
				n, _ := strconv.Atoi(v)
				travelGuide[i].distance = n
				travelGuide[i].wayMin = -1
				travelGuide[i].wayMax = -1
				i++
			}
		}
	}

	// fmt.Println(travelGuide)

	for ind, v := range travelGuide {
		speed := 0
		for i := 0; i < v.time; i++ {
			if d := speed * (v.time - i); d > v.distance && v.wayMin == -1 {
				travelGuide[ind].wayMin = i
				// fmt.Println("min:", travelGuide[ind].wayMin)
				break
			}
			speed++
		}
		speed = travelGuide[ind].time
		for i := travelGuide[ind].time; i >= 0; i-- {
			if d := speed * (v.time - i); d > v.distance && v.wayMax == -1 {
				travelGuide[ind].wayMax = i
				// fmt.Println("max:", travelGuide[ind].wayMax)
				break
			}
			speed--
		}
	}

	res := 0
	for _, v := range travelGuide {
		if res == 0 && v.wayMax-v.wayMin+1 > 0 {
			res = v.wayMax - v.wayMin + 1
			continue
		}
		res *= v.wayMax - v.wayMin + 1
	}

	fmt.Println(res)
}

func part2(f *os.File) {
	fmt.Println("Part 2")
	fmt.Println("======")

	travelGuide := make([]travel, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		// fmt.Println(s)

		row := strings.Split(s, ":")
		i := 0
		if row[0] == "Time" {
			l := strings.Split(row[1], " ")
			for _, v := range l {
				if v == "" {
					continue
				}
				travelGuide = append(travelGuide, travel{timeStr: v})
			}
		} else if row[0] == "Distance" {
			l := strings.Split(row[1], " ")
			for _, v := range l {
				if v == "" {
					continue
				}
				travelGuide[i].distanceStr = v
				i++
			}
		}
	}

	// fmt.Println(travelGuide)
	totalTime, totalDistance := "", ""
	for _, v := range travelGuide {
		totalTime += v.timeStr
		totalDistance += v.distanceStr
	}
	totalTimeInt, _ := strconv.Atoi(totalTime)
	totalDistanceInt, _ := strconv.Atoi(totalDistance)
	// fmt.Println(totalTime, totalTimeInt, totalDistance, totalDistanceInt)

	minWay, maxWay := -1, -1
	for i := 0; i < totalTimeInt; i++ {
		if d := i * (totalTimeInt - i); d > totalDistanceInt {
			// fmt.Println("min:", i)
			minWay = i
			break
		}
	}

	for i := totalTimeInt; i >= 0; i-- {
		if d := i * (totalTimeInt - i); d > totalDistanceInt {
			// fmt.Println("max:", i)
			maxWay = i
			break
		}
	}

	fmt.Println(maxWay - minWay + 1)
}
