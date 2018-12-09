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

	var output int

	s := bufio.NewScanner(f)
	for s.Scan() {
		var i int
		if _, err = fmt.Sscanf(s.Text(), "%d", &i); err != nil {
			log.Fatalln("could not parse line:", err)
		}

		output += i
	}
	if err = s.Err(); err != nil {
		log.Fatalln("cannot read input file:", err)
	}

	fmt.Println("answer:", output)
}
