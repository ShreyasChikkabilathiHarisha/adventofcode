package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("*** ERROR: ", err)
	}

	// part1(string(f))
	part2(string(f))
}

func part1(f string) {
	fmt.Println("Part 1")
	fmt.Println("======")

	maps := strings.Split(f, "\n\n")
	seedPattern := []string{}
	wholeMap := map[string][]mapping{}
	for i, m := range maps {
		// fmt.Println("=====")
		// fmt.Println(m)

		if i == 0 {
			seedPattern = strings.Split(strings.Split(m, ": ")[1], " ")
			continue
		}

		details := strings.Split(m, "\n")
		// fmt.Println("det: ", len(details), details)
		srcdst := strings.Split(details[0], "-")
		src := strings.TrimSpace(srcdst[0])
		dst := strings.TrimSpace(strings.Split(srcdst[2], " ")[0])
		// fmt.Println("s and d", src, dst)

		for _, detail := range details[1:] {
			det := strings.Split(detail, " ")
			srcRangeStart, _ := strconv.Atoi(det[1])
			dstRangeStart, _ := strconv.Atoi(det[0])
			_range, _ := strconv.Atoi(det[2])
			if _, ok := wholeMap[src+dst]; !ok {
				wholeMap[src+dst] = []mapping{}
			}
			wholeMap[src+dst] = append(wholeMap[src+dst], mapping{srcRangeStart, dstRangeStart, _range})
		}
	}

	res := math.MaxInt64
	for _, s := range seedPattern {
		src, _ := strconv.Atoi(s)
		soil := getMappedValue(src, wholeMap["seedsoil"])
		fertilizer := getMappedValue(soil, wholeMap["soilfertilizer"])
		water := getMappedValue(fertilizer, wholeMap["fertilizerwater"])
		light := getMappedValue(water, wholeMap["waterlight"])
		temperature := getMappedValue(light, wholeMap["lighttemperature"])
		humidity := getMappedValue(temperature, wholeMap["temperaturehumidity"])
		location := getMappedValue(humidity, wholeMap["humiditylocation"])
		// fmt.Println(s, soil, fertilizer, water, light, temperature, humidity, location)

		if location < res {
			res = location
		}
	}

	fmt.Println(seedPattern, res)
}

type mapping struct {
	src    int
	dst    int
	rang_e int
}

func getMappedValue(src int, srcToDst []mapping) int {
	// fmt.Println("getMappedValue", src, srcToDst)
	for _, m := range srcToDst {
		if src >= m.src && src <= m.src+m.rang_e {
			// fmt.Println("returning", m.dst+src-m.src)
			return m.dst + src - m.src
		}
	}

	return src
}

func part2(f string) {
	fmt.Println("Part 2")
	fmt.Println("======")

	maps := strings.Split(f, "\n\n")
	seedPattern := []string{}
	wholeMap := map[string][]mapping{}
	for i, m := range maps {
		// fmt.Println("=====")
		// fmt.Println(m)

		if i == 0 {
			seedPattern = strings.Split(strings.Split(m, ": ")[1], " ")
			continue
		}

		details := strings.Split(m, "\n")
		// fmt.Println("det: ", len(details), details)
		srcdst := strings.Split(details[0], "-")
		src := strings.TrimSpace(srcdst[0])
		dst := strings.TrimSpace(strings.Split(srcdst[2], " ")[0])
		// fmt.Println("s and d", src, dst)

		for _, detail := range details[1:] {
			det := strings.Split(detail, " ")
			srcRangeStart, _ := strconv.Atoi(det[1])
			dstRangeStart, _ := strconv.Atoi(det[0])
			_range, _ := strconv.Atoi(det[2])
			if _, ok := wholeMap[src+dst]; !ok {
				wholeMap[src+dst] = []mapping{}
			}
			wholeMap[src+dst] = append(wholeMap[src+dst], mapping{srcRangeStart, dstRangeStart, _range})
		}
	}

	res := math.MaxInt64
	history := map[int]int{}
	locChan := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < len(seedPattern)-1; i += 2 {
		srcStart, _ := strconv.Atoi(seedPattern[i])
		range_, _ := strconv.Atoi(seedPattern[i+1])
		// fmt.Println("***** srcStart", srcStart, "range_", range_)

		// split the range into 1000000 parts and calculate the location for each part in a goroutine
		for range_ > 0 {
			r := 1000000
			if range_ < r {
				r = range_
				range_ = 0
			} else {
				// fmt.Println("srcStart", srcStart, "r", r)
				range_ -= r
			}

			src, ran := srcStart, r

			wg.Add(1)
			go func(srcStart, r int, wholeMap map[string][]mapping, history map[int]int, wg *sync.WaitGroup) {
				defer wg.Done()
				loc := getSmallestLocationForRange(srcStart, r, wholeMap, history)
				// fmt.Println("srcStart", srcStart, "r", r, "loc", loc)
				locChan <- loc
			}(src, ran, wholeMap, history, &wg)

			srcStart += r
		}
	}

	go func() {
		defer close(locChan)
		wg.Wait()
	}()

	for loc := range locChan {
		// fmt.Println("loc", loc)
		if loc < res {
			res = loc
		}
	}

	fmt.Println("***** results ***** :", res, " The pattern was: ", seedPattern)
}

func getSmallestLocationForRange(srcStart, range_ int, wholeMap map[string][]mapping, history map[int]int) int {
	res := math.MaxInt64
	for src := srcStart; src <= srcStart+range_; src++ {
		// if we wanna use the history (memoization) to speed up the process, we need to make that map write
		// thread safe and then uncomment the following lines and the write to history map below

		// if the location of this seed is already calculated, skip since the location will be the
		// same and won't be any less than the previous one
		// if _, ok := history[src]; ok {
		// 	continue
		// }
		soil := getMappedValue(src, wholeMap["seedsoil"])
		fertilizer := getMappedValue(soil, wholeMap["soilfertilizer"])
		water := getMappedValue(fertilizer, wholeMap["fertilizerwater"])
		light := getMappedValue(water, wholeMap["waterlight"])
		temperature := getMappedValue(light, wholeMap["lighttemperature"])
		humidity := getMappedValue(temperature, wholeMap["temperaturehumidity"])
		location := getMappedValue(humidity, wholeMap["humiditylocation"])
		// history[src] = location
		// fmt.Println(s, soil, fertilizer, water, light, temperature, humidity, location)

		if location < res {
			res = location
		}
	}

	return res
}
