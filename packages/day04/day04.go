package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Passport struct {
	data map[string]string
}

func (p Passport) requiredKey() []string {
	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	sort.Strings(keys)
	return keys
}

func (p Passport) optionalKey() []string {
	keys := []string{"cid"}
	return keys
}

func (p Passport) keys() []string {
	keys := make([]string, 0)
	for k, _ := range p.data {
		if k != "cid" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

func (p Passport) checkKeys() bool {
	required := p.requiredKey()
	contains := p.keys()
	if len(contains) != len(required) {
		return false
	}

	// we know they are the same length and they are sorted
	for i := range (contains) {
		if contains[i] != required[i] {
			return false
		}
	}
	return true
}

func (p Passport) validData() bool {
	for k, v := range p.data {
		if k == "byr" {
			yr, _ := strconv.Atoi(v)
			if yr < int(1920) || yr > int(2002) {
				return false
			}
		} else if k == "iyr" {
			yr, _ := strconv.Atoi(v)
			if yr < int(2010) || yr > int(2020) {
				return false
			}
		} else if k == "eyr" {
			yr, _ := strconv.Atoi(v)
			if yr < int(2020) || yr > int(2030) {
				return false
			}
		} else if k == "hgt" {
			re := regexp.MustCompile(`^((1([5-8][0-9]|9[0-3])cm)|((59|6[0-9]|7[0-6])in))$`)
			if !re.MatchString(v) {
				return false
			}
		} else if k == "hcl" {
			re := regexp.MustCompile(`^#[a-f0-9]{6}$`)
			if !re.MatchString(v) {
				return false
			}
		} else if k == "ecl" {
			re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
			if !re.MatchString(v) {
				return false
			}
		} else if k == "pid" {
			if len(v) != 9 {
				return false
			}
		}
	}
	return true
}

func main() {
	data := filereader("../../data/day4/input.txt")
	passports := parseFile(data)
	valid1 := checkCounts1(passports)
	valid2 := checkCounts2(passports)
	fmt.Println("Number of valid passports from puzzle 1: ", valid1)
	fmt.Println("Number of valid passports from puzzle 2: ", valid2)
}

func checkCounts1(p []Passport) int {
	count := 0
	for _, passport := range p {
		if passport.checkKeys() {
			count++
		}
	}
	return count
}

func checkCounts2(p []Passport) int {
	count := 0
	for _, passport := range p {
		if passport.checkKeys() {
			if passport.validData() {
				count++
			}
		}
	}
	return count
}

func parseLine(s string) map[string]string {
	line := strings.TrimSpace(s)
	pairs := strings.Split(line, " ")
	items := make(map[string]string)
	for _, pair := range pairs {
		key := strings.Split(pair,":")[0]
		val := strings.Split(pair,":")[1]
		items[key] = val
	}
	return items
}

func parseFile(ss []string) []Passport {
	passports := make([]Passport, 0)
	group := ""
	for _, s := range ss {
		if s != "" {
			group = group + " " + s
		} else {
			passports = append(passports, Passport{parseLine(group)})
			group = ""
		}
	}
	// eof doesn't contain the final space so we miss the last passport if we don't append again
	passports = append(passports, Passport{parseLine(group)})
	return passports
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
