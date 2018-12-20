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

// regionSize returns the size of the region containing all locations which have
// a total distance to all given coordinates of less than d.
func regionSize(coords []*coord, d int) int {
	w, h := grid(coords)

	var size int
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			var (
				gc   = &coord{x, y} // Current grid coordinate.
				dist int
			)
			for _, c := range coords {
				dist += manhattanDist(gc, c)
			}

			if dist < d {
				size++
			}
		}
	}

	return size
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

	fmt.Println("answer:", regionSize(coords, 10000))
}
