package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day1p2() {
	data, err := os.Open("./input/day-1-input.txt")

	if err != nil {
		fmt.Println("err")
	}

	reader := bufio.NewReader(data)

	scanner := bufio.NewScanner(reader)

	var sum int

	for scanner.Scan() {
		sum = sum + getCalibrationForLineWithStringDigitis(scanner.Bytes())
	}

	fmt.Println(sum)
}

func getCalibrationForLineWithStringDigitis(line []byte) int {
	digits := []Digit{
		{name: "one", value: 1},
		{name: "two", value: 2},
		{name: "three", value: 3},
		{name: "four", value: 4},
		{name: "five", value: 5},
		{name: "six", value: 6},
		{name: "seven", value: 7},
		{name: "eight", value: 8},
		{name: "nine", value: 9},
	}

	firstStringPlace, firstStringDigit := getFirstNumWithPlace(digits, string(line))
	lastStringPlace, lastStringDigit := getLastNumWithPlace(digits, string(line))
	firstIntPlace, firstIntDigit := getFirstIntWithPlace(line)
	lastIntPlace, lastIntDigit := getLastIntWithPlace(line)

	var first string
	var last string

	if firstIntPlace <= firstStringPlace {
		first = firstIntDigit
	} else {
		first = firstStringDigit
	}

	if lastIntPlace >= lastStringPlace {
		last = lastIntDigit
	} else {
		last = lastStringDigit
	}

	combined := first + last

	calibration, _ := strconv.Atoi(combined)

	return calibration
}

func getFirstIntWithPlace(line []byte) (int, string) {
	for i := 0; i < len(line); i++ {
		b := line[i]
		if checkByteIsDigit(b) {
			return i, string(rune(b))
		}
	}
	return len(line), "/"
}

func getLastIntWithPlace(line []byte) (int, string) {
	for i := len(line) - 1; i >= 0; i-- {
		b := line[i]
		if checkByteIsDigit(b) {
			return i, string(rune(b))
		}
	}
	return 0, "/"
}

func getFirstNumWithPlace(digits []Digit, line string) (int, string) {
	firstPlace := len(line)
	strMap := make(map[int]Digit, 0)

	for _, digit := range digits {
		place := strings.Index(line, digit.name)
		if place != -1 {
			strMap[place] = digit
			firstPlace = min(firstPlace, place)
		}
	}

	digit, prs := strMap[firstPlace]

	if prs {
		return firstPlace, strconv.Itoa(digit.value)
	}

	return firstPlace, "*"
}

func getLastNumWithPlace(digits []Digit, line string) (int, string) {
	lastPlace := 0
	strMap := make(map[int]Digit, 0)

	for _, digit := range digits {
		place := strings.LastIndex(line, digit.name)
		if place != -1 {
			strMap[place] = digit
			lastPlace = max(lastPlace, place)
		}
	}

	digit, prs := strMap[lastPlace]

	if prs {
		return lastPlace, strconv.Itoa(digit.value)
	}

	return lastPlace, "*"
}

type Digit struct {
	name  string
	value int
}
