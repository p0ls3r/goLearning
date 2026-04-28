// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf("  %s\n", comma("1234567我"))
	fmt.Printf("  %t\n", anogram("1234567我", "我7654321"))
}

func anogram(s string, s2 string) bool {
	if s == reverse(s2) {
		return true
	}
	return false
}

func comma(s string) string {
	var buf bytes.Buffer
	var n int
	comma := ","
	runes := []rune(s)

	for i := len(runes); i > 0; i-- {
		v := runes[i-1 : i]
		k := []rune(v)

		if n == 3 {
			buf.WriteString(comma)
			n = 0
		}
		n++

		buf.Write(runesToUTF8(k))
		fmt.Println(buf.String())
	}
	return reverse(buf.String())
}

func runesToUTF8(rs []rune) []byte {
	size := 0
	for _, r := range rs {
		size += utf8.RuneLen(r)
	}

	bs := make([]byte, size)

	count := 0
	for _, r := range rs {
		count += utf8.EncodeRune(bs[count:], r)
	}

	return bs
}

func reverseRunes(runes *[]rune) {
	for i, j := 0, len(*runes)-1; i < j; i, j = i+1, j-1 {
		(*runes)[i], (*runes)[j] = (*runes)[j], (*runes)[i]
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
