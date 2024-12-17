package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func atoi(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return result
}

//go:embed day14.txt
var inputDay14 string

type day14robot struct {
	p vec2
	v vec2
}

type day14 struct {
	height int
	width  int
	robots []day14robot
}

func readDay14Input(input string, width, height int) day14 {
	result := day14{
		width:  width,
		height: height,
	}

	r := bufio.NewReader(strings.NewReader(input))

	for {
		bytes, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		line := string(bytes)

		sep := strings.IndexByte(line, ' ')
		p := strings.Split(line[2:sep], ",")
		v := strings.Split(line[sep+3:], ",")

		d := day14robot{
			p: vec2{atoi(p[0]), atoi(p[1])},
			v: vec2{atoi(v[0]), atoi(v[1])},
		}

		result.robots = append(result.robots, d)
	}

	return result
}

func (d *day14) moveAll(count int) {
	for i := range d.robots {
		d.move(i, count)
	}
}

func (d *day14) move(idx int, count int) {
	pos := d.robots[idx].p
	v := d.robots[idx].v

	pos.x = (pos.x + count*(d.width+v.x)) % d.width
	pos.y = (pos.y + count*(d.height+v.y)) % d.height

	d.robots[idx].p = pos
}

func (d *day14) safetyFactor() int {
	var q1, q2, q3, q4 int

	midw := d.width / 2
	midh := d.height / 2

	for _, robot := range d.robots {
		switch {
		case robot.p.x < midw && robot.p.y < midh:
			q1++
		case robot.p.x > midw && robot.p.y < midh:
			q2++
		case robot.p.x > midw && robot.p.y > midh:
			q3++
		case robot.p.x < midw && robot.p.y > midh:
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func (d *day14) image() {
	counts := make(map[vec2]int)
	for _, robot := range d.robots {
		counts[robot.p]++
	}
	for y := range d.height {
		for x := range d.width {
			count := counts[vec2{x, y}]
			if count == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", count)
			}
		}
		fmt.Println()
	}
}

func (AoC) SolveDay14() {
	// Actual sizes:
	d := readDay14Input(inputDay14, 101, 103)
	d.moveAll(100)
	fmt.Printf("Part 1: %d\n", d.safetyFactor())

	// I had no idea what to look for here, so I printed the output until I
	// could find a pattern. Starting on 11 seconds, you could see all the
	// counts gather in the middle of the image, and then the same thing
	// would reoccur every 202 seconds. Eventually, a christmas tree
	// appeared on the screen.
	p2 := 11 + 202*38
	fmt.Printf("Part 2: %d\n", p2)
	d = readDay14Input(inputDay14, 101, 103)
	d.moveAll(p2)
	d.image()
}
