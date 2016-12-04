package main

import (
	"fmt"
)

func validTriangle(a, b, c int) bool {
	fmt.Println("checking", a, b, c)

	if a+b <= c {
		return false
	} else if a+c <= b {
		return false
	} else if b+c <= a {
		return false
	}
	return true
}

func main() {
	var matrix [3][3]int
	valid := 0
	for _, err := fmt.Scanln(&matrix[0][0], &matrix[0][1], &matrix[0][2]); err == nil; _, err = fmt.Scanln(&matrix[0][0], &matrix[0][1], &matrix[0][2]) {
		fmt.Scanln(&matrix[1][0], &matrix[1][1], &matrix[1][2])
		fmt.Scanln(&matrix[2][0], &matrix[2][1], &matrix[2][2])

		for j := 0; j < 3; j++ {
			if validTriangle(matrix[0][j], matrix[1][j], matrix[2][j]) {
				valid++
			}
		}
	}
	fmt.Println(valid)
}
