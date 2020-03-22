package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var o int64
	s := bufio.NewScanner(f)
	for s.Scan() {
		i, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		n := fuel(i)
		o += n
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(o)
}

func fuel(i int64) int64 {
	o := i/3 - 2
	if o <= 0 {
		return 0
	}
	return o + fuel(o)
}
