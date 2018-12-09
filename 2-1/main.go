package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalln("cannot open input file:", err)
	}
	defer f.Close()

	var twos, threes uint

	s := bufio.NewScanner(f)
	for s.Scan() {
		if hasTwo := count(s.Text(), 2); hasTwo {
			twos++
		}
		if hasThree := count(s.Text(), 3); hasThree {
			threes++
		}
	}
	if err = s.Err(); err != nil {
		log.Fatalln("cannot read input file:", err)
	}

	fmt.Println("answer:", twos*threes)
}

// count returns true if id contains exactly n of a character within itself.
// e.g. id = abbcde n = 2 returns true because it has 2 'b'.
func count(id string, n uint) bool {
	runes := make(map[rune]uint)
	for _, char := range id {
		runes[char]++
	}

	for _, rn := range runes {
		if rn == n {
			return true
		}
	}

	return false
}
