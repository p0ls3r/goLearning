// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

//!+

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
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
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]

		if arg == "sha256" {
			c1 := sha256.Sum256([]byte("x"))
			c2 := sha256.Sum256([]byte("X"))
			fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
			fmt.Printf("pop count: %d\n", PopCountSHA256(c1, c2))
		}

		if arg == "sha384" {
			c1 := sha512.Sum384([]byte("x"))
			c2 := sha512.Sum384([]byte("X"))
			fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
		}

		if arg == "sha512" {
			c1 := sha512.Sum512([]byte("x"))
			c2 := sha512.Sum512([]byte("X"))
			fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
		}
	}
}

func PopCountSHA256(a, b [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		count = count + int(a[i]^b[i])
	}
	return count
}

//!-
