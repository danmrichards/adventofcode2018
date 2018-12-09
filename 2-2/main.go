package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// compare returns two strings containing the common and differing runes from
// input strings a and b.
func compare(a, b string) (diff string, common string) {
	d := make([]rune, 0, len(a))
	c := make([]rune, 0, len(a))
	for i := range a {
		if a[i] == b[i] {
			c = append(c, rune(b[i]))
		} else {
			d = append(d, rune(b[i]))
		}
	}

	return string(d), string(c)
}

func main() {
	var ids []string
	f, err := os.Open("input")
	if err != nil {
		log.Fatalln("cannot open input file:", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		ids = append(ids, s.Text())
	}
	if err = s.Err(); err != nil {
		log.Fatalln("cannot read input file:", err)
	}

	sort.Strings(ids)

	for i := range ids {
		d, c := compare(ids[i], ids[i+1])
		if len(d) == 1 {
			fmt.Println("id 1:", ids[i])
			fmt.Println("id 2:", ids[i+1])
			fmt.Println(strings.Repeat("-", len(ids[i])+6))
			fmt.Println("common:", c)
			fmt.Println("diff:", d)
			os.Exit(0)
		}
	}
}
