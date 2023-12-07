package main

import (
	"fmt"
	"strings"

	"github.com/collinforsyth/aoc-2023/util"
)

func main() {
	input := util.HandleInput()
	p1 := part1(input)
	fmt.Println("Part 1:", p1)
	p2 := part2(input)
	fmt.Println("Part 2:", p2)
}

type race struct {
	time int
	dist int
}

func part1(input string) int {
	races := parseInputP1(input)
	c := 1
	for _, r := range races {
		c1 := 0
		for i := 0; i < r.time; i++ {
			if distance(r.time, i) > r.dist {
				c1++
			}
		}
		c *= c1
	}
	return c
}

func part2(input string) int {
	r := parseInputP2(input)
	i := 0
	for i < r.time {
		if distance(r.time, i) > r.dist {
			break
		}
		i++
	}
	c := r.time - (i * 2)
	if r.time%1 == 0 {
		c++
	}
	return c
}

func distance(totalTime, heldTime int) int {
	return heldTime * (totalTime - heldTime)
}

func parseInputP2(input string) race {
	s := strings.Split(input, "\n")
	time := strings.Join(strings.Fields(s[0])[1:], "")
	dist := strings.Join(strings.Fields(s[1])[1:], "")
	return race{time: util.MustAtoi(time), dist: util.MustAtoi(dist)}
}

func parseInputP1(input string) []race {
	s := strings.Split(input, "\n")
	time := strings.Fields(s[0])[1:]
	dist := strings.Fields(s[1])[1:]
	r := make([]race, len(time))
	for i := 0; i < len(time); i++ {
		r[i] = race{time: util.MustAtoi(time[i]), dist: util.MustAtoi(dist[i])}
	}
	return r
}
