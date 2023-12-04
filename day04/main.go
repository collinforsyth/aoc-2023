package main

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/collinforsyth/aoc-2023/util"
)

func main() {
	input := util.HandleInput()
	cards := parseInput(input)
	p1 := part1(cards)
	fmt.Println("Part 1:", p1)
	p2 := part2(cards)
	fmt.Println("Part 2:", p2)
}

type card struct {
	number  int
	winners []int
	input   []int
}

func parseInput(input string) []card {
	cards := make([]card, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		s1 := strings.Split(scanner.Text(), ":")
		cardID, _ := strconv.Atoi(strings.Fields(s1[0])[1])
		c := card{number: cardID}
		s2 := strings.Split(s1[1], "|")
		for _, i := range strings.Fields(s2[0]) {
			n, _ := strconv.Atoi(i)
			c.winners = append(c.winners, n)
		}
		for _, i := range strings.Fields(s2[1]) {
			n, _ := strconv.Atoi(i)
			c.input = append(c.input, n)
		}
		cards = append(cards, c)
	}
	return cards
}

func part1(cards []card) int {
	p1 := 0
	for _, c := range cards {
		res := 0
		n := numWinners(c)
		if n > 0 {
			res += int(math.Exp2(float64(n - 1)))
		}
		p1 += res
	}
	return p1
}

func part2(cards []card) int {
	p2 := 0
	copies := make([]int, len(cards))
	for i := range copies {
		copies[i] = 1
	}
	for i, c := range cards {
		for l := 0; l < copies[i]; l++ {
			n := numWinners(c)
			if n == 0 {
				continue
			}
			for j := i + 1; j < i+n+1; j++ {
				copies[j]++
			}
		}
	}
	for _, c := range copies {
		p2 += c
	}
	return p2
}

func numWinners(c card) int {
	i := 0
	for _, n := range c.input {
		if isWinner(c.winners, n) {
			i++
		}
	}
	return i
}

func isWinner(winners []int, x int) bool {
	return slices.Contains(winners, x)
}
