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

//go:embed day2.txt
var inputDay2 string

func safeReport(report []int) bool {
	safe := true

	dir := 1
	if report[0] > report[1] {
		dir = -1
	}

	prev := report[0]
	for _, cur := range report[1:] {
		diff := dir * (cur - prev)
		if diff < 1 || diff > 3 {
			safe = false
			break
		}
		prev = cur
	}

	return safe
}

func day2Part1(reports [][]int) int {
	tot := 0

	for _, report := range reports {
		if safeReport(report) {
			tot++
		}
	}

	return tot
}

func day2Part2(reports [][]int) int {
	tot := 0

	for _, report := range reports {
		if safeReport(report) {
			tot++
			continue
		}
		for i := 0; i < len(report); i++ {
			report2 := slices.Clone(report)
			report2 = slices.Delete(report2, i, i+1)
			if safeReport(report2) {
				tot++
				break
			}
		}
	}

	return tot
}

func (AoC) SolveDay2() {
	var reports [][]int
	r := bufio.NewReader(strings.NewReader(inputDay2))
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}

		parts := strings.Split(string(line), " ")

		report := make([]int, 0, len(parts))
		for _, part := range parts {
			level, _ := strconv.Atoi(part)
			report = append(report, level)
		}

		reports = append(reports, report)
	}

	fmt.Printf("Part 1 = %d\n", day2Part1(reports))
	fmt.Printf("Part 2 = %d\n", day2Part2(reports))
}
