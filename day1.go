package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

//go:embed day1.txt
var inputDay1 string

func solvePart1(lhs, rhs []int) int {
	answer := 0
	for i := range len(lhs) {
		tmp := lhs[i] - rhs[i]
		if tmp < 0 {
			tmp *= -1
		}
		answer += tmp
	}
	return answer
}

func solvePart2(lhs, rhs []int) int {
	m := make(map[int]int)
	for _, n := range rhs {
		m[n]++
	}

	answer := 0
	for _, n := range lhs {
		answer += n * m[n]
	}
	return answer
}

func (AoC) SolveDay1() {
	var lhs []int
	var rhs []int

	r := bufio.NewReader(strings.NewReader(inputDay1))
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		parts := strings.Split(string(line), "   ")

		n1, _ := strconv.Atoi(parts[0])
		n2, _ := strconv.Atoi(parts[1])

		lhs = append(lhs, n1)
		rhs = append(rhs, n2)
	}

	slices.Sort(lhs)
	slices.Sort(rhs)

	fmt.Printf("Part 1: %d\n", solvePart1(lhs, rhs))
	fmt.Printf("Part 2: %d\n", solvePart2(lhs, rhs))
}
