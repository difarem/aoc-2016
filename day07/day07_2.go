package main

import (
	"fmt"
	"bufio"
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
				s = i+1
			}
		}
		seq = append(seq, ip[s:])
		
		// iterate through all sequences to find ABAs and BABs
		var aba []string
		var bab []string
		for i := range seq {
			last := [2]byte{seq[i][0], seq[i][1]}
			l := len(seq[i])
			for j := 2; j < l; j++ {
				if last[0] == seq[i][j] && last[0] != last[1] {
					if i%2 == 0 {
						// supernet
						aba = append(aba, seq[i][j-2:j+1])
					} else {
						// hypernet
						bab = append(bab, seq[i][j-2:j+1])
					}
				}
				last[0], last[1] = last[1], seq[i][j]
			}
		}
		// find a pair
		found := false
		for i := range aba {
			for j := range bab {
				if aba[i][0] == bab[j][1] && aba[i][1] == bab[j][0] {
					found = true
					goto end
				}
			}
		}
		end: if found {
			fmt.Println(ip, "supports SSL")
			count++
		} else {
			fmt.Println(ip, "does NOT support SSL")
		}
	}
	fmt.Println(count)
}
