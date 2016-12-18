package main

import (
	"fmt"
)

const (
	T byte = '^'
	S = '.'
)

func isTrap(l, c, r byte) bool {
	if l == T && c == T && r == S {
		return true
	} else if c == T && r == T && l == S {
		return true
	} else if l == T && c == S && r == S {
		return true
	} else if l == S && c == S && r == T {
		return true
	}
	return false
}

func main() {
	var rows int
	fmt.Printf("rows: ")
	fmt.Scanln(&rows)

	var input string
	fmt.Scanln(&input)

	lastLine := input
	st := 0
	for i := range input {
		if input[i] == S {
			st++
		}
	}
	
	for i := 1; i < rows; i++ {
		var currLine []byte
		for j := range lastLine {
			var left, right byte
			
			if j == 0 {
				left = S
			} else {
				left = lastLine[j-1]
			}
			if j == len(lastLine)-1 {
				right = S
			} else {
				right = lastLine[j+1]
			}
			
			if isTrap(left, lastLine[j], right) {
				currLine = append(currLine, T)
			} else {
				currLine = append(currLine, S)
				st++
			}
		}
		lastLine = string(currLine)
	}
	fmt.Println(st)
}
