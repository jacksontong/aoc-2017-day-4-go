package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jacksontong/aoc-2017-day-4-go/security"
)

func main() {
	path := flag.String("file", "input.txt", "Input file")
	flag.Parse()

	lines, err := readInput(*path)
	if err != nil {
		log.Fatal("error reading input", err)
	}

	valid := 0
	for _, line := range lines {
		if ok := security.ValidatePassword(line); ok {
			valid++
		}
	}
	fmt.Printf("Answer: %d\n", valid)
}

// Open the input file, return a 2d slice
// each child slice element is a word
func readInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var out []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return out, scanner.Err()
}
