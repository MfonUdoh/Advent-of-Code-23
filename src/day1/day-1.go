package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Day1() {
	data, err := os.Open("./input/day-1-input.txt")

	if err != nil {
		fmt.Println("err")
	}

	reader := bufio.NewReader(data)

	scanner := bufio.NewScanner(reader)

	var sum int

	for scanner.Scan() {
		sum = sum + getCalibrationForLine(scanner.Bytes())
	}

	// _ = scanner
	// getCalibrationForLine([]byte("9ninejcngjshghz"))

	fmt.Println(sum)
}

func getCalibrationForLine(line []byte) int {
	first := string(findFirstDigit(line))
	last := string(findLastDigit(line))

	combined := first + last

	calibration, _ := strconv.Atoi(combined)

	return calibration
}

func findFirstDigit(line []byte) rune {
	for i := 0; i < len(line); i++ {
		b := line[i]
		if checkByteIsDigit(b) {
			return rune(b)
		}
	}
	return 0
}

func findLastDigit(line []byte) rune {
	for i := len(line) - 1; i >= 0; i-- {
		b := line[i]
		if checkByteIsDigit(b) {
			return rune(b)
		}
	}
	return 0
}

func checkByteIsDigit(b byte) bool {
	return unicode.IsDigit(rune(b))
}
