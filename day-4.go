package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func day4() {
	file, err := os.Open("day-4-input.txt")

	if err != nil {
		panic(err.Error())
	}

	scanner := bufio.NewScanner(file)

	text := make([][]byte, 0)

	for scanner.Scan() {
		src := scanner.Bytes()
		dst := make([]byte, len(src))
		copy(dst, src)
		text = append(text, dst)
	}

	var sum int

	for _, line := range text {
		stripped := stripPrefix(string(line))
		card := parseCard(stripped)

		sum += card.getScore()
	}

	fmt.Println(sum)
}

func stripPrefix(text string) string {
	split := strings.Split(text, ":")

	return split[1]
}

func parseCard(line string) Card {
	trim := strings.TrimSpace(line)
	leftRight := strings.Split(trim, "|")

	leftStr := strings.TrimSpace(leftRight[0])
	rightStr := strings.TrimSpace(leftRight[1])

	left := strings.Split(leftStr, " ")
	right := strings.Split(rightStr, " ")

	leftFilter := filterEmpty(left)
	rightFilter := filterEmpty(right)

	return Card{
		left:  strToInt(leftFilter),
		right: strToInt(rightFilter),
	}
}

func strToInt(nums []string) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		x, err := strconv.Atoi(n)
		if err != nil {
			panic(fmt.Sprintf("Failed to convert %s", n))
		}
		result[i] = x
	}
	return result
}

type Card struct {
	left  []int
	right []int
}

func (this Card) getScore() int {
	matches := this.findMatches()

	return int(math.Pow(2, float64(len(matches)-1)))
}

func (this Card) findMatches() []int {
	matches := make([]int, 0)
	for _, v := range this.right {
		if contains(this.left, v) {
			matches = append(matches, v)
		}
	}
	return matches
}

func contains(hay []int, needle int) bool {
	for _, v := range hay {
		if needle == v {
			return true
		}
	}
	return false
}

func filterEmpty(str []string) []string {
	result := make([]string, 0)
	for _, v := range str {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}
