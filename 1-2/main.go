package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
		if _, err := fmt.Sscanf(s.Text(), "%d", &i); err != nil {
			log.Fatalln("could not parse line:", err)
		}

		delta = append(delta, i)
	}

	fmt.Println("answer:", duplicate(delta))
}

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
