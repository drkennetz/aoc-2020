package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	f "aoc-2020/packages/filereader"
)

var encoding = map[string]string {
	"F" : "0",
	"B" : "1",
	"L" : "0",
	"R" : "1",
}

func main() {
	data := f.Filereader("../../data/day5/input.txt")
	ints := convertAllSlice(data)
	fmt.Println("Max seat ID: ", findMax(ints))
	fmt.Println("Missing seat is: ", getEmptySeat(ints))
}

func convertAllSlice(ss []string) []int {
	ints := make([]int, 0)
	for _, v := range ss {
		encode := encodeSeat(v)
		i := binaryToInt(encode)
		ints = append(ints, int(i))
	}
	return ints
}

func getEmptySeat(ii []int) int {
	var missing int
	sort.Ints(ii)
	for i := 0; i+1 < len(ii); i++ {
		if ii[i] != (ii[i+1] - 1) {
			missing = (ii[i] + 1)
		}
	}
	return missing

}
func encodeSeat(s string) string {

	encoded := ""
	for i := 0; i < len(s); i++ {
		encoded += encoding[s[i:i+1]]
	}
	return encoded
}

func binaryToInt(s string) int64 {
	num, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return num
}

func findMax(ii []int) int {
	var max int = ii[0]
	for _, v := range ii {
		if max < v {
			max = v
		}
	}
	return max
}
