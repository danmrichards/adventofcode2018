package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
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

func replace(input string, char rune) string {
	return strings.Replace(input, string(char), "", -1)
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

	answers := make([]string, 0, 26)
	for i := 'a'; i <= 'z'; i++ {
		answers = append(answers, react(replace(replace(polymer, unicode.ToUpper(i)), i)))
	}

	sort.Slice(answers, func(i, j int) bool {
		return len(answers[i]) < len(answers[j])
	})

	fmt.Println("answer:", len(answers[0]))
}
