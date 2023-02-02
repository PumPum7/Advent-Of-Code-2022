package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack []string

func (inputStack *stack) move(toStack *stack, moveBy int) {
	for index := 0; index < moveBy; index++ {
		*toStack = append(*toStack, (*inputStack)[len(*inputStack)-1])
		*inputStack = (*inputStack)[:len(*inputStack)-1]
	}
}

func tops(inputStack []stack) string {
	result := ""
	for _, crate := range inputStack {
		result = result + crate[len(crate)-1]
	}
	return result
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	lines := []string{}
	for sc.Scan() {
		if sc.Text() == "" {
			break
		}
		lines = append(lines, sc.Text())
	}
	stacks := []stack{{}, {}, {}, {}, {}, {}, {}, {}, {}}
	for i := len(lines) - 2; i >= 0; i-- {
		for j, n := 1, 0; j < len(lines[i]); j, n = j+4, n+1 {
			crate := string(lines[i][j])
			if crate != " " {
				stacks[n] = append(stacks[n], crate)
			}
		}
	}
	for sc.Scan() {
		var count, from, to int
		_, err := fmt.Sscanf(sc.Text(), "move %d from %d to %d", &count, &from, &to)
		if err != nil {
			continue
		}
		from -= 1
		to -= 1
		tmp := stack{}
		(&stacks[from]).move(&tmp, count)
		(&tmp).move(&stacks[to], count)
	}
	fmt.Println(tops(stacks))
}
