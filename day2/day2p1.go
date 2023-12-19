package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Max: 12 red cubes, 13 green cubes, and 14 blue cubes
var colorMaxMap = map[string]int{"red": 12, "green": 13, "blue": 14}

var re = regexp.MustCompile(`([0-9]+|red|green|blue)`)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
Lines:
	for scanner.Scan() {
		parts := re.FindAllString(scanner.Text(), -1)
		gameNum, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		for i := 2; i < len(parts); i += 2 {
			val, err := strconv.Atoi(parts[i-1])
			if err != nil {
				log.Fatal(err)
			}
			if val > colorMaxMap[parts[i]] {
				continue Lines
			}
		}
		sum += gameNum
	}
	fmt.Println(sum)
}
