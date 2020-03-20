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
	var o int64
	for i/3-2 >= 0 {
		i = i/3 - 2
		o += i
	}
	return o
}
