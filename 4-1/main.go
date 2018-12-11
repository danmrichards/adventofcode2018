package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	dre = regexp.MustCompile(`\[([^\]]+)\]`) // Match event date.
	gre = regexp.MustCompile(`Guard #(\d+)`) // Match guard ID.
)

const (
	dateFormat = "2006-01-02 15:04"

	eventBeginShift eventTyp = iota
	eventFallAsleep
	eventWakeUp
)

// eventType is the type of event.
type eventTyp int

// event represents a timed event when a guard did something.
type event struct {
	guard     int
	typ       eventTyp
	timestamp time.Time
}

// parseEvent returns a new event parsed from the input string.
func parseEvent(input string) (*event, error) {
	var (
		e   = &event{}
		err error
	)
	if rawTime := dre.FindStringSubmatch(input); len(rawTime) == 2 && rawTime[1] != "" {
		e.timestamp, err = time.Parse(dateFormat, rawTime[1])
		if err != nil {
			return nil, err
		}
	}
	input = strings.TrimSpace(dre.ReplaceAllString(input, ""))

	if rawGuard := gre.FindStringSubmatch(input); len(rawGuard) == 2 && rawGuard[1] != "" {
		e.guard, err = strconv.Atoi(rawGuard[1])
		if err != nil {
			return nil, err
		}
	}

	rawTyp := strings.TrimSpace(gre.ReplaceAllString(input, ""))
	switch {
	case strings.Contains(rawTyp, "begins shift"):
		e.typ = eventBeginShift
	case strings.Contains(rawTyp, "falls asleep"):
		e.typ = eventFallAsleep
	case strings.Contains(rawTyp, "wakes up"):
		e.typ = eventWakeUp
	}

	return e, nil
}

// sleepGuard returns the ID of the guard who slept the most.
func sleepyGuard(events []*event) int {
	// Build a map of guard ids to minutes slept.
	guardSleep := make(map[int]time.Duration)
	for i, e := range events {
		if e.typ == eventWakeUp {
			if events[i-1].typ != eventFallAsleep {
				log.Printf("somehow guard %d woke up from not being asleep\n", events[i-1].guard)
			}
			guardSleep[e.guard] += e.timestamp.Sub(events[i-1].timestamp)
		}
	}

	// Extract the guard who slept the most.
	var (
		lastGuard int
		sleepMax  time.Duration
	)
	for guard, sleep := range guardSleep {
		if sleep > sleepMax {
			sleepMax = sleep
			lastGuard = guard
		}
	}

	return lastGuard
}

// mostSlept returns the minute at which a guard slept the most during the events.
func mostSlept(events []*event, guard int) int {
	mins := make([]int, 60)
	for i, e := range events {
		if e.guard != guard || e.typ != eventWakeUp {
			continue
		}
		for i := events[i-1].timestamp.Minute(); i < e.timestamp.Minute(); i++ {
			mins[i]++
		}
	}

	var min, count int
	for i, c := range mins {
		if c > count {
			min = i
			count = c
		}
	}

	return min
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalln("cannot open input file:", err)
	}
	defer f.Close()

	var rawEvents []string

	s := bufio.NewScanner(f)
	for s.Scan() {
		rawEvents = append(rawEvents, s.Text())
	}
	if err = s.Err(); err != nil {
		log.Fatalln("cannot read input file:", err)
	}

	// Put events in chronological order.
	sort.Slice(rawEvents, func(i, j int) bool {
		var (
			timeI, timeJ time.Time
			err          error
		)

		if rawI := dre.FindStringSubmatch(rawEvents[i]); len(rawI) == 2 && rawI[1] != "" {
			timeI, err = time.Parse(dateFormat, rawI[1])
			if err != nil {
				log.Fatalln(err)
			}
		}

		if rawJ := dre.FindStringSubmatch(rawEvents[j]); len(rawJ) == 2 && rawJ[1] != "" {
			timeJ, err = time.Parse(dateFormat, rawJ[1])
			if err != nil {
				log.Fatalln(err)
			}
		}

		return timeI.Before(timeJ)
	})

	var (
		events    = make([]*event, 0, len(rawEvents))
		lastGuard int
	)
	for _, re := range rawEvents {
		e, err := parseEvent(re)
		if err != nil {
			log.Fatalln(err)
		} else if e.guard == 0 {
			e.guard = lastGuard
		}

		events = append(events, e)
		lastGuard = e.guard
	}

	g := sleepyGuard(events)
	m := mostSlept(events, g)

	fmt.Println("answer:", g*m)
}
