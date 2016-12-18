package main

import (
	"crypto/md5"
	"fmt"
)

type State struct {
	x, y int
	code string
}

func (st State) hash() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(st.code)))[0:4]
}

func valid(c byte) bool {
	switch c {
	case 'b':
		fallthrough
	case 'c':
		fallthrough
	case 'd':
		fallthrough
	case 'e':
		fallthrough
	case 'f':
		return true
	}
	return false
}

func main() {
	var passcode string
	fmt.Scanln(&passcode)

	var queue []State
	queue = append(queue, State{
		x: 0, y: 0,
		code: passcode,
	})
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.x < 0 || state.x >= 4 {
			continue
		}
		if state.y < 0 || state.y >= 4 {
			continue
		}
		if state.x == 3 && state.y == 3 {
			// we did it
			// woo
			fmt.Println("reached vault,", len(state.code)-len(passcode), "steps away")
		} else {
			hash := state.hash()

			// up
			if valid(hash[0]) {
				queue = append(queue, State{
					x: state.x, y: state.y - 1,
					code: state.code + "U",
				})
			}
			// down
			if valid(hash[1]) {
				queue = append(queue, State{
					x: state.x, y: state.y + 1,
					code: state.code + "D",
				})
			}
			// left
			if valid(hash[2]) {
				queue = append(queue, State{
					x: state.x - 1, y: state.y,
					code: state.code + "L",
				})
			}
			// right
			if valid(hash[3]) {
				queue = append(queue, State{
					x: state.x + 1, y: state.y,
					code: state.code + "R",
				})
			}
		}
	}
	fmt.Println("done")
}
