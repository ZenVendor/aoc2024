package main

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

type Position struct {
	x int
	y int
}

func translateDir() map[int]Position {
	return map[int]Position{UP: {0, -1}, RIGHT: {1, 0}, DOWN: {0, 1}, LEFT: {-1, 0}}
}

func displayDir() map[int]string {
	return map[int]string{UP: "^", RIGHT: ">", DOWN: "v", LEFT: "<"}
}
