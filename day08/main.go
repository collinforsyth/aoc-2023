package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/collinforsyth/aoc-2023/util"
)

func main() {
	input := util.HandleInput()
	instructions := parseInput(input)
	p1 := part1(instructions)
	fmt.Println("Part 1:", p1)
	p2 := part2(instructions)
	fmt.Println("Part 2:", p2)

}

type direction struct {
	left, right, parent string
}

type instructions struct {
	lr string
	m  map[string]direction
}

func part1(instr instructions) int {
	return find("AAA", func(s string) bool { return s == "ZZZ" }, instr)
}

func part2(instr instructions) int {
	sp := make([]string, 0)
	for _, v := range instr.m {
		if v.parent[2] == 'A' {
			sp = append(sp, v.parent)
		}
	}

	cycles := make([]int, len(sp))
	for i, v := range sp {
		cycles[i] = find(v, func(s string) bool { return s[2] == 'Z' }, instr)
	}

	return lcm(cycles...)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a ...int) int {
	if len(a) == 1 {
		return a[0]
	} else if len(a) == 2 {
		return a[0] * a[1] / gcd(a[0], a[1])
	}
	return lcm(a[0], lcm(a[1:]...))
}

func find(start string, endFn func(string) bool, instr instructions) int {
	c := 0
	var next direction = instr.m[start]
	found := false
	for !found {
		for i := 0; i < len(instr.lr); i++ {
			switch instr.lr[i] {
			case 'R':
				next = instr.m[next.right]
			case 'L':
				next = instr.m[next.left]
			}
			c++
		}
		if endFn(next.parent) {
			found = true
		}
	}
	return c

}

func parseInput(input string) instructions {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	lr := scanner.Text()
	scanner.Scan()
	m := make(map[string]direction)
	for scanner.Scan() {
		l := strings.Fields(scanner.Text())
		m[l[0]] = direction{
			left:   strings.Trim(l[2], "(,"),
			right:  strings.Trim(l[3], ")"),
			parent: l[0],
		}
	}

	return instructions{
		lr: lr,
		m:  m,
	}
}
