package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data := filereader("../../data/day3/input.txt")
	treeCountR3D1 := findTreesR3D1(data)
	treeCountR1D1 := findTreesR1D1(data)
	treeCountR5D1 := findTreesR5D1(data)
	treeCountR7D1 := findTreesR7D1(data)
	treeCountR1D2 := findTreesR1D2(data)
	puzzle4answer := treeCountR1D1*treeCountR1D2*treeCountR3D1*treeCountR5D1*treeCountR7D1
	fmt.Println("total trees encountered, R3D1 (puzzle1 answer): ", treeCountR3D1)
	fmt.Println("total trees encountered, R1D1: ", treeCountR1D1)
	fmt.Println("total trees encountered, R5D1: ", treeCountR5D1)
	fmt.Println("total trees encountered, R7D1: ", treeCountR7D1)
	fmt.Println("total trees encountered, R1D2: ", treeCountR1D2)
	fmt.Println("puzzle2 answer: ", puzzle4answer)
}

func findTreesR3D1(s []string) int {
	treeCount := 0
	var currentIndex int
	for i, v := range s {
		if i == 0 {
			currentIndex = 3
		} else {
			if string(v[currentIndex]) == "#" {
				treeCount += 1
			}
			if currentIndex+3 < len(v) {
				currentIndex += 3
			} else {
				currentIndex = currentIndex - (len(v)-3)
			}
		}
	}
	return treeCount
}

func findTreesR1D1(s []string) int {
	treeCount := 0
	var currentIndex int
	for i, v := range s {
		if i == 0 {
			currentIndex = 1
		} else {
			if string(v[currentIndex]) == "#" {
				treeCount += 1
			}
			if currentIndex+1 < len(v) {
				currentIndex += 1
			} else {
				currentIndex = currentIndex - (len(v)-1)
			}
		}
	}
	return treeCount
}

func findTreesR5D1(s []string) int {
	treeCount := 0
	var currentIndex int
	for i, v := range s {
		if i == 0 {
			currentIndex = 5
		} else {
			if string(v[currentIndex]) == "#" {
				treeCount += 1
			}
			if currentIndex+5 < len(v) {
				currentIndex += 5
			} else {
				currentIndex = currentIndex - (len(v)-5)
			}
		}
	}
	return treeCount
}
func findTreesR1D2(s []string) int {
	treeCount := 0
	currentColIndex := 1
	for i := 2; i < len(s); i+=2 {
		fmt.Println("CurrentRowIndex:", i, "CurrentColIndex:", currentColIndex)
		fmt.Println("treeCount:", treeCount)
		if string(s[i][currentColIndex]) == "#" {
			treeCount += 1
		}
		if currentColIndex+1 < len(s[i]) {
			currentColIndex += 1
		} else {
			currentColIndex = currentColIndex - (len(s[i])-1)
		}
	}
	return treeCount
}
func findTreesR7D1(s []string) int {
	treeCount := 0
	var currentIndex int
	for i, v := range s {
		if i == 0 {
			currentIndex = 7
		} else {
			if string(v[currentIndex]) == "#" {
				treeCount += 1
			}
			if currentIndex+7 < len(v) {
				currentIndex += 7
			} else {
				currentIndex = currentIndex - (len(v)-7)
			}
		}
	}
	return treeCount
}

func filereader(f string) []string {
	lines := make([]string, 0)
	file, err := os.Open(f)

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