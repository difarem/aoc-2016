package main

import (
	"fmt"
)

func validTriangle(a, b, c int) bool {
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
	var a, b, c int
	valid := 0
	for _, err := fmt.Scanln(&a, &b, &c); err == nil; _, err = fmt.Scanln(&a, &b, &c) {
		if validTriangle(a, b, c) {
			fmt.Println("valid")
			valid++
		}
	}
	fmt.Println(valid)
}
