package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()
	sc := bufio.NewScanner(inputFile)

	maxCaloriesSum := 0
	currentElf := 0

	for sc.Scan() {
		calories, err := strconv.Atoi(sc.Text())

		// if an error happens it means that the line is empty -> new elf
		if err != nil {
			if currentElf > maxCaloriesSum {
				maxCaloriesSum = currentElf
			}
			currentElf = 0
		} else {
			currentElf += calories
		}
	}

	fmt.Printf("max calories %d", maxCaloriesSum)


}