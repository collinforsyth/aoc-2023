package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"

	"github.com/collinforsyth/aoc-2023/util"
)

func main() {
	input := util.HandleInput()
	hands := parseInput(input)
	p1 := part1(hands)
	fmt.Println("Part 1:", p1)
	p2 := part2(hands)
	fmt.Println("Part 2:", p2)
}

type hand struct {
	bid   int
	cards string
}

type winner int

const (
	fiveOfAKind winner = iota
	fourOfAKind
	fullHouse
	threeOfAKind
	twoPair
	onePair
	highCard
)

func part1(hands []hand) int {
	c := 0
	slices.SortFunc(hands, handCmp(false))
	for i := range hands {
		c += hands[i].bid * (i + 1)
	}
	return c
}

func part2(hands []hand) int {
	c := 0
	slices.SortFunc(hands, handCmp(true))
	for i := range hands {
		c += hands[i].bid * (i + 1)
	}

	return c
}

// implement custom cmp function
func handCmp(includeJokers bool) func(a, b hand) int {
	return func(a, b hand) int {
		t1, t2 := winnerType(a.cards, includeJokers), winnerType(b.cards, includeJokers)
		if t1 == t2 {
			if includeJokers {
				return tieBreaker(a.cards, b.cards, 0)
			} else {
				return tieBreaker(a.cards, b.cards, 11)
			}
		}
		return int(t2 - t1)

	}
}

func cardRk(c byte, jVal int) int {
	switch c {
	case 'T':
		return 10
	case 'J':
		return jVal
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return int(c - '0')
	}
}

type pair struct {
	c byte
	n int
}

func winnerType(cards string, includeJoker bool) winner {
	m := make(map[byte]int)
	for i := 0; i < len(cards); i++ {
		m[cards[i]]++
	}
	// pull out jokers if possible
	jokerCount := 0
	if includeJoker {
		jokerCount = m['J']
		delete(m, 'J')
	}
	p := make([]pair, 0, len(m))
	for k, v := range m {
		p = append(p, pair{k, v})
	}
	slices.SortFunc(p, func(a, b pair) int {
		return b.n - a.n
	})

	if includeJoker && jokerCount > 0 {
		if len(p) == 0 {
			p = append(p, pair{'J', 0})
		}
		// edge case for all jokers:
		// put them back into highest pair
		p[0].n += jokerCount
	}

	wt := highCard
	switch p[0].n {
	case 5:
		wt = fiveOfAKind
	case 4:
		wt = fourOfAKind
	case 3:
		switch p[1].n {
		case 2:
			wt = fullHouse
		default:
			wt = threeOfAKind
		}
	case 2:
		switch p[1].n {
		case 2:
			wt = twoPair
		default:
			wt = onePair
		}
	}

	return wt
}

func tieBreaker(a, b string, jVal int) int {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return cardRk(a[i], jVal) - cardRk(b[i], jVal)
		}
	}
	panic("undefined behavior")
}

func parseInput(input string) []hand {
	hands := make([]hand, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		hands = append(hands, hand{
			bid:   util.MustAtoi(line[1]),
			cards: line[0],
		})
	}
	return hands
}
