package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// claim represents an area of fabric.
type claim struct {
	id   int
	x, y int
	w, h int
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
	x, y int
}

// fabric represents the set of claims, each coordinate having a slice of
// claim ids.
type fabric struct {
	ids    map[int]struct{}
	claims map[coord][]int
}

// claim claims points on the fabric for the area in claim c.
func (f *fabric) claim(c *claim) {
	f.ids[c.id] = struct{}{}
	for x := 0; x < c.w; x++ {
		for y := 0; y < c.h; y++ {
			xy := coord{x: c.x + x, y: c.y + y}
			f.claims[xy] = append(f.claims[xy], c.id)
		}
	}
}

// free returns the id of the claim which does not overlap with any other claims.
func (f *fabric) free() int {
	ids := copyMap(f.ids)

	for _, cids := range f.claims {
		if len(cids) <= 1 {
			continue
		}

		for _, id := range cids {
			delete(ids, id)
		}

		if len(ids) == 1 {
			break
		}
	}

	return pop(ids)
}

// newFabric returns a new fabric.
func newFabric() fabric {
	return fabric{
		ids:    make(map[int]struct{}),
		claims: make(map[coord][]int),
	}
}

// copyMap returns a new map with the values copied from src.
func copyMap(src map[int]struct{}) map[int]struct{} {
	dst := make(map[int]struct{})
	for s := range src {
		dst[s] = struct{}{}
	}

	return dst
}

// pop returns the "first" index from the src map.
// Unreliable for maps with len > 1 as range iterates in an non-deterministic
// order.
func pop(src map[int]struct{}) int {
	for i := range src {
		return i
	}
	return 0
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

	fmt.Println("answer:", fabric.free())
}
