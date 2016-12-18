package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		ip := scanner.Text()
		// odd-numbered sequences are hypernet sequences
		var seq []string
		s := 0
		for i := range ip {
			if ip[i] == '[' || ip[i] == ']' {
				seq = append(seq, ip[s:i])
				s = i
			}
		}
		seq = append(seq, ip[s:])

		// iterate through all sequences to find ABBAs
		hasAbba := false
		for i := range seq {
			l := len(seq[i])
			last := [3]byte{seq[i][0], seq[i][1], seq[i][2]}
			for j := 3; j < l; j++ {
				if last[0] == seq[i][j] && last[1] == last[2] && last[0] != last[1] {
					if i%2 == 0 {
						hasAbba = true
						break
					} else {
						hasAbba = false
						goto end
					}
				}

				last[0], last[1], last[2] = last[1], last[2], seq[i][j]
			}
		}
	end:
		if hasAbba {
			fmt.Println(ip, "supports TLS")
			count++
		} else {
			fmt.Println(ip, "does NOT support TLS")
		}
	}
	fmt.Println(count)
}
