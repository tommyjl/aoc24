package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed day12.txt
var inputDay12 string

type day12 struct {
	width  int
	height int
	m      []byte
	v      []bool
}

func newDay12(input string) *day12 {
	width := strings.IndexByte(input, '\n')
	height := strings.Count(input, "\n")
	if input[len(input)-1] != '\n' {
		height++
	}

	size := width * height
	m := make([]byte, size)
	v := make([]bool, size)

	i := 0
	for _, ch := range input {
		if ch == '\n' {
			continue
		}
		m[i] = byte(ch)
		i++
	}

	return &day12{
		width:  width,
		height: height,
		m:      m,
		v:      v,
	}
}

func (d *day12) regions() [][]vec2 {
	var regions [][]vec2

	for y := range d.height {
		for x := range d.width {
			if d.v[y*d.width+x] {
				continue
			}
			next := d.region(x, y)
			regions = append(regions, next)
		}
	}

	return regions
}

func (d *day12) region(x, y int) []vec2 {
	var region []vec2

	region = append(region, vec2{x, y})
	d.v[y*d.width+x] = true
	label := d.m[y*d.width+x]
	region = d.fillRegion(region, label, vec2{x, y})

	return region
}

func getNeighbours(pos vec2) []vec2 {
	return []vec2{
		pos.sub(vec2{0, 1}),
		pos.add(vec2{1, 0}),
		pos.add(vec2{0, 1}),
		pos.sub(vec2{1, 0}),
	}
}

func (d *day12) fillRegion(region []vec2, label byte, pos vec2) []vec2 {
	neighbours := getNeighbours(pos)
	if false && pos.x < 5 && pos.y < 5 {
		fmt.Printf("Neighbours (%d,%d): %v\n", pos.x, pos.y, neighbours)
	}

	for _, n := range neighbours {
		idx := n.y*d.width + n.x
		if d.oob(n) || d.v[idx] || d.m[idx] != label {
			continue
		}
		region = append(region, n)
		d.v[n.y*d.width+n.x] = true
		region = d.fillRegion(region, label, n)
	}

	return region
}

func (d *day12) oob(pos vec2) bool {
	return pos.x < 0 || pos.x >= d.width ||
		pos.y < 0 || pos.y >= d.height
}

func solveDay12Part1(input string) int {
	d := newDay12(input)
	price := 0
	for _, region := range d.regions() {
		price += len(region) * d.getPerimeter(region)
	}
	return price
}

func (d *day12) getPerimeter(region []vec2) int {
	perimeter := 0
	this := region[0].y*d.width + region[0].x
	for _, pos := range region {
		for _, npos := range getNeighbours(pos) {
			that := npos.y*d.width + npos.x
			if d.oob(npos) || d.m[this] != d.m[that] {
				perimeter++
			}
		}
	}
	return perimeter
}

func solveDay12Part2(input string) int {
	d := newDay12(input)
	price := 0
	for _, region := range d.regions() {
		price += len(region) * d.getSideCount(region)
	}
	return price
}

func (d *day12) getSideCount(region []vec2) int {
	sideCount := 0
	this := region[0].y*d.width + region[0].x

	sidesByDir := make([][]vec2, 4)
	for _, pos := range region {
		for i, npos := range getNeighbours(pos) {
			that := npos.y*d.width + npos.x
			if d.oob(npos) || d.m[this] != d.m[that] {
				sidesByDir[i] = append(sidesByDir[i], npos)
			}
		}
	}

	for i, side := range sidesByDir {
		slices.SortFunc(side, func(a, b vec2) int {
			if i%2 == 0 {
				if a.x > b.x {
					return 1
				} else {
					return -1
				}
			} else {
				if a.y > b.y {
					return 1
				} else {
					return -1
				}
			}
		})
	}

	for i, side := range sidesByDir {
		axes := make(map[int][]int)
		for _, cur := range side {
			if i%2 == 0 {
				axes[cur.y] = append(axes[cur.y], cur.x)
			} else {
				axes[cur.x] = append(axes[cur.x], cur.y)
			}
		}

		for _, axis := range axes {
			sideCount++
			prev := axis[0]
			for _, cur := range axis[1:] {
				if cur-prev > 1 {
					sideCount++
				}
				prev = cur
			}
		}
	}

	return sideCount
}

func (AoC) SolveDay12() {
	fmt.Printf("Part 1: %d\n", solveDay12Part1(inputDay12))
	fmt.Printf("Part 2: %d\n", solveDay12Part2(inputDay12))
}
