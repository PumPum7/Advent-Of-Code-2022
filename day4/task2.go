package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var numberOfContained int

	for sc.Scan(){
		var startFirst, endFirst, startSecond, endSecond int
		fmt.Sscanf(sc.Text(), "%d-%d,%d-%d", &startFirst, &endFirst, &startSecond, &endSecond)
		if startSecond <= endFirst && endSecond >= startFirst || startFirst <= endSecond && endFirst >= startSecond{
			numberOfContained += 1
		}
	}
	fmt.Println(numberOfContained)
}