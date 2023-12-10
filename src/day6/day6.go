package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day6() {
	file, _ := os.Open("./input/day-6-input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timeBytes := scanner.Bytes()
	timeText := string(timeBytes)
	scanner.Scan()
	distanceBytes := scanner.Bytes()
	distanceText := string(distanceBytes)

	times := getDistanceTimes(timeText)
	distances := getDistanceTimes(distanceText)

	races := make([]Race, len(times))

	for i, time := range times {
		races[i] = Race{
			duration: time,
			record:   distances[i],
		}
	}

	ways := make([]int, len(races))
	for i, race := range races {
		ways[i] = race.getVariations()
	}

	fmt.Println(getProduct(ways))

}

func getProduct(ints []int) int {
	product := 1
	for _, x := range ints {
		product = product * x
	}
	return product
}

func getDistanceTimes(line string) []int {
	splitLine := strings.Split(line, " ")
	result := make([]int, 0)

	for _, item := range splitLine {
		num, err := strconv.Atoi(item)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

type Race struct {
	duration int
	record   int
}

func (this Race) getVariations() int {
	variations := make([]int, 0)
	for i := 1; i < this.duration; i++ {
		attemptDistance := this.getDistance(i)
		if this.isRecordBreaking(attemptDistance) {
			variations = append(variations, attemptDistance)
		}
	}
	return len(variations)
}

func (this Race) getDistance(velocity int) int {
	remainingTime := this.duration - velocity

	return velocity * remainingTime
}

func (this Race) isRecordBreaking(distance int) bool {
	return this.record < distance
}
