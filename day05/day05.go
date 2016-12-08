package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	var inp string
	fmt.Scanf("%s\n", &inp)
	cn := 0
	for i := 0; ; i++ {
		hashed := fmt.Sprintf("%s%d", inp, i)
		hash := md5.Sum([]byte(hashed))
		hex := fmt.Sprintf("%x", hash)
		if hex[0] == '0' && hex[1] == '0' && hex[2] == '0' && hex[3] == '0' && hex[4] == '0' {
			fmt.Printf("%c", hex[5])
			cn++
		}

		if cn == 8 {
			break
		}
	}
	fmt.Printf("\n")
}
