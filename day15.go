package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strings"
)

func (v vec2) moveDir(dir Direction) vec2 {
	switch dir {
	case N:
		return v.sub(vec2{0, 1})
	case E:
		return v.add(vec2{1, 0})
	case S:
		return v.add(vec2{0, 1})
	case W:
		return v.sub(vec2{1, 0})
	default:
		panic("unexpected dir")
	}
}

func (v vec2) outOfBounds(width, height int) bool {
	return v.x < 0 || v.x >= width ||
		v.y < 0 || v.y >= height
}

func (dir Direction) String() string {
	switch dir {
	case N:
		return "N"
	case E:
		return "E"
	case S:
		return "S"
	case W:
		return "W"
	default:
		return fmt.Sprintf("%d", dir)
	}
}

//go:embed day15.txt
var inputDay15 string

type day15tile int

const (
	freeTile day15tile = iota
	wallTile
	boxTile
	boxLTile
	boxRTile
)

type day15 struct {
	height int
	width  int
	tiles  []day15tile
	robot  vec2
	moves  []Direction
	isP2   bool
}

func newDay15(input string, width, height int) *day15 {
	tiles := make([]day15tile, width*height)
	robot := vec2{0, 0}

	rd := bufio.NewReader(strings.NewReader(input))

	y := 0
	for {
		bytes, _, err := rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if len(bytes) == 0 {
			break
		}

		for x, ch := range bytes {
			switch ch {
			case '.':
				tiles[y*width+x] = freeTile
			case 'O':
				tiles[y*width+x] = boxTile
			case '#':
				tiles[y*width+x] = wallTile
			case '@':
				robot = vec2{x, y}
			}
		}

		y++
	}

	var moves []Direction
	for {
		bytes, _, err := rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		for _, byte := range bytes {
			switch byte {
			case '^':
				moves = append(moves, N)
			case '>':
				moves = append(moves, E)
			case 'v':
				moves = append(moves, S)
			case '<':
				moves = append(moves, W)
			case '\n':
				continue
			default:
				panic(fmt.Sprintf("unexpected move: '%c'", byte))
			}
		}
	}

	return &day15{
		width:  width,
		height: height,
		tiles:  tiles,
		robot:  robot,
		moves:  moves,
	}
}

func (day *day15) toP2() {
	tiles := make([]day15tile, len(day.tiles)*2)

	for i, tile := range day.tiles {
		switch tile {
		case freeTile:
			tiles[2*i] = freeTile
			tiles[2*i+1] = freeTile
		case boxTile:
			tiles[2*i] = boxLTile
			tiles[2*i+1] = boxRTile
		case wallTile:
			tiles[2*i] = wallTile
			tiles[2*i+1] = wallTile
		}
	}

	day.robot.x *= 2
	day.width *= 2
	day.tiles = tiles
	day.isP2 = true
}

func (day *day15) moveRobot(dir Direction) {
	nextPos := day.robot.moveDir(dir)
	if day.getTile(nextPos) == freeTile || day.moveTile(nextPos, dir) {
		day.robot = nextPos
	}
}

func (day *day15) moveRobotAll() {
	for _, move := range day.moves {
		day.moveRobot(move)
	}
}

func (day *day15) getTile(pos vec2) day15tile {
	return day.tiles[pos.y*day.width+pos.x]
}

func (day *day15) setTile(pos vec2, tile day15tile) {
	day.tiles[pos.y*day.width+pos.x] = tile
}

func (day *day15) moveTile(pos vec2, dir Direction) bool {
	switch day.getTile(pos) {
	case freeTile:
		return true

	case wallTile:
		return false

	case boxLTile:
		if dir == N || dir == S {
			return day.moveLBoxTile(pos, dir)
		} else {
			return day.moveBoxTile(pos, dir)
		}

	case boxRTile:
		if dir == N || dir == S {
			return day.moveLBoxTile(pos.sub(vec2{1, 0}), dir)
		} else {
			return day.moveBoxTile(pos, dir)
		}

	case boxTile:
		return day.moveBoxTile(pos, dir)

	default:
		return false
	}
}

func (day *day15) moveBoxTile(pos vec2, dir Direction) bool {
	nextPos := pos.moveDir(dir)
	if day.moveTile(nextPos, dir) {
		day.setTile(nextPos, day.getTile(pos))
		day.setTile(pos, freeTile)
		return true
	} else {
		return false
	}
}

func (day *day15) moveLBoxTile(pos vec2, dir Direction) bool {
	if day.getTile(pos.moveDir(dir)) == wallTile ||
		day.getTile(pos.moveDir(dir).moveDir(E)) == wallTile {
		return false
	}

	broken := false
	ls := []vec2{pos}

	for i := 0; i < len(ls); i++ {
		switch next := ls[i].moveDir(dir); day.getTile(next) {
		case boxLTile:
			ls = append(ls, next)
		case boxRTile:
			ls = append(ls, next.moveDir(W))
		case wallTile:
			broken = true
		}

		switch next := ls[i].moveDir(E).moveDir(dir); day.getTile(next) {
		case boxLTile:
			ls = append(ls, next)
		case wallTile:
			broken = true
		}
	}

	if broken {
		return false
	}

	for i := len(ls) - 1; i >= 0; i-- {
		day.setTile(ls[i], freeTile)
		day.setTile(ls[i].moveDir(dir), boxLTile)
		day.setTile(ls[i].moveDir(E), freeTile)
		day.setTile(ls[i].moveDir(E).moveDir(dir), boxRTile)
	}

	return true
}

func (day *day15) calculateAnswer() int {
	tot := 0
	for y := range day.height {
		for x := range day.width {
			switch day.getTile(vec2{x, y}) {
			case boxTile, boxLTile:
				tot += 100*y + x
			}
		}
	}
	return tot
}

func (AoC) SolveDay15() {
	width := strings.IndexByte(inputDay15, '\n')
	height := strings.Index(inputDay15, "\n\n") / width

	day := newDay15(inputDay15, width, height)
	day.moveRobotAll()
	fmt.Printf("Part 1: %d\n", day.calculateAnswer())

	day = newDay15(inputDay15, width, height)
	day.toP2()
	day.moveRobotAll()
	fmt.Printf("Part 2: %d\n", day.calculateAnswer())
}
