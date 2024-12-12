package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed day10.txt
var inputDay10 string

type day10 struct {
	width  int
	height int
	m      []int
}

func newDay10(input string) day10 {
	width := strings.IndexByte(input, '\n')
	height := strings.Count(input, "\n")
	if input[len(input)-1] != '\n' {
		height++
	}

	m := make([]int, width*height)

	i := 0
	for _, ch := range input {
		if ch != '\n' {
			if ch == '.' {
				m[i] = -1
			} else {
				height, err := strconv.Atoi(string(ch))
				if err != nil {
					panic(err)
				}
				m[i] = height
			}
			i++
		}
	}

	return day10{
		width:  width,
		height: height,
		m:      m,
	}
}

func (d day10) getTrailheads() []vec2 {
	var trailheads []vec2
	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {
			if d.m[y*d.width+x] == 0 {
				trailheads = append(trailheads, vec2{x, y})
			}
			if y != 0 && y != d.height-1 && x == 0 {
				// TODO: Delete this if irrelevant in part 2
				//x += d.width - 2
			}
		}
	}
	return trailheads
}

func (d day10) walk(pos vec2) [][]vec2 {
	var paths [][]vec2
	path := []vec2{pos}
	result := d.walkInner(paths, path, 0)
	return result

}

func (d day10) walkInner(result [][]vec2, path []vec2, val int) [][]vec2 {
	pos := path[len(path)-1]
	dirs := []Direction{N, E, S, W}

	for _, dir := range dirs {
		nextPos := d.step(pos, dir)
		if d.outofbounds(nextPos) {
			continue
		}

		nextVal := d.get(nextPos)
		if nextVal == val+1 {
			nextPath := append(slices.Clone(path), nextPos)
			result = d.walkInner(result, nextPath, nextVal)

			if nextVal == 9 {
				result = append(result, nextPath)
			}
		}
	}

	return result
}

func (d day10) outofbounds(pos vec2) bool {
	return pos.x < 0 || pos.x >= d.width ||
		pos.y < 0 || pos.y >= d.height
}

func (d day10) step(pos vec2, dir Direction) vec2 {
	switch dir {
	case N:
		return pos.sub(vec2{0, 1})
	case S:
		return pos.add(vec2{0, 1})
	case E:
		return pos.add(vec2{1, 0})
	case W:
		return pos.sub(vec2{1, 0})
	default:
		panic("unexpected direction")
	}
}

func (d day10) get(pos vec2) int {
	return d.m[pos.y*d.width+pos.x]
}

func (d day10) solve() int {
	total := 0
	for _, trailhead := range d.getTrailheads() {
		paths := d.walk(trailhead)

		ends := make(map[vec2]int)
		for _, r := range paths {
			end := r[len(r)-1]
			ends[end]++
		}

		total += len(ends)
	}
	return total
}

func (d day10) solve2() int {
	total := 0
	for _, trailhead := range d.getTrailheads() {
		paths := d.walk(trailhead)
		total += len(paths)
	}
	return total
}

func (AoC) SolveDay10() {
	fmt.Printf("Part 1: %d\n", newDay10(inputDay10).solve())
	fmt.Printf("Part 2: %d\n", newDay10(inputDay10).solve2())
}
