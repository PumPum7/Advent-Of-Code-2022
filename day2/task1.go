package main

import (
	"bufio"
	"fmt"
	"os"
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

	scores := map[string]int{"B X":1, "C Y":2, "A Z":3, "A X":4,"B Y":5,"C Z":6, "C X":7, "A Y":8, "B Z":9}


	for sc.Scan() {
		score += scores[sc.Text()]
	}

	fmt.Printf("score %d", score)

}
