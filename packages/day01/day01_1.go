package main

import (
	"fmt"
	"strconv"
	f "aoc-2020/packages/filereader"
)

func main() {
	data := f.Filereader("../../data/day1/input.txt")
	intData := toInt(data)
	twentytwentyTwo := find2020Two(intData)
	fmt.Println(twentytwentyTwo)
	fmt.Println(multiplyTwentyTwenty(twentytwentyTwo))
	twentyTwentyThree := find2020Three(intData)
	fmt.Println(twentyTwentyThree)
	fmt.Println(multiplyTwentyTwenty(twentyTwentyThree))
}

func toInt(s []string) []int {
	l := len(s)

	// don't allocate if empty
	if l == 0 {
		return nil
	}

	ints := make([]int, l)

	for i := 0; i < l; i++ {
		f, _ := strconv.Atoi(s[i])

		ints[i] = int(f)
	}

	return ints
}

// find2020 finds the two ints in the slice that sum to 2020
func find2020Two(ints []int) []int {
	vals := make([]int, 2)

	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints); j++ {
			if j == i {
				continue
			} else if ints[i] + ints[j] == 2020 {
				vals[0] = ints[i]
				vals[1] = ints[j]
				return vals
			} else {
				vals[0] = 0
				vals[1] = 0
			}
		}
	}
	return vals
}

func find2020Three(ints []int) []int {
	vals := make([]int, 3)

	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints); j++ {
			for k := 0; k < len(ints); k++ {
				if i == j || i == k || j == k {
					continue
				} else if ints[i] + ints[j] + ints[k] == 2020 {
					vals[0] = ints[i]
					vals[1] = ints[j]
					vals[2] = ints[k]
					return vals
				} else {
					vals[0] = 0; vals[1] = 0; vals[2] = 0
				}
			}
		}
	}
	return vals
}
func multiplyTwentyTwenty(ints []int) int {
	var final int = 1
	for _, v := range ints {
		final *= v
	}
	return final
}