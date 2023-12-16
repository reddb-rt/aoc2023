package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var reDigit = regexp.MustCompile("[0-9]")
var reSymbol = regexp.MustCompile("[^\\.0-9]")
var reNum = regexp.MustCompile("[0-9]+")

func getLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var lines []string
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return lines
}

func adjacentSymbol(engine [][]string, i int, j int) bool {
	rows := len(engine)
	cols := len(engine[0])

	// left, i, j-1
	if j-1 >= 0 {
		if reSymbol.MatchString(engine[i][j-1]) {
			return true
		}
	}
	// top left, i-1, j-1
	if i-1 >= 0 && j-1 >= 0 {
		if reSymbol.MatchString(engine[i-1][j-1]) {
			return true
		}
	}
	// top, i-1, j
	if i-1 >= 0 {
		if reSymbol.MatchString(engine[i-1][j]) {
			return true
		}
	}
	// top right, i-1, j+1
	if i-1 >= 0 && j+1 < cols {
		if reSymbol.MatchString(engine[i-1][j+1]) {
			return true
		}
	}
	// right, i, j+1
	if j+1 < cols {
		if reSymbol.MatchString(engine[i][j+1]) {
			return true
		}
	}
	// bottom right, i+1, j+1
	if i+1 < rows && j+1 < cols {
		if reSymbol.MatchString(engine[i+1][j+1]) {
			return true
		}
	}
	// bottom, i+1, j
	if i+1 < rows {
		if reSymbol.MatchString(engine[i+1][j]) {
			return true
		}
	}
	// bottom left, i+1, j-1
	if i+1 < rows && j-1 >= 0 {
		if reSymbol.MatchString(engine[i+1][j-1]) {
			return true
		}
	}
	return false
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := getLines(file)

	rows := len(lines)
	cols := len(lines[0])
	engine := make([][]string, rows)
	for i := 0; i < rows; i++ {
		engine[i] = make([]string, cols)
		for j := 0; j < cols; j++ {
			engine[i][j] = string(lines[i][j])
		}
	}

	var sum int
	var symbolFound bool
	var current string

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !reNum.MatchString(engine[i][j]) {
				if !symbolFound {
					current = ""
					continue
				}
				if symbolFound {
					partNumber, err := strconv.Atoi(current)
					if err != nil {
						log.Fatal(err)
					}
					sum += partNumber
					symbolFound = false
					current = ""
					continue
				}
			} else {
				if symbolFound {
					current += engine[i][j]
					continue
				}
				if adjacentSymbol(engine, i, j) {
					symbolFound = true
					current += engine[i][j]
					continue
				}
				current += engine[i][j]
			}
		}
	}
	fmt.Println(sum)
}
