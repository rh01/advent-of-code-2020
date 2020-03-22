package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	textByte, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	nums := strings.Split(string(textByte), ",")
	cells := make([]int, len(nums))
	for i, num := range nums {
		cells[i], err = strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			log.Fatalf("Couldn't parse number  %q, %v", num, err)
		}
	}

	cells[1] = 12
	cells[2] = 2
	c := computer{cells: cells}
	for !c.done {
		fmt.Println(c.nextInst, c.cells)
		if err := c.next(); err != nil {
			log.Fatal(err)
		}
	}
}

type computer struct {
	cells    []int
	nextInst int
	done     bool
}

const (
	opADD  int = 1
	opMult int = 2
	opHalt int = 99
)

func (c *computer) read(pos int) int {
	return c.cells[pos]
}

func (c *computer) write(pos int, val int) {
	c.cells[pos] = val
}

func (c *computer) next() error {
	op := c.cells[c.nextInst]
	switch op {
	case opADD:
		args := c.cells[c.nextInst+1 : c.nextInst+4]
		a := c.read(args[0])
		b := c.read(args[1])
		c.write(args[2], a+b)
		c.nextInst += 4
	case opMult:
		args := c.cells[c.nextInst+1 : c.nextInst+4]
		a := c.read(args[0])
		b := c.read(args[1])
		c.write(args[2], a*b)
		c.nextInst += 4
	case opHalt:
		c.done = true
		return nil
	default:
		// c.done = true
		return fmt.Errorf("unknown opcode %d", op)
	}

	return nil
}
