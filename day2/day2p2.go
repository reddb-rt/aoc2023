package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`([0-9]+|red|green|blue)`)

var colorIndex = map[string]int{"red": 0, "green": 1, "blue": 2}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sumPowers int
	for scanner.Scan() {
		// [red, green, blue]
		var cubeMax [3]int
		parts := re.FindAllString(scanner.Text(), -1)
		for i := 2; i < len(parts); i += 2 {
			val, err := strconv.Atoi(parts[i-1])
			if err != nil {
				log.Fatal(err)
			}
			if val > cubeMax[colorIndex[parts[i]]] {
				cubeMax[colorIndex[parts[i]]] = val
			}
		}
		sumPowers += (cubeMax[0] * cubeMax[1] * cubeMax[2])
	}
	fmt.Println(sumPowers)
}
