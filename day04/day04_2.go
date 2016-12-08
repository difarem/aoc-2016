package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

var re = regexp.MustCompile(`([a-z\-]*)-([0-9]*)\[([a-z]*)\]`)

type letter struct {
	letter   rune
	quantity int
}

// letter sorter
type letters []letter

func (l letters) Len() int      { return len(l) }
func (l letters) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l letters) Less(i, j int) bool {
	if l[i].quantity == l[j].quantity {
		return l[i].letter < l[j].letter
	}
	return l[i].quantity > l[j].quantity
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		roomStr := scanner.Text()
		result := re.FindAllStringSubmatch(roomStr, -1)

		roomName := result[0][1]
		sectorID, _ := strconv.Atoi(result[0][2])
		checksum := result[0][3]

		// find the five most common letters
		var lts letters
		ltsm := make(map[rune]int)
		for _, r := range roomName {
			if r != '-' {
				if i, ok := ltsm[r]; ok {
					lts[i].quantity++
				} else {
					ltsm[r] = len(lts)
					lts = append(lts, letter{letter: r, quantity: 1})
				}
			}
		}
		sort.Sort(lts)
		checksum2 := make([]byte, 5)
		for i := 0; i < 5; i++ {
			checksum2[i] = byte(lts[i].letter)
		}

		if checksum == string(checksum2) {
			// now, decrypt the room name
			decrypted := make([]byte, len(roomName))
			for i := range roomName {
				if roomName[i] == '-' {
					decrypted[i] = ' '
				} else {
					shift := sectorID % ('z' - 'a' + 1)
					decrypted[i] = byte(int(roomName[i])-'a'+shift)%('z'-'a'+1) + 'a'
				}
			}
			fmt.Println(roomName, sectorID, "=>", string(decrypted))
		} else {
			//fmt.Println(roomName, "is fake, actual checksum would be", string(checksum2))
		}
	}
}
