package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	file, err := os.Open("./input/day-2-input.txt")

	if err != nil {
		fmt.Println("Err:", err.Error())
	}

	scanner := bufio.NewScanner(file)

	var sum int

	for scanner.Scan() {
		sum += addGameIfValid(string(scanner.Bytes()))
	}

	fmt.Println(sum)
}

func addGameIfValid(line string) int {
	split_game := strings.Split(line, ":")

	id_string := split_game[0]

	id, _ := strconv.Atoi(strings.Split(id_string, " ")[1])
	games_string := strings.Split(split_game[1], ";")
	if isGameValid(games_string) {
		return id
	}
	return 0
}

func isGameValid(game_string []string) bool {
	for _, s := range game_string {
		if !isValidSelection(s) {
			return false
		}
	}
	return true
}

func isValidSelection(selection string) bool {

	trimmed := strings.TrimSpace(selection)
	split := strings.Split(trimmed, " ")

	for i := 0; i < len(split); i += 2 {
		num, _ := strconv.Atoi(split[i])
		colour := split[i+1]
		if !isValidColour(num, colour) {
			return false
		}
	}
	return true
}

func isValidColour(num int, col string) bool {
	max_red := 12
	max_green := 13
	max_blue := 14

	colour := strings.Trim(strings.Trim(col, ","), ";")

	max := 0
	switch colour {
	case "red":
		max = max_red
	case "green":
		max = max_green
	case "blue":
		max = max_blue
	}

	return num <= max
}
