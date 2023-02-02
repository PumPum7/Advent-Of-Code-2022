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

	maxCalories1, maxCalories2, maxCalories3 := 0, 0, 0
	currentElf := 0

	for sc.Scan() {
		calories, err := strconv.Atoi(sc.Text())

		// if an error happens it means that the line is empty -> new elf
		if err != nil {
			if currentElf > maxCalories1 {
				maxCalories1 = currentElf
			} else if currentElf > maxCalories2 {
				maxCalories2 = currentElf
			} else if currentElf > maxCalories3 {
				maxCalories3 = currentElf
			}
			currentElf = 0
		} else {
			currentElf += calories
		}
	}

	fmt.Printf("max calories %d", maxCalories1+maxCalories2+maxCalories3)

}
