package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	var inp string
	fmt.Scanf("%s\n", &inp)
	cn := 0
	code := []byte("________")
	for i := 0; ; i++ {
		hashed := fmt.Sprintf("%s%d", inp, i)
		hash := md5.Sum([]byte(hashed))
		hex := fmt.Sprintf("%x", hash)
		if hex[0] == '0' && hex[1] == '0' && hex[2] == '0' && hex[3] == '0' && hex[4] == '0' {
			pos := int(hex[5]) - '0'
			char := hex[6]
			if pos < 0 || pos >= 8 {
				continue
			} else if code[pos] != '_' {
				continue
			}
			code[pos] = char
			fmt.Println(string(code))
			cn++
		}

		if cn == 8 {
			break
		}
	}
}
