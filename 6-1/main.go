package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord struct {
	x, y int
}

// parseCoord returns a new coordinate parsed from the input string.
func parseCoord(input string) (*coord, error) {
	var c coord
	if _, err := fmt.Sscanf(input, "%d, %d", &c.x, &c.y); err != nil {
		return nil, err
	}

	return &c, nil
}

// grid returns the dimensions (w, h) by getting the max x, y values from coords.
func grid(coords []*coord) (w, h int) {
	for _, c := range coords {
		if c.y > h {
			h = c.y
		}
		if c.x > w {
			w = c.x
		}
	}
	return w, h
}

// abs returns the absolute representation of integer i.
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// manhattanDist returns the manhattan distance between two coords a and b
// Calculated using the taxicab geometry formula.
// See: https://en.wikipedia.org/wiki/Taxicab_geometry
func manhattanDist(a, b *coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

// coordAtEdge returns true if coordinate c has one of it's axis points touch
// against the edge of a grid as defined by w and h.
func coordAtEdge(c *coord, w, h int) bool {
	return c.x == 0 || c.y == 0 || c.x == w || c.y == h
}

// maxFinite returns the size of the maximum finite area within coords. Areas
// are groups of grid coordinates referenced against the closest point in coords.
// Infinite areas are bound against the edge of the "visible" grid, but extend
// out into infinity. Finite areas are bounded by other areas.
func maxFinite(coords []*coord) int {
	infLocs := make(map[*coord]struct{}) // Locations with infinite areas.
	locs := make(map[*coord]int)         // Locations with finite areas.

	w, h := grid(coords)

	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			var (
				gc  = &coord{x, y} // Current grid coordinate.
				cc  *coord         // Closest star coordinate.
				min = -1
			)
			for _, c := range coords {
				if dist := manhattanDist(gc, c); dist < min || min == -1 {
					min = dist
					cc = c
				} else if dist == min {
					cc = &coord{-1, -1}
				}
			}

			if coordAtEdge(gc, w, h) {
				infLocs[cc] = struct{}{}
			}
			locs[cc]++
		}
	}

	var max int
	for c, a := range locs {
		// Find finite coords with areas bigger than max.
		if _, found := infLocs[c]; a > max && !found {
			max = a
		}
	}

	return max
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalln("cannot open input file:", err)
	}
	defer f.Close()

	var coords []*coord
	s := bufio.NewScanner(f)
	for s.Scan() {
		c, err := parseCoord(s.Text())
		if err != nil {
			log.Fatalln("cannot parse coord:", err)
		}
		coords = append(coords, c)
	}
	if err = s.Err(); err != nil {
		log.Fatalln("cannot read input file:", err)
	}

	fmt.Println("answer:", maxFinite(coords))
}
