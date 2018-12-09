package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// claim represents an area of fabric.
type claim struct {
	id   uint
	x, y uint
	w, h uint
}

// parseClaim returns a new claim parsed from the input string.
func parseClaim(input string) (*claim, error) {
	c := claim{}

	if _, err := fmt.Sscanf(
		input, "#%d @ %d,%d: %dx%d",
		&c.id, &c.x, &c.y, &c.w, &c.h,
	); err != nil {
		return nil, err
	}

	return &c, nil
}

// coord is a point on the fabric.
type coord struct {
	x, y uint
}

// fabric represents the set of claims.
type fabric struct {
	claims map[coord]uint
}

// claim claims points on the fabric for the area in claim c.
func (f *fabric) claim(c *claim) {
	for x := uint(0); x < c.w; x++ {
		for y := uint(0); y < c.h; y++ {
			f.claims[coord{x: c.x + x, y: c.y + y}]++
		}
	}
}

// overlap returns the amount of fabric claimed by >= n claims.
func (f *fabric) overlap(n uint) uint {
	var count uint
	for _, c := range f.claims {
		if c >= n {
			count++
		}
	}
	return count
}

// newFabric returns a new fabric.
func newFabric() fabric {
	return fabric{
		claims: make(map[coord]uint),
	}
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalln("cannot open input file:", err)
	}
	defer f.Close()

	fabric := newFabric()

	s := bufio.NewScanner(f)
	for s.Scan() {
		c, err := parseClaim(s.Text())
		if err != nil {
			log.Fatalln("could not parse claim:", err)
		}
		fabric.claim(c)
	}
	if err = s.Err(); err != nil {
		log.Fatalln("cannot read input file:", err)
	}

	fmt.Println("answer:", fabric.overlap(2))
}
