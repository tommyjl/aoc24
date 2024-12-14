package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed day11.txt
var inputDay11 string

type day11stone struct {
	value int
	next  []int
}

type day11 struct {
	input  []int
	stones map[int]*day11stone
	counts map[vec2]int
}

func newDay11(input string) *day11 {
	stones := make(map[int]*day11stone)
	counts := make(map[vec2]int)
	result := &day11{stones: stones, counts: counts}

	parts := strings.Split(input, " ")
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}

		result.input = append(result.input, num)
		result.insert(num)
	}

	return result
}

func (d *day11) insert(value int) {
	if _, ok := d.stones[value]; !ok {
		d.stones[value] = &day11stone{value: value}
	}
}

func (d *day11) blink() {
	var stones []*day11stone
	for _, stone := range d.stones {
		stones = append(stones, stone)
	}

	for _, stone := range stones {
		if len(stone.next) != 0 {
			continue
		}

		size := int(math.Log10(float64(stone.value)) + 1.0)
		if stone.value == 0 {
			stone.next = append(stone.next, 1)
		} else if size%2 == 0 {
			str := strconv.Itoa(stone.value)
			lhs, _ := strconv.Atoi(str[:size/2])
			rhs, _ := strconv.Atoi(str[size/2:])
			stone.next = append(stone.next, lhs)
			stone.next = append(stone.next, rhs)
		} else {
			stone.next = append(stone.next, stone.value*2024)
		}

		for _, next := range stone.next {
			if _, ok := d.stones[next]; !ok {
				d.stones[next] = &day11stone{value: next}
			}
		}
	}
}

func (d *day11) count(depth int) int {
	tot := len(d.input)
	for _, num := range d.input {
		stone := d.stones[num]
		tot += d.countStone(stone, depth)
	}
	return tot
}

func (d *day11) countStone(stone *day11stone, depth int) int {
	if depth == 0 {
		return 0
	}

	if count, ok := d.counts[vec2{stone.value, depth}]; ok {
		return count
	}

	count := 0
	if len(stone.next) == 2 {
		count++
	}
	for _, next := range stone.next {
		count += d.countStone(d.stones[next], depth-1)
	}

	d.counts[vec2{stone.value, depth}] = count

	return count
}

func solveDay11(input string, blinkCount int) int {
	d := newDay11(input)
	for range blinkCount {
		d.blink()
	}

	return d.count(blinkCount)
}

func (AoC) SolveDay11() {
	fmt.Printf("Part 1: %d\n", solveDay11(inputDay11, 25))
	fmt.Printf("Part 2: %d\n", solveDay11(inputDay11, 75))
}
