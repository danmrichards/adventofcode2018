package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func react(input string) string {
	var i int
	for {
		// Figure out the matching polymer element.
		r := rune(input[i])
		var opposite rune
		if unicode.IsLower(r) {
			opposite = unicode.ToUpper(r)
		} else {
			opposite = unicode.ToLower(r)
		}

		if i+1 >= len(input) {
			break
		}

		// If there is no match carry on.
		if rune(input[i+1]) != opposite {
			i++
			continue
		}

		// Remove the matches and start again.
		input = input[:i] + input[i+2:]
		i = 0
	}

	return input
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalln("cannot open input file:", err)
	}
	defer f.Close()

	var polymer string
	s := bufio.NewScanner(f)
	for s.Scan() {
		polymer += s.Text()
	}
	if err = s.Err(); err != nil {
		log.Fatalln("cannot read input file:", err)
	}

	polymer = react(polymer)
	fmt.Println("answer:", len(polymer))
}
