package main

import (
	"os"
	"piscine"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return
	}
	if a <= 0 || b <= 0 {
		return
	}
	piscine.QuadD(a, b)
}
