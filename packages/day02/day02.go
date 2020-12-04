package main

import (
	"log"
	"os"
)

func main() {
	inputs := filereader("../../data/day2/input.txt")
}



func filereader(f string) []string {
	lines := make([]string, 0)
	file, err := os.Open(f)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

}