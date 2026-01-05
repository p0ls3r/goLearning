package main

import (
	"bufio"
	"fmt"
	"os"
)

type countsInfo struct {
	count     int
	files     map[string]bool
	fileNames string
}

func (ci *countsInfo) filesString() {
	sep := ""
	for line, _ := range ci.files {
		ci.fileNames += sep + line
		sep = ", "
	}
}

func main() {
	counts := make(map[string]countsInfo)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n.count > 1 {
			n.filesString()
			println(n.count, n.fileNames, line)
		}
	}
}

func countLines(f *os.File, counts map[string]countsInfo) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, exists := counts[input.Text()]; !exists {
			counts[input.Text()] = countsInfo{count: 1, files: make(map[string]bool)}
			counts[input.Text()].files[f.Name()] = true
		} else {
			info := counts[input.Text()]
			info.count++
			info.files[f.Name()] = true
			counts[input.Text()] = info
		}
	}
}
