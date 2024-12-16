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

//go:embed day13.txt
var inputDay13 string

type day13Machine struct {
	a     vec2
	b     vec2
	prize vec2
}

func readDay13Machines(input string) []day13Machine {
	var result []day13Machine

	var a, b, prize vec2
	r := bufio.NewReader(strings.NewReader(input))

	for {
		bytes, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				result = append(result, day13Machine{a, b, prize})
				break
			}
			panic(err)
		}
		line := string(bytes)

		switch {
		case len(line) == 0:
			result = append(result, day13Machine{a, b, prize})

		case strings.HasPrefix(line, "Button A: "):
			comma := strings.IndexByte(line, ',')
			x, _ := strconv.Atoi(line[12:comma])
			y, _ := strconv.Atoi(line[comma+4:])
			a = vec2{x, y}

		case strings.HasPrefix(line, "Button B: "):
			comma := strings.IndexByte(line, ',')
			x, _ := strconv.Atoi(line[12:comma])
			y, _ := strconv.Atoi(line[comma+4:])
			b = vec2{x, y}

		case strings.HasPrefix(line, "Prize: "):
			comma := strings.IndexByte(line, ',')
			x, _ := strconv.Atoi(line[9:comma])
			y, _ := strconv.Atoi(line[comma+4:])
			prize = vec2{x, y}
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func (d day13Machine) fewestTokens() int {
	ax, ay := float64(d.a.x), float64(d.a.y)
	bx, by := float64(d.b.x), float64(d.b.y)
	px, py := float64(d.prize.x), float64(d.prize.y)

	// Cramer's rule
	det := ax*by - bx*ay
	n := (px*by - bx*py) / det
	m := (ax*py - px*ay) / det

	if math.Floor(m) == m && math.Floor(n) == n {
		return 3*int(n) + int(m)
	} else {
		return 0
	}
}

func (AoC) SolveDay13() {
	ms := readDay13Machines(inputDay13)

	p1 := 0
	for _, m := range ms {
		p1 += m.fewestTokens()
	}
	fmt.Printf("Part 1: %d\n", p1)

	p2 := 0
	for _, m := range ms {
		m.prize = m.prize.add(vec2{10000000000000, 10000000000000})
		p2 += m.fewestTokens()
	}
	fmt.Printf("Part 2: %d\n", p2)
}
