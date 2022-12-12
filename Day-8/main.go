package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dat, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		tempText := scanner.Text()
		if err != nil {
			panic(err)
		}
		fmt.Println(tempText)

	}

}
