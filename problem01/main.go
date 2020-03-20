package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	s := bufio.NewScanner(fd)
	var res int64
	for i := 1; s.Scan(); i++ {
		i, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		f := fuel(i)
		res += f
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func fuel(i int64) int64 {
	return i/3 - 2
}
