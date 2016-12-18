package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var freq []map[byte]int
	l := 0
	for scanner.Scan() {
		text := scanner.Text()
		if freq == nil {
			l = len(text)
			freq = make([]map[byte]int, l)
			for i := 0; i < l; i++ {
				freq[i] = make(map[byte]int)
			}
		}

		for i := 0; i < l; i++ {
			freq[i][text[i]]++
		}
	}

	for i := 0; i < l; i++ {
		var most byte
		fr := -1
		for b, f := range freq[i] {
			if fr == -1 {
				most = b
				fr = f
			} else if fr > f {
				most = b
				fr = f
			}
		}
		fmt.Printf("%c", most)
	}
	fmt.Printf("\n")
}
