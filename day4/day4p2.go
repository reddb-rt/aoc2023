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

var reGame = regexp.MustCompile(`([0-9]+): (.*) \| (.*)$`)

type Card struct {
	cardCount int
}

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

	cardMap := make(map[int]*Card)
	// var lastCardNum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		gameMatch := reGame.FindAllStringSubmatch(line, -1)
		var numbers []int
		var winningNumbers []int
		for _, game := range gameMatch {
			if len(game) != 4 {
				log.Fatal("Regexp failure")
			}

			cardNum, err := strconv.Atoi(game[1])
			if err != nil {
				log.Fatal(err)
			}

			// lastCardNum = cardNum

			card, ok := cardMap[cardNum]
			if !ok {
				cardMap[cardNum] = &Card{1}
				card = cardMap[cardNum]
			} else {
				card.cardCount += 1
			}

			nums := strings.Split(game[2], " ")
			for i := range nums {
				num, err := strconv.Atoi(nums[i])
				if err != nil {
					continue
				}
				numbers = append(numbers, num)
			}

			winningNums := strings.Split(game[3], " ")
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

			currentCard := cardMap[cardNum]
			for i := cardNum + 1; i <= cardNum+wins; i++ {
				card, ok := cardMap[i]
				if ok {
					card.cardCount += currentCard.cardCount
				} else {
					cardMap[i] = &Card{currentCard.cardCount}
				}
			}
		}
	}

	var totalCards int
	for i := 1; ; i++ {
		card, ok := cardMap[i]
		if ok {
			totalCards += card.cardCount
		} else {
			break
		}
	}
	fmt.Println(totalCards)
}
