package main

import (
	_ "embed"
	"fmt"
	"strings"
)

type vec2 struct {
	x int
	y int
}

func (v vec2) add(w vec2) vec2 {
	return vec2{v.x + w.x, v.y + w.y}
}

func (v vec2) sub(w vec2) vec2 {
	return vec2{v.x - w.x, v.y - w.y}
}

//go:embed day8.txt
var inputDay8 string

func (AoC) SolveDay8() {
	width := strings.IndexByte(inputDay8, '\n')
	height := strings.Count(inputDay8, "\n")
	if inputDay8[len(inputDay8)-1] != '\n' {
		height++
	}

	inside := func(antinode vec2) bool {
		return antinode.x >= 0 &&
			antinode.x < width &&
			antinode.y >= 0 &&
			antinode.y < height
	}

	antennas := make(map[byte][]vec2)
	for y := range height {
		for x := range width {
			switch v := inputDay8[y*(width+1)+x]; v {
			case '.':
			case '\n':
			default:
				antennas[v] = append(antennas[v], vec2{x, y})
			}
		}
	}

	antinodes := make(map[vec2]int)
	antinodes2 := make(map[vec2]int)
	for _, as := range antennas {
		for i, a := range as {
			for _, b := range as[i+1:] {
				distance := b.sub(a)
				ns := []vec2{a.sub(distance), b.add(distance)}
				for _, n := range ns {
					antinodes[n]++
				}
				for a2 := a; inside(a2); a2 = a2.sub(distance) {
					antinodes2[a2]++
				}
				for b2 := b; inside(b2); b2 = b2.add(distance) {
					antinodes2[b2]++
				}
			}
		}
	}

	answer := 0
	for antinode := range antinodes {
		if inside(antinode) {
			answer++
		}
	}
	fmt.Printf("Part 1: %d\n", answer)

	answer2 := 0
	for range antinodes2 {
		answer2++
	}
	fmt.Printf("Part 2: %d\n", answer2)
}
