package main

import "slices"

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

func translateDir() map[int]Position {
	return map[int]Position{
		UP:    {x: 0, y: -1, value: 0},
		RIGHT: {x: 1, y: 0, value: 1},
		DOWN:  {x: 0, y: 1, value: 2},
		LEFT:  {x: -1, y: 0, value: 3},
	}
}

func displayDir() map[int]string {
	return map[int]string{
		UP:    "^",
		RIGHT: ">",
		DOWN:  "v",
		LEFT:  "<",
	}
}

func arrowToDir() map[string]int {
	return map[string]int{
		"^": UP,
		">": RIGHT,
		"v": DOWN,
		"<": LEFT,
	}
}

type Position struct {
	x     int
	y     int
	value int
}

func (pos Position) OutOfBounds(width, height int) bool {
	if pos.x < 0 || pos.x >= width {
		return true
	}
	if pos.y < 0 || pos.y >= height {
		return true
	}
	return false
}

func (pos Position) Move(width, height int) (newPos Position, ok bool) {
	newPos.x = pos.x + translateDir()[pos.value].x
	newPos.y = pos.y + translateDir()[pos.value].y
	newPos.value = pos.value
	if newPos.OutOfBounds(width, height) {
		return pos, false
	}
	return newPos, true
}

type Positions []Position

func (ps Positions) FirstIndex(pos Position) int {
	return slices.IndexFunc(ps, func(p Position) bool {
		return p.x == pos.x && p.y == pos.y
	})
}

func (ps Positions) FirstValueIndex(val int) int {
	return slices.IndexFunc(ps, func(p Position) bool {
		return p.value == val
	})
}

func (ps Positions) Contains(pos Position) bool {
	return slices.ContainsFunc(ps, func(p Position) bool {
		return p.x == pos.x && p.y == pos.y
	})
}

func (ps Positions) ContainsPair(pair Position) bool {
	return slices.ContainsFunc(ps, func(p Position) bool {
		return (p.x == pair.x && p.y == pair.y) ||
			(p.x == pair.y && p.y == pair.x)
	})
}

func (ps Positions) FindAllWithValue(val int) Positions {
	ps.SortByValueYX()
	var rs Positions
	idx := ps.FirstValueIndex(val)
	for i := idx; i < len(ps); i++ {
		if ps[i].value == val {
			rs = append(rs, ps[i])
		}
	}
	return rs
}

func (ps Positions) SortByValueXY() {
	slices.SortFunc(ps, func(a, b Position) int {
		if n := int(a.value) - int(b.value); n != 0 {
			return n
		}
		if n := a.x - b.x; n != 0 {
			return n
		}
		return a.y - b.y
	})
}
func (ps Positions) SortByValueYX() {
	slices.SortFunc(ps, func(a, b Position) int {
		if n := int(a.value) - int(b.value); n != 0 {
			return n
		}
		if n := a.y - b.y; n != 0 {
			return n
		}
		return a.x - b.x
	})
}
func (ps Positions) SortByXYValue() {
	slices.SortFunc(ps, func(a, b Position) int {
		if n := int(a.x) - int(b.x); n != 0 {
			return n
		}
		if n := a.y - b.y; n != 0 {
			return n
		}
		return a.value - b.value
	})
}
func (ps Positions) SortByYXValue() {
	slices.SortFunc(ps, func(a, b Position) int {
		if n := a.y - b.y; n != 0 {
			return n
		}
		if n := int(a.x) - int(b.x); n != 0 {
			return n
		}
		return a.value - b.value
	})
}

type Plan struct {
	plan     [][]rune
	width    int
	height   int
	startPos Position
	endPos   Position
	objects  Positions
}

func (p *Plan) UpdateDimensions() {
	p.width = len(p.plan[0])
	p.height = len(p.plan)
}
