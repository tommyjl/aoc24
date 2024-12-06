package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

//go:embed day5.txt
var inputDay5 string

func sortUpdate(rules map[string][]string, update []string) {
	for rhs := len(update) - 1; rhs >= 0; rhs-- {
		for lhs := 0; lhs <= rhs; lhs++ {
			found := false
			for _, r := range rules[update[lhs]] {
				if slices.Contains(update[0:rhs+1], r) {
					found = true
				}
			}
			if !found {
				update[rhs], update[lhs] = update[lhs], update[rhs]
				break
			}
		}
	}
}

func (AoC) SolveDay5() {
	answer1 := 0
	answer2 := 0
	r := bufio.NewReader(strings.NewReader(inputDay5))
	rules := make(map[string][]string)

	// Rules section
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if len(line) == 0 {
			break
		}

		parts := strings.Split(string(line), "|")
		rules[parts[0]] = append(rules[parts[0]], parts[1])
	}

	// Updates section
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		update := strings.Split(string(line), ",")
		sortUpdate(rules, update)

		num, _ := strconv.Atoi(update[len(update)/2])
		if strings.Join(update, ",") == string(line) {
			answer1 += num
		} else {
			answer2 += num
		}
	}

	fmt.Printf("Part 1: %d\n", answer1)
	fmt.Printf("Part 2: %d\n", answer2)
}
