package main

import (
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed day9.txt
var inputDay9 string

type day9 struct {
	blocks []int
}

func newDay9(input string) day9 {
	var blocks []int

	id := 0
	for i, ch := range input {
		size, err := strconv.Atoi(string(ch))
		if err != nil {
			panic(err)
		}

		v := -1
		if i%2 == 0 {
			v = id
			id++
		}

		for range size {
			blocks = append(blocks, v)
		}
	}

	return day9{blocks: blocks}
}

func (d day9) checksum() int {
	answer := 0
	for i, blk := range d.blocks {
		if blk >= 0 {
			answer += i * blk
		}
	}
	return answer
}

func (d day9) compact() day9 {
	rhs := len(d.blocks) - 1
	for lhs, blk := range d.blocks {
		if blk == -1 {
			for d.blocks[rhs] == -1 {
				rhs--
			}
			if lhs > rhs {
				break
			}
			d.blocks[lhs] = d.blocks[rhs]
			d.blocks[rhs] = -1
		}
	}
	return d
}

func (d day9) compact2() day9 {
	for rhs := len(d.blocks) - 1; rhs >= 0; rhs-- {
		if d.blocks[rhs] == -1 {
			continue
		}
		if rhs > 0 && d.blocks[rhs-1] == d.blocks[rhs] {
			continue
		}

		rsize := d.sizeof(rhs)

		for lhs := 0; lhs < rhs; lhs++ {
			lsize := d.sizeof(lhs)
			if d.blocks[lhs] == -1 && lsize >= rsize {
				val := d.blocks[rhs]
				for i := 0; i < rsize; i++ {
					d.blocks[lhs+i] = val
					d.blocks[rhs+i] = -1
				}
				break
			}
		}
	}
	return d
}

func (d day9) sizeof(idx int) int {
	tot := 0
	for i := idx; i < len(d.blocks) && d.blocks[i] == d.blocks[idx]; i++ {
		tot++
	}
	return tot
}

func (AoC) SolveDay9() {
	fmt.Printf("Part 1: %d\n", newDay9(inputDay9).compact().checksum())
	fmt.Printf("Part 2: %d\n", newDay9(inputDay9).compact2().checksum())
}
