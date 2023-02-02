package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)


	sc.Scan()

	inputText := sc.Text()

	for i := 0; i < len(inputText); i++ {
		set := make(map[string]interface{})
		seq := inputText[i : i+14]
		for _, character := range seq {
			set[string(character)] = ""
		}
		if len(set) == 14 {
			fmt.Println(i + 14)
			break
		}
	}
}
