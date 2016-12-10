package main

import (
	"fmt"
	"regexp"
	"os"
	"bufio"
	"strconv"
)

var opRect = regexp.MustCompile("rect ([0-9]*)x([0-9]*)")
var opRotCol = regexp.MustCompile("rotate column x=([0-9]*) by ([0-9]*)")
var opRotRow = regexp.MustCompile("rotate row y=([0-9]*) by ([0-9]*)")

func main() {
	var screen [50][6]bool
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		op := scanner.Text()
		
		if rect := opRect.FindStringSubmatch(op); rect != nil {
			w, _ := strconv.Atoi(rect[1])
			h, _ := strconv.Atoi(rect[2])
			for i := 0; i < w; i++ {
				for j := 0; j < h; j++ {
					screen[i][j] = true
				}
			}
		} else if col := opRotCol.FindStringSubmatch(op); col != nil {
			i, _ := strconv.Atoi(col[1])
			amt, _ := strconv.Atoi(col[2])
			for n := 0; n < amt; n++ {
				for j := 5; j >= 1; j-- {
					screen[i][j], screen[i][j-1] = screen[i][j-1], screen[i][j]
				}
			}
		} else if row := opRotRow.FindStringSubmatch(op); row != nil {
			j, _ := strconv.Atoi(row[1])
			amt, _ := strconv.Atoi(row[2])
			for n := 0; n < amt; n++ {
				for i := 49; i >= 1; i-- {
					screen[i][j], screen[i-1][j] = screen[i-1][j], screen[i][j]
				}
			}
		}
	}
	count := 0
	for j := range screen[0] {
		for i := range screen {
			if screen[i][j] {
				fmt.Printf("#")
				count++
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println(count)
}
