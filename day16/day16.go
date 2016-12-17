package main

import (
	"fmt"
)

func checksum(data []byte) []byte {
	result := make([]byte, len(data)/2)
	for i := range result {
		if data[i*2] == data[i*2+1] {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}
	if len(result)%2 == 0 {
		return checksum(result)
	}
	return result
}

func main() {
	fmt.Printf("disk length: ")
	var diskLen int
	fmt.Scanln(&diskLen)
	var initialState string
	fmt.Scanln(&initialState)
	
	state := []byte(initialState)
	for len(state) < diskLen {
		l := len(state)
		state = append(state, '0')
		for i := l-1; i >= 0; i-- {
			if state[i] == '0' {
				state = append(state, '1')
			} else {
				state = append(state, '0')
			}
		}		
	}
	state = state[:diskLen]
	sum := checksum(state)
	fmt.Printf("checksum is %s\n", sum)
}
