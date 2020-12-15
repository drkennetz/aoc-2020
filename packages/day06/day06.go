package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data := filereader("../../data/day6/input.txt")
	//fmt.Println(data)
	nums := concatGroups(data)
	fmt.Println("Puzzle 1 answer: ", sumSlice(nums))
	nums2 := compareFirst(data)
	fmt.Println("Puzzle 2 answer: ", sumSlice(nums2))
}

func sumSlice(ii []int) int {
	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum
}

func compareFirst(ss []string) []int {
	ints := make([]int, 0)
	str := ""
	for _, s := range ss {
		if s == "" {
			i := getCount(str)
			ints = append(ints, i)
			str = ""
		} else {
			str += s + "+"
		}
	}
	ints = append(ints, getCount(str))
	return ints
}

func getCount(s string) int {
	num_groups := strings.Count(s, "+")
	total_count := 0
	for i := range s {
		if s[i:i+1] == "+" {
			break
		}
		if strings.Count(s, s[i:i+1]) == num_groups {
			total_count++
		}
	}
	return total_count
}

// puzzle1
func concatGroups(ss []string) []int {
	ints := make([]int, 0)
	str := ""
	for _, s := range ss {
		if s == "" {
			i := getUniqs(str)
			ints = append(ints, i)
			str = ""
		} else {
			str += s
		}
	}
	ints = append(ints, getUniqs(str))
	return ints
}

func getUniqs(s string) int {
	chars := make([]string, 0)
	for i := range s {
		if findStringInSlice(chars, s[i:i+1]) {
			continue
		} else {
			chars = append(chars, s[i:i+1])
		}
	}
	return len(chars)
}

func findStringInSlice(ss []string, s string) bool {
	for _, item := range ss {
		if item == s {
			return true
		}
	}
	return false
}
// Turns a string literal into slice of string
func readData(s string) []string {
	lines := make([]string, 0)
	data := strings.NewReader(s)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return lines
}

func filereader(s string) []string {
	lines := make([]string, 0)

	file, err := os.Open(s)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return lines
}