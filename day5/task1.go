package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack []string

func (s *stack) move(to *stack, n int) {
	for i := 0; i < n; i++ {
		*to = append(*to, (*s)[len(*s)-1])
		*s = (*s)[:len(*s)-1]
	}
}

func tops(ss []stack) string {
	result := ""
	for _, s := range ss {
		result = result + s[len(s)-1]
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
		from = from - 1
		to = to - 1
		(&stacks[from]).move(&stacks[to], count)
	}
	fmt.Println(tops(stacks))
}
