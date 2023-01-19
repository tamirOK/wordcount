package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.TrimSpace(os.Args[1])

	if input == "" {
		fmt.Println(0)
		return
	}

	splitted := strings.Split(input, " ")
	fmt.Println(len(splitted))
}
