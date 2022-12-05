package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dat, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(dat)

	var listOfElves []int

	listOfElves = append(listOfElves, 0)

	var counter = 0
	var createNewElf = false
	for scanner.Scan() {
		tempText := scanner.Text()
		if tempText == "" {
			counter++
			createNewElf = true
		} else {
			tempInt, err := strconv.Atoi(tempText)

			if err != nil {
				panic(err)
			}

			if createNewElf {
				listOfElves = append(listOfElves, tempInt)
				createNewElf = false
			} else {
				listOfElves[counter] = tempInt + listOfElves[counter]
			}
		}
		// listOfElves = append(listOfElves, tempText)
	}

	var biggest = 0
	var position = 0
	counter = 0
	for i := 0; i < len(listOfElves); i++ {
		if listOfElves[i] > biggest {
			biggest = listOfElves[i]
			position = counter + 1
		}
		counter++
	}
	fmt.Println("biggest: ", biggest)
	fmt.Println("position: ", position)

}
