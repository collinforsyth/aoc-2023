package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/collinforsyth/aoc-2023/util"
)

func main() {
	input := util.HandleInput()
	p1, err := solve(input, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Part 1:", p1)
	p2, err := solve(input, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Part 2:", p2)
}

func solve(input string, includeSpelled bool) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	i := 0
	for scanner.Scan() {
		t := scanner.Text()
		var f, l int
		var err error
		if f, err = firstDigit(t, includeSpelled); err != nil {
			return -1, err
		}
		if l, err = lastDigit(t, includeSpelled); err != nil {
			return -1, err
		}
		i += (f * 10) + l
	}
	return i, nil
}

func firstDigit(l string, includeSpelled bool) (int, error) {
	for i := 0; i < len(l); i++ {
		if i, ok := digit(l[i]); ok {
			return i, nil
		}
		if includeSpelled {
			if i, ok := spelledDigit(l[i:]); ok {
				return i, nil
			}
		}
	}
	return -1, errors.New("no digit found")
}

func lastDigit(l string, includeSpelled bool) (int, error) {
	for i := len(l) - 1; i >= 0; i-- {
		if i, ok := digit(l[i]); ok {
			return i, nil
		}
		if includeSpelled {
			if i, ok := reverseSpelledDigit(l[:i+1]); ok {
				return i, nil
			}
		}
	}
	return -1, errors.New("no digit found")
}

func digit(b byte) (int, bool) {
	if '0' <= b && b <= '9' {
		i, _ := strconv.Atoi(string(b))
		return i, true
	}
	return -1, false
}

var spelled map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func spelledDigit(s string) (int, bool) {
	for k, v := range spelled {
		if strings.HasPrefix(s, k) {
			return v, true
		}
	}
	return -1, false
}

func reverseSpelledDigit(s string) (int, bool) {
	for k, v := range spelled {
		if strings.HasSuffix(s, k) {
			return v, true
		}
	}
	return -1, false
}
