package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reGame = regexp.MustCompile(`[0-9]+: (.*) \| (.*)$`)

func intInSlice(i int, nums []int) bool {
	for _, num := range nums {
		if i == num {
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

	var score int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := reGame.FindAllStringSubmatch(line, -1)
		var numbers []int
		var winningNumbers []int
		for _, card := range matches {
			if len(card) != 3 {
				log.Fatal("Regexp failure")
			}

			nums := strings.Split(card[1], " ")
			for i := range nums {
				num, err := strconv.Atoi(nums[i])
				if err != nil {
					continue
				}
				numbers = append(numbers, num)
			}

			winningNums := strings.Split(card[2], " ")
			for i := range winningNums {
				winningNum, err := strconv.Atoi(winningNums[i])
				if err != nil {
					continue
				}
				winningNumbers = append(winningNumbers, winningNum)
			}

			var wins int
			for i := range numbers {	
				if intInSlice(numbers[i], winningNumbers) {
					wins += 1
				}
			}
			// Since score is doubled, everything after 2 is equiv to 2^(wins-1)
			if wins > 2 {
				score += int(math.Pow(2, float64(wins - 1)))
			} else {
				score += wins
			}
		}
	}
	fmt.Println(score)
}	
