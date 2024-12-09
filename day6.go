package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed day6.txt
var inputDay6 string

type day6Map struct {
	// map
	m      []bool
	width  int
	height int

	// guard
	gx   int
	gy   int
	gdir Direction
}

func newDay6Map() day6Map {
	var result day6Map

	result.width = strings.IndexByte(inputDay6, '\n')
	result.height = strings.Count(inputDay6, "\n")
	if inputDay6[len(inputDay6)-1] != '\n' {
		result.height++
	}

	size := result.width * result.height
	result.m = make([]bool, size, size)

	for y := 0; y < result.height; y++ {
		for x := 0; x < result.width; x++ {
			switch tile := inputDay6[y*(result.height+1)+x]; tile {
			case '^':
				result.gx, result.gy = x, y
				result.gdir = N
			case '#':
				result.m[y*result.height+x] = true
			case '.':
				result.m[y*result.height+x] = false
			default:
				panic(fmt.Sprintf("unexpected character: %c", tile))
			}
		}
	}

	return result
}

func (m day6Map) clone() day6Map {
	next := m
	next.m = slices.Clone(m.m)
	return next
}

func (m day6Map) nextPos() (int, int) {
	x, y := m.gx, m.gy
	switch m.gdir {
	case N:
		y--
	case E:
		x++
	case S:
		y++
	case W:
		x--
	default:
		panic("unexpected guard direction")
	}
	return x, y
}

func (m *day6Map) turn() {
	switch m.gdir {
	case N:
		m.gdir = E
	case E:
		m.gdir = S
	case S:
		m.gdir = W
	case W:
		m.gdir = N
	default:
		panic("unexpected guard direction")
	}
}

func (m *day6Map) walk(fn func(int, int, Direction) bool) {
	for {
		if !fn(m.gx, m.gy, m.gdir) {
			break
		}

		nextx, nexty := m.nextPos()
		i := nexty*m.height + nextx
		if !(nextx >= 0 && nextx < m.width && nexty >= 0 && nexty < m.height) {
			m.gx, m.gy = nextx, nexty
			break
		}

		if m.m[i] {
			m.turn()
		} else {
			m.gx, m.gy = nextx, nexty
		}
	}
}

func solveDay6Part1() int {
	answer := 0
	m := newDay6Map()
	v := make([]bool, m.width*m.height, m.width*m.height)
	m.walk(func(gx, gy int, gdir Direction) bool {
		i := gy*m.height + gx
		if !v[i] {
			answer++
		}
		v[i] = true
		return true
	})
	return answer
}

func solveDay6Part2() int {
	mstart := newDay6Map()

	answer := 0
	for y := range mstart.height {
		for x := range mstart.width {
			m := mstart.clone()
			if x == m.gx && y == m.gy {
				continue
			}
			if i := y*m.height + x; m.m[i] {
				continue
			} else {
				m.m[i] = true
			}

			history := make(map[int]bool)
			olddir := m.gdir
			m.walk(func(gx, gy int, gdir Direction) bool {
				if gdir != olddir {
					i := 10*(gy*m.height+gx) + int(olddir)
					if history[i] {
						answer++
						return false
					} else {
						history[i] = true
						olddir = gdir
					}
				}
				return true
			})
		}
	}
	return answer
}

func (AoC) SolveDay6() {
	fmt.Printf("Part 1: %d\n", solveDay6Part1())
	fmt.Printf("Part 2: %d\n", solveDay6Part2())
}
