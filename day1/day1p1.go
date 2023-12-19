package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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
	l, r := 0, len(line)-1
	var l_val, r_val rune
	for {
		if unicode.IsDigit(rune(line[l])) {
			l_val = rune(line[l])
			break
		}
		l++
	}
	for {
		if unicode.IsDigit(rune(line[r])) {
			r_val = rune(line[r])
			break
		}
		r--
	}
	result, _ := strconv.Atoi(string(l_val) + string(r_val))
	return result
}
