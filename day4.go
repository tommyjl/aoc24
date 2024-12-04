package main

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
)

//go:embed day4.txt
var day4Input string

type Direction int

const (
	E Direction = iota
	SE
	S
	SW
	W
	NW
	N
	NE
)

func (dir Direction) toPos(width int, i int, j int) int {
	switch dir {
	case E:
		return i + j
	case W:
		return i - j
	case S:
		return i + j*width
	case N:
		return i - j*width
	case NE:
		return i - j*width + j
	case NW:
		return i - j*width - j
	case SE:
		return i + j*width + j
	case SW:
		return i + j*width - j
	default:
		panic("invalid direction")
	}
}

func solveDay4Part1() int {
	answer := 0

	width := strings.IndexByte(day4Input, '\n') + 1
	target := `XMAS`
	r := strings.NewReader(day4Input)

	dirs := []Direction{E, SE, S, SW, W, NW, N, NE}
	for _, dir := range dirs {
		for i := 0; i < len(day4Input); i++ {
			found := true
			for j := 0; j < len(target); j++ {
				pos := dir.toPos(width, i, j)

				_, err := r.Seek(int64(pos), io.SeekStart)
				if err != nil {
					found = false
					break
				}

				b, err := r.ReadByte()
				if err != nil {
					found = false
					break
				}
				found = found && b == target[j]
			}
			if found {
				answer++
			}
		}
	}

	return answer
}

func solveDay4Part2() int {
	getByte := func(r *strings.Reader, dir Direction, width, i, j int) byte {
		pos := dir.toPos(width, i, j)
		_, err := r.Seek(int64(pos), io.SeekStart)
		if err != nil {
			return 0
		}
		b, err := r.ReadByte()
		if err != nil {
			return 0
		}
		return b
	}

	answer := 0

	width := strings.IndexByte(day4Input, '\n') + 1
	r := strings.NewReader(day4Input)

	for i := 0; i < len(day4Input); i++ {
		mid := getByte(r, N, width, i, 0)
		nw := getByte(r, NW, width, i, 1)
		ne := getByte(r, NE, width, i, 1)
		sw := getByte(r, SW, width, i, 1)
		se := getByte(r, SE, width, i, 1)

		if nw == 'S' {
			nw, se = se, nw
		}
		if sw == 'S' {
			sw, ne = ne, sw
		}

		if mid == 'A' && nw == 'M' && se == 'S' && sw == 'M' && ne == 'S' {
			answer++
		}
	}

	return answer
}

func (AoC) SolveDay4() {
	fmt.Printf("Part 1: %d\n", solveDay4Part1())
	fmt.Printf("Part 2: %d\n", solveDay4Part2())
}
