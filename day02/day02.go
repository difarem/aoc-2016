package main

import (
	"bufio"
	"fmt"
	"os"
)

func coordsToButton(x, y int) int {
	return x + y*3 + 1
}

func normalizeCoords(x, y int) (int, int) {
	return normalize(x), normalize(y)
}

func normalize(c int) int {
	if c < 0 {
		return 0
	} else if c > 2 {
		return 2
	}
	return c
}

func main() {
	x, y := 1, 1
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, r := range line {
			switch r {
			case 'U':
				y--
			case 'D':
				y++
			case 'R':
				x++
			case 'L':
				x--
			}
			x, y = normalizeCoords(x, y)
		}
		fmt.Printf("%d\n", coordsToButton(x, y))
	}
}
