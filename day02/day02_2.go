package main

import (
	"bufio"
	"fmt"
	"os"
)

const keypad = "..1...234.56789.ABC...D.."

func coordsToButton(x, y int) byte {
	return keypad[x+y*5]
}

func validCoords(x, y int) bool {
	if x < 0 || x >= 5 || y < 0 || y >= 5 {
		return false
	}
	return keypad[x+y*5] != '.'
}

func move(x, y, dx, dy int) (int, int) {
	if validCoords(x+dx, y+dy) {
		return x + dx, y + dy
	}
	return x, y
}

func main() {
	x, y := 0, 2
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, r := range line {
			switch r {
			case 'U':
				x, y = move(x, y, 0, -1)
			case 'D':
				x, y = move(x, y, 0, 1)
			case 'R':
				x, y = move(x, y, 1, 0)
			case 'L':
				x, y = move(x, y, -1, 0)
			}
		}
		fmt.Printf("%c\n", coordsToButton(x, y))
	}
}
