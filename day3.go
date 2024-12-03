package main

import (
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed day3.txt
var inputDay3 string

type day3Scanner struct {
	pattern string
	input   string
	pos     int
	do      bool
}

func newDay3scanner(input string, do bool) *day3Scanner {
	return &day3Scanner{
		pattern: `mul(0,0)`,
		input:   input,
		do:      do,
	}
}

func (s *day3Scanner) backup() {
	s.pos--
	if s.pos < 0 {
		panic("invalid backup")
	}
}

func (s *day3Scanner) next() byte {
	if s.pos < len(s.input) {
		result := s.input[s.pos]
		s.pos++
		return result
	} else {
		return 0
	}
}

func (s *day3Scanner) peek() byte {
	result := s.next()
	s.backup()
	return result
}

func (s *day3Scanner) match() (string, string, bool) {
	enabled := true

	for start := s.pos; start < len(s.input); start++ {
		var tmp []string
		s.pos = start
		found := true

		if s.do {
			a := min(s.pos, len(s.input))

			dont := `don't()`
			b := min(s.pos+len(dont), len(s.input))
			if s.input[a:b] == dont {
				enabled = false
			}

			do := `do()`
			b = min(s.pos+len(do), len(s.input))
			if s.input[a:b] == do {
				enabled = true
			}
		}

		if s.do && !enabled {
			continue
		}

		for _, target := range s.pattern {
			if target == '0' {
				start := s.pos
				n := s.peek()
				for n >= '0' && n <= '9' {
					s.next()
					n = s.peek()
				}
				tmp = append(tmp, s.input[start:s.pos])
			} else {
				n := s.next()
				if rune(n) != target {
					found = false
					break
				}
			}
		}
		if found {
			return tmp[0], tmp[1], true
		}
	}

	return "", "", false
}

func solveDay3Part1() int {
	answer := 0
	s := newDay3scanner(inputDay3, false)

	for lhs, rhs, ok := s.match(); ok; lhs, rhs, ok = s.match() {
		l, _ := strconv.Atoi(lhs)
		r, _ := strconv.Atoi(rhs)
		answer += l * r
	}

	return answer
}

func solveDay3Part2() int {
	answer := 0
	s := newDay3scanner(inputDay3, true)

	for lhs, rhs, ok := s.match(); ok; lhs, rhs, ok = s.match() {
		l, _ := strconv.Atoi(lhs)
		r, _ := strconv.Atoi(rhs)
		answer += l * r
	}

	return answer
}

func (AoC) SolveDay3() {
	fmt.Printf("Part 1: %d\n", solveDay3Part1())
	fmt.Printf("Part 2: %d\n", solveDay3Part2())
}
