package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// duplicate returns the first value seen twice when summing the inputs from delta.
func duplicate(delta []int) int {
	seen := make(map[int]struct{})
	var freq int

	for {
		for _, d := range delta {
			freq += d
			if _, ok := seen[freq]; ok {
				return freq
			}
			seen[freq] = struct{}{}
		}
	}
}

func main() {
	var delta []int

	f, err := os.Open("input")
	if err != nil {
		log.Fatalln("cannot open input file:", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var i int
		if _, err = fmt.Sscanf(s.Text(), "%d", &i); err != nil {
			log.Fatalln("could not parse line:", err)
		}

		delta = append(delta, i)
	}
	if err = s.Err(); err != nil {
		log.Fatalln("cannot read input file:", err)
	}

	fmt.Println("answer:", duplicate(delta))
}
