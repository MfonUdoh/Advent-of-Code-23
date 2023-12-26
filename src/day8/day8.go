package day8

import (
	"bufio"
	"fmt"
	"os"
)

func Day8() {
	file, err := os.Open("./input/day-8-input.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()

	networkText := make([]string, 0)
	scanner.Scan()
	for scanner.Scan() {
		text := scanner.Text()
		networkText = append(networkText, text)
	}

	fmt.Println(instructions)

	network := make(map[string]Pair)

	curr := "AAA"

	for _, line := range networkText {
		key := line[:3]
		left := line[7:10]
		right := line[12:15]

		network[key] = Pair{left: left, right: right}
	}

	counter := 0
	instructionIndex := 0

	for curr != "ZZZ" {
		counter++
		pair := network[curr]
		if instructionIndex < len(instructions) {
			fmt.Println("instruction", instructionIndex)
			curr = pair.getNext(string(instructions[instructionIndex]))
			instructionIndex++
		} else {
			instructionIndex = 0
			fmt.Println("instruction", instructionIndex)
			curr = pair.getNext(string(instructions[instructionIndex]))
			instructionIndex++
		}
	}

}

type Pair struct {
	left  string
	right string
}

func (this Pair) getNext(instruction string) string {
	if instruction == "L" {
		return this.left
	} else {
		return this.right
	}
}
