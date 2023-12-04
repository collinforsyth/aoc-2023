package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"slices"

	"github.com/collinforsyth/aoc-2023/util"
)

func main() {
	input := util.HandleInput()
	arr := createArray(input)
	p1 := part1(arr)
	fmt.Println("Part 1:", p1)
	p2 := part2(arr)
	fmt.Println("Part 2:", p2)
}

func part1(arr []string) int {
	r := regexp.MustCompile(`[0-9]+`)
	c := 0
	for y := 0; y < len(arr); y++ {
		digitIdxs := r.FindAllStringIndex(arr[y], -1)
		for _, i := range digitIdxs {
			isTouchingSymbol := false
			for x := i[0]; x < i[1]; x++ {
				if isAdjacent(arr, x, y) {
					isTouchingSymbol = true
				}
			}
			if isTouchingSymbol {
				s, _ := strconv.Atoi(arr[y][i[0]:i[1]])
				c += s
			}
		}
	}
	return c
}

func part2(arr []string) int {
	c := 0
	digitIdxs := make([][][]int, len(arr))
	r := regexp.MustCompile(`[0-9]+`)
	for y := range arr {
		// create index of all digit locations
		di := r.FindAllStringIndex(arr[y], -1)
		digitIdxs[y] = di
	}
	for y := 0; y < len(arr); y++ {
		for x := 0; x < len(arr[y]); x++ {
			if arr[y][x] == '*' {
				adj := adjacentDigits(arr, digitIdxs, x, y)
				if len(adj) == 2 {
					c += adj[0] * adj[1]
				}
			}
		}
	}
	return c
}

func adjacentDigits(arr []string, digitIdxs [][][]int, currX, currY int) []int {
	m := make(map[int]struct{})
	upDown := []int{-1, 0, 1}
	for _, x := range upDown {
		for _, y := range upDown {
			if i, ok := overlap(arr, digitIdxs, currX+x, currY+y); ok {
				m[i] = struct{}{}
			}
		}
	}
	ret := make([]int, 0, len(m))
	for v := range m {
		ret = append(ret, v)

	}
	return ret
}

func overlap(arr []string, digitIdxs [][][]int, x, y int) (int, bool) {
	// clamp check
	if x < 0 || y < 0 || y >= len(arr) || x >= len(arr[y]) {
		return -1, false
	}
	v, ok := slices.BinarySearchFunc(digitIdxs[y], []int{x}, func(a []int, j []int) int {
		if a[1] > j[0] && a[0] > j[0] {
			return 1
		} else if a[0] <= j[0] && a[1] > j[0] {
			return 0
		} else {
			return -1
		}
	})
	if ok {
		num, _ := strconv.Atoi(arr[y][digitIdxs[y][v][0]:digitIdxs[y][v][1]])
		return num, true
	}
	return -1, false
}

func createArray(input string) []string {
	arr := make([]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	// turn input text into 2d array
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	return arr
}

func isAdjacent(arr []string, currX, currY int) bool {
	upDown := []int{-1, 0, 1}
	for _, x := range upDown {
		for _, y := range upDown {
			if isSymbol(arr, currX+x, currY+y) {
				return true
			}
		}
	}
	return false
}

func isSymbol(arr []string, x, y int) bool {
	// clamp check
	if x < 0 || y < 0 || y >= len(arr) || x >= len(arr[y]) {
		return false
	}
	return ((arr[y][x] < '0' && arr[y][x] != '.') || arr[y][x] > '9')
}
