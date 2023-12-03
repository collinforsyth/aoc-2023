package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/collinforsyth/aoc-2023/util"
)

func main() {
	input := util.HandleInput()
	games, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p1 := part1(games, 12, 13, 14)
	fmt.Println("Part 1:", p1)
	p2 := part2(games)
	fmt.Println("Part 2:", p2)
}

func part1(games []game, maxRed, maxGreen, maxBlue int) int {
	c := 0
	for _, g := range games {
		possible := true
		for _, s := range g.subsets {
			if !isSubsetValid(s, maxRed, maxGreen, maxBlue) {
				possible = false
			}
		}
		if possible {
			c += g.id
		}
	}
	return c
}

func part2(games []game) int {
	c := 0
	for _, g := range games {
		r, g, b := minCubes(g.subsets)
		c += r * g * b
	}
	return c
}

func minCubes(subsets []subset) (red int, green int, blue int) {
	for _, s := range subsets {
		red = max(red, s.red)
		green = max(green, s.green)
		blue = max(blue, s.blue)
	}
	return

}

func isSubsetValid(s subset, maxRed, maxGreen, maxBlue int) bool {
	return s.red <= maxRed && s.green <= maxGreen && s.blue <= maxBlue
}

type game struct {
	id      int
	subsets []subset
}

type subset struct {
	blue  int
	red   int
	green int
}

func parseInput(input string) ([]game, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	games := make([]game, 0)
	for scanner.Scan() {
		t := scanner.Text()
		s1 := strings.Split(t, ":")
		gameInput, subsetInput := s1[0], s1[1]
		gameID, _ := strconv.Atoi(strings.Split(gameInput, " ")[1])
		g := game{id: gameID}
		for _, subset := range strings.Split(subsetInput, ";") {
			s := parseSubset(subset)
			g.subsets = append(g.subsets, s)
		}
		games = append(games, g)
	}
	return games, nil
}

func parseSubset(ss string) subset {
	subset := subset{}
	for _, pull := range strings.Split(ss, ",") {
		s := strings.Split(pull, " ")
		switch s[2] {
		case "blue":
			subset.blue, _ = strconv.Atoi(s[1])
		case "red":
			subset.red, _ = strconv.Atoi(s[1])
		case "green":
			subset.green, _ = strconv.Atoi(s[1])
		}
	}
	return subset
}
