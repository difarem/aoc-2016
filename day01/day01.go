package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	N = 0
	W = 1
	S = 2
	E = 3
)

func left(d int) int {
	return (d + 1) % 4
}

func right(d int) int {
	return (d + 3) % 4
}

func advance(d int, q int, x int, y int) (x2, y2 int) {
	switch d {
	case N:
		return x, y - q
	case W:
		return x + q, y
	case S:
		return x, y + q
	case E:
		return x - q, y
	}
	return x, y
}

var visited = make(map[int]map[int]int)

func visit(x, y int) {
	if visited[x] == nil {
		visited[x] = make(map[int]int)
	}
	visited[x][y]++
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	directions := strings.Split(text, ", ")
	x, y := 0, 0
	visit(0, 0)
	facing := N
	for _, direction := range directions {
		d := direction[0]
		n, _ := strconv.Atoi(direction[1:])

		if d == 'R' {
			facing = right(facing)
		} else {
			facing = left(facing)
		}
		for i := 0; i < n; i++ {
			x, y = advance(facing, 1, x, y)
			visit(x, y)
			if visited[x][y] == 2 {
				fmt.Printf("VISITED %d %d TWICE (%d blocks away)\n", x, y, abs(x)+abs(y))
			}
		}
	}
	fmt.Printf("final position: %d %d (%d blocks away)\n", x, y, abs(x)+abs(y))
}
