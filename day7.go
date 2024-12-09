package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

//go:embed day7.txt
var inputDay7 string

func foo(line string, ops []byte) int {
	i := strings.IndexByte(line, ':')

	answer, _ := strconv.Atoi(line[0:i])

	parts := strings.Split(line[i+2:], " ")
	nums := make([]int, len(parts))
	for i, part := range parts {
		nums[i], _ = strconv.Atoi(part)
	}

	for i := 0; i < int(math.Pow(float64(len(ops)), float64(len(nums[1:])))); i++ {
		i2 := i
		total := nums[0]
		for _, num := range nums[1:] {
			op := ops[i2%len(ops)]
			switch op {
			case '|':
				for j := num; j > 0; j /= 10 {
					total *= 10
				}
				total += num
			case '*':
				total *= num
			default:
				total += num
			}
			i2 = i2 / len(ops)
		}
		if total == answer {
			return answer
		}
	}

	return 0
}

func solveDay6(ops []byte) int {
	answer := 0
	r := bufio.NewReader(strings.NewReader(inputDay7))
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		answer += foo(string(line), ops)
	}
	return answer
}

func (AoC) SolveDay7() {
	fmt.Printf("Part 1: %d\n", solveDay6([]byte{'+', '*'}))
	fmt.Printf("Part 2: %d\n", solveDay6([]byte{'+', '*', '|'}))
}
