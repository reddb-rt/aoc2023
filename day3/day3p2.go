package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reGear = regexp.MustCompile("\\*")
var reDigit = regexp.MustCompile("[0-9]")

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

func partNumIdxs(matrix *[][]string, i int, j int, rowLength int) (int, int, int) {
	engine := *matrix
	l, r := j, j
	for l > 0 {
		lNext := engine[i][l-1]
		if !reDigit.MatchString(lNext) {
			break
		}
		l -= 1
	}
	for r < rowLength-1 {
		rNext := engine[i][r+1]
		if !reDigit.MatchString(rNext) {
			break
		}
		r += 1
	}
	return i, l, r
}

func gearRatio(matrix *[][]string, i int, j int, rowLength int) int {
	engine := *matrix
	partOne, partTwo := findAdjacentParts(&engine, i, j, rowLength)
	return partOne * partTwo
}

func findAdjacentParts(matrix *[][]string, i int, j int, rowLength int) (int, int) {
	engine := *matrix
	var partIdxs []int
	var adjacencies = [8][2]int{
		{i, j - 1},
		{i, j + 1},
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i + 1, j + 1},
		{i + 1, j},
		{i + 1, j - 1},
	}

	for p := 0; p < 8; p++ {
		pi, pj := adjacencies[p][0], adjacencies[p][1]
		if pi < 0 || pi > len(engine) || pj < 0 || pj > rowLength {
			continue
		}
		if reDigit.MatchString(engine[pi][pj]) {
			partRowIdx, partStartIdx, partEndIdx := partNumIdxs(&engine, pi, pj, rowLength)
			if len(partIdxs) > 0 {
				if partRowIdx == partIdxs[len(partIdxs)-3] && partStartIdx == partIdxs[len(partIdxs)-2] && partEndIdx == partIdxs[len(partIdxs)-1] {
					continue
				}
			}
			partIdxs = append(partIdxs, partRowIdx, partStartIdx, partEndIdx)
		}
	}
	if len(partIdxs) != 6 {
		return 0, 0
	}
	partOne, err := strconv.Atoi(strings.Join(engine[partIdxs[0]][partIdxs[1]:partIdxs[2]+1], ""))
	if err != nil {
		log.Fatal("Failed to convert string to int")
	}
	partTwo, err := strconv.Atoi(strings.Join(engine[partIdxs[3]][partIdxs[4]:partIdxs[5]+1], ""))
	if err != nil {
		log.Fatal("Failed to convert string to int")
	}
	return partOne, partTwo
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
	rowLength := len(lines[0])
	engine := make([][]string, rows)
	for i := 0; i < rows; i++ {
		engine[i] = make([]string, rowLength)
		for j := 0; j < rowLength; j++ {
			engine[i][j] = string(lines[i][j])
		}
	}

	var sum int
	for i := 0; i < rows; i++ {
		for j := 0; j < rowLength; j++ {
			if reGear.MatchString(engine[i][j]) {
				ratio := gearRatio(&engine, i, j, rowLength)
				sum += ratio
			}
		}
	}
	fmt.Println(sum)
}
