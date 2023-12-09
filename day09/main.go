package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/collinforsyth/aoc-2023/util"
)

func main() {
	input := util.HandleInput()
	sequences := parseInput(input)
	fmt.Println("Part 1:", part1(sequences))
	fmt.Println("Part 2:", part2(sequences))
}

func part1(sequences [][]int) int {
	c := 0
	for _, seq := range sequences {
		// c += sequenceDegree(seq)
		c += sequenceDegree(seq, func(seq []int, d int) int {
			return seq[len(seq)-1] + d
		})

	}
	return c
}

func part2(sequences [][]int) int {
	c := 0
	for _, seq := range sequences {
		c += sequenceDegree(seq, func(seq []int, d int) int {
			return seq[0] - d
		})
	}
	return c
}

func sequenceDegree(sequence []int, adder func(seq []int, d int) int) int {
	if constant(sequence) {
		return sequence[0]
	}
	diff := make([]int, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		diff[i] = sequence[i+1] - sequence[i]
	}
	d := sequenceDegree(diff, adder)
	return adder(sequence, d)
}

func constant(a []int) bool {
	if len(a) <= 1 {
		return true
	}
	d1 := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] != d1 {
			return false
		}
	}
	return true

}

func parseInput(input string) [][]int {
	history := [][]int{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		a := make([]int, len(nums))
		for i := range nums {
			a[i] = util.MustAtoi(nums[i])
		}
		history = append(history, a)
	}
	return history
}
