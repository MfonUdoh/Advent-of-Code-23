package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Day3() {
	file, _ := os.Open("./input/day-3-input.txt")

	scanner := bufio.NewScanner(file)

	text := make([][]byte, 0)

	for scanner.Scan() {
		src := scanner.Bytes()
		dst := make([]byte, len(src))
		copy(dst, src)
		text = append(text, dst)
	}

	matrix := Matrix{bytes: text}

	var sum int

	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			if isDigit(text[i][j]) {
				numbers := collectNumberIndices(text[i], j)
				addNumber := false
				for _, n := range numbers {
					if isAdjacentToSymbol(matrix, i, n) {
						addNumber = true
						break
					}
				}
				if addNumber {
					sum += collectNumbers(matrix, i, numbers)
				}
				lastNumber := numbers[len(numbers)-1]
				j = lastNumber
			}
		}
	}

	fmt.Println(sum)
}

func collectNumbers(m Matrix, row int, cols []int) int {
	result := ""
	for _, c := range cols {
		b, _ := m.get(row, c)
		result += string(b)
	}

	num, _ := strconv.Atoi(result)
	return num
}

func collectNumberIndices(line []byte, start int) []int {
	indices := make([]int, 0)

	for i := start; i < len(line); i++ {
		if isDigit(line[i]) {
			indices = append(indices, i)
		} else {
			break
		}
	}
	return indices
}

func isDigit(b byte) bool {
	return unicode.IsDigit(rune(b))
}

func isSymbol(b byte) bool {
	return !isDigit(b) && b != 46
}

func isAdjacentToSymbol(matrix Matrix, i int, j int) bool {
	tl, ok := matrix.get(i-1, j-1)
	if ok {
		if isSymbol(tl) {
			return true
		}
	}

	t, ok := matrix.get(i-1, j)
	if ok {
		if isSymbol(t) {
			return true
		}
	}

	tr, ok := matrix.get(i-1, j+1)
	if ok {
		if isSymbol(tr) {
			return true
		}
	}

	l, ok := matrix.get(i, j-1)
	if ok {
		if isSymbol(l) {
			return true
		}
	}

	r, ok := matrix.get(i, j+1)
	if ok {
		if isSymbol(r) {
			return true
		}
	}
	bl, ok := matrix.get(i+1, j-1)
	if ok {
		if isSymbol(bl) {
			return true
		}
	}

	b, ok := matrix.get(i+1, j)
	if ok {
		if isSymbol(b) {
			return true
		}
	}

	br, ok := matrix.get(i+1, j+1)
	if ok {
		if isSymbol(br) {
			return true
		}
	}
	return false
}

type Matrix struct {
	bytes [][]byte
}

func (this Matrix) get(i int, j int) (byte, bool) {
	if i >= 0 && i < len(this.bytes) {
		if j >= 0 && j < len(this.bytes[i]) {
			return this.bytes[i][j], true
		}
	}
	return 0, false
}
