package filereader

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Filereader reads in AoC input files and returns slice of string
func Filereader(s string) []string {
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