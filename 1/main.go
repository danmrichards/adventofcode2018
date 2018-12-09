package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalln("cannot open input file:", err)
	}
	defer func () {
		if err = f.Close(); err != nil {
			log.Fatalln("cannot close input file:", err)
		}
	}()

	var output int

	s := bufio.NewScanner(f)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalln("could not parse line:", err)
		}

		output += i
	}

	fmt.Println("answer:", output)
}
