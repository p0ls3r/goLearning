package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	r := 0
	for i := 0; i < 8; i++ {
		r += int(pc[byte(x>>(i*8))])
	}

	return r
}

func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}

func main() {
	//fmt.Printf("pc: %v\n", pc)
	fmt.Printf("pc: %v\n", byte(1&6))
}
