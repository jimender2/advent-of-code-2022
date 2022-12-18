package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(dat)

	totalContains := 0
	for scanner.Scan() {
		tempText := scanner.Text()
		if err != nil {
			panic(err)
		}
		// fmt.Println(tempText)

		values := strings.Split(tempText, ",")

		firstExpand := expandRange(values[0])
		secondExpand := expandRange(values[1])

		if strings.Contains(firstExpand, secondExpand) {
			totalContains = totalContains + 1
		} else if strings.Contains(secondExpand, firstExpand) {
			totalContains = totalContains + 1
		}
	}

	fmt.Println(totalContains)
}

func expandRange(str string) string {

	values := strings.Split(str, "-")

	// convert from string to int
	firstInt, err := strconv.Atoi(values[0])
	if err != nil {
		panic(err)
	}

	secondInt, err := strconv.Atoi(values[1])
	if err != nil {
		panic(err)
	}

	str = ","
	for i := firstInt; i <= secondInt; i++ {
		str = str + strconv.Itoa(i) + ","
	}

	return str
}
