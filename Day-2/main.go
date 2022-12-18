package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	possibleOutcomes := make(map[string]int)

	possibleOutcomes["win"] = 6
	possibleOutcomes["draw"] = 3
	possibleOutcomes["lose"] = 0

	dat, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(dat)

	points := 0

	for scanner.Scan() {
		tempText := scanner.Text()
		if err != nil {
			panic(err)
		}
		fmt.Println(tempText)

		inputs := strings.Split(tempText, " ")
		points = points + determineOutcome(inputs[0], inputs[1])
	}

	fmt.Println(points)
}

func determineOutcome(player1 string, player2 string) int {
	possibleOutcomes := make(map[string]int)

	possibleOutcomes["win"] = 6
	possibleOutcomes["draw"] = 3
	possibleOutcomes["lose"] = 0

	typePoints := make(map[string]int)

	typePoints["X"] = 1
	typePoints["Y"] = 2
	typePoints["Z"] = 3

	// 1 for Rock, 2 for Paper, and 3 for Scissors
	// X for Rock, Y for Paper, and Z for Scissors

	points := typePoints[player2]

	// A for Rock, B for Paper, and C for Scissors
	// X for Rock, Y for Paper, and Z for Scissors

	//rock
	if player1 == "A" {
		//rock
		if player2 == "X" {
			points = points + possibleOutcomes["draw"]
		} else if player2 == "Y" {
			points = points + possibleOutcomes["win"]
		} else if player2 == "Z" {
			points = points + possibleOutcomes["lose"]
		}
	} else if player1 == "B" {
		//rock
		if player2 == "X" {
			points = points + possibleOutcomes["lose"]
		} else if player2 == "Y" {
			points = points + possibleOutcomes["draw"]
		} else if player2 == "Z" {
			points = points + possibleOutcomes["win"]
		}
	} else if player1 == "C" {
		//rock
		if player2 == "X" {
			points = points + possibleOutcomes["win"]
		} else if player2 == "Y" {
			points = points + possibleOutcomes["lose"]
		} else if player2 == "Z" {
			points = points + possibleOutcomes["draw"]
		}
	}

	return points
}
