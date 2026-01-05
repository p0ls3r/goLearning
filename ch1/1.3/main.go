package main

import (
	"fmt"
	"os"
)

func main() {

	for i, arg := range os.Args {
		s := fmt.Sprintf("index: %d, value: %s", i, arg)
		fmt.Println(s)
	}

}
