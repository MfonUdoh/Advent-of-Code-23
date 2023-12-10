package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day6p2() {
	file, _ := os.Open("./input/day-6-input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timeBytes := scanner.Bytes()
	timeText := string(timeBytes)
	scanner.Scan()
	distanceBytes := scanner.Bytes()
	distanceText := string(distanceBytes)

	time := getNoKernDistanceTimes(timeText)
	distance := getNoKernDistanceTimes(distanceText)

	race := Race{
		duration: time,
		record:   distance,
	}

	fmt.Println(race.getVariations())

}

func getNoKernDistanceTimes(line string) int {
	splitLine := strings.Split(line, " ")
	resultString := ""

	for _, item := range splitLine {
		_, err := strconv.Atoi(item)
		if err == nil {
			resultString += item
		}
	}

	result, _ := strconv.Atoi(resultString)
	return result
}
