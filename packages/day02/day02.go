package main

import (
	"fmt"
	"regexp"
	"log"
	"strconv"
	"strings"
	f "aoc-2020/packages/filereader"
)

type Pass struct {
	min int
	max int
	letter string
	pass string
}

func main() {
	inputs := f.Filereader("../../data/day2/input.txt")
	total := iterateStringOne(inputs)
	fmt.Println("The number of matching passwords: ", total)
	total2 := iterateStringTwo(inputs)
	fmt.Println("The number of matching passwords for the second case: ", total2)
}

func iterateStringOne(s []string) int{
	count := 0
	for _, v := range s {
		passInfo := parsePassString(v)
		matcher := regexp.MustCompile(passInfo.letter)
		matches := matcher.FindAllStringIndex(passInfo.pass, -1)
		charTotal := len(matches)
		if passInfo.min <= charTotal && charTotal <= passInfo.max {
			count += 1
		}
	}
	return count
}

func iterateStringTwo(s []string) int {
	count := 0
	for _, v := range s {
		passInfo := parsePassString(v)
		strs := strings.Split(passInfo.pass, "")
		// rules said no concept of 0 value
		substring := []string{strs[passInfo.min-1], strs[passInfo.max-1]}
		matches := make([]string, 0)
		for _, s := range substring {
			if passInfo.letter == s {
				matches = append(matches, s)
			}
		}
		if len(matches) == 1 {
			count += 1
		}
	}
	return count
}

func parsePassString(s string) (data Pass) {
	compiled, err := regexp.Compile(`^(\d+)-(\d+) ([a-z]): (\w+)$`)
	if err != nil {
		log.Fatalln(err)
	}
	items := compiled.FindStringSubmatch(s)[:]
	min, _ := strconv.Atoi(items[1])
	max, _ := strconv.Atoi(items[2])

	data = Pass{
			min,
			max,
			items[3],
			items[4],
	}
	return data
}
