package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var number_map = map[string]string{
	"one": "1",
	"two": "2",
	"three": "3",
	"four": "4",
	"five": "5",
	"six": "6",
	"seven": "7",
	"eight": "8",
	"nine": "9",
	"1": "1",
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
}

var number_match = regexp.MustCompile(
	`[0-9]|one|two|three|four|five|six|seven|eight|nine`,
)


func main() {
	var sum int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := getValue(scanner.Text())
		sum += val
	}
	fmt.Println(sum)
}

func getValue(line string) int {
	r := len(line) - 1
	match := number_match.FindString(line)
	l_val := number_map[match]
	var r_val string
	for {
		window := line[r:]
		match := number_match.FindString(window)
		if match != "" {
			r_val = number_map[match]
			break
		}
		r--
	}
	result, _ := strconv.Atoi(l_val + r_val)
	return result
}
