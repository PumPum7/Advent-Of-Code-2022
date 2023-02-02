package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()
	sc := bufio.NewScanner(inputFile)

	score := 0

	for sc.Scan() {
		input := strings.Split(sc.Text(), " ")

		if input[0] == "A" {
			if input[1] == "X" {
				score += 3
			} else if input[1] == "Y" {
				score += 4
			} else {
				score += 8
			}
		} else if input[0] == "B" {
			if input[1] == "X" {
				score += 1
			} else if input[1] == "Y" {
				score += 5
			} else {
				score += 9
			}
		} else {
			if input[1] == "X" {
				score += 2
			} else if input[1] == "Y" {
				score += 6
			} else {
				score += 7
			}
		}
	}

	fmt.Printf("score %d", score)

}
