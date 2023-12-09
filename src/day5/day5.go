package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day5() {
	fileSeeds, _ := os.Open("./input/day-5-seeds.txt")
	scannerSeeds := bufio.NewScanner(fileSeeds)
	scannerSeeds.Scan()
	seedsBytes := scannerSeeds.Bytes()
	textSeeds := make([]byte, len(seedsBytes))

	copy(textSeeds, seedsBytes)

	fileMaps, _ := os.Open("./input/day-5-maps.txt")
	scannerMaps := bufio.NewScanner(fileMaps)
	textMaps := make([][]byte, 0)

	for scannerMaps.Scan() {
		src := scannerMaps.Bytes()
		dst := make([]byte, len(src))
		copy(dst, src)
		textMaps = append(textMaps, dst)
	}

	seeds := getSeeds(textSeeds)

	seed2soilMap := createMaps(textMaps, 1, 48)
	soil2fertMap := createMaps(textMaps, 51, 83)
	fert2waterMap := createMaps(textMaps, 86, 95)
	water2lightMap := createMaps(textMaps, 98, 109)
	light2tempMap := createMaps(textMaps, 112, 134)
	temp2humMap := createMaps(textMaps, 137, 174)
	hum2locMap := createMaps(textMaps, 177, 210)

	locations := make([]int, len(seeds))

	for i, seed := range seeds {
		soil := useMaps(seed2soilMap, seed)
		fert := useMaps(soil2fertMap, soil)
		water := useMaps(fert2waterMap, fert)
		light := useMaps(water2lightMap, water)
		temp := useMaps(light2tempMap, light)
		hum := useMaps(temp2humMap, temp)
		loc := useMaps(hum2locMap, hum)

		locations[i] = loc
	}

	fmt.Println(getMin(locations))

}

func getMin(locations []int) int {
	minVal := locations[0]
	for _, v := range locations {
		minVal = min(minVal, v)
	}
	return minVal
}

func getSeeds(line []byte) []int {
	seedsBytes := string(line)
	seedsStrings := strings.Split(seedsBytes, " ")
	seeds := make([]int, len(seedsStrings))
	for i := 0; i < len(seedsStrings); i++ {
		seed, _ := strconv.Atoi(seedsStrings[i])
		seeds[i] = seed
	}
	return seeds
}

func useMaps(maps []AlmanacMap, x int) int {
	for _, m := range maps {
		v, ok := m.mapValue(x)
		if ok {
			return v
		}
	}
	return x
}

func createMaps(lines [][]byte, start int, end int) []AlmanacMap {
	maps := make([]AlmanacMap, 0)
	for i := start; i <= end; i++ {
		aMap := createMap(string(lines[i]))
		maps = append(maps, aMap)
	}
	return maps
}

func createMap(line string) AlmanacMap {
	split := strings.Split(line, " ")

	from, _ := strconv.Atoi(split[1])
	to, _ := strconv.Atoi(split[0])
	span, _ := strconv.Atoi(split[2])

	return AlmanacMap{
		from: from,
		to:   to,
		span: span,
	}
}

type AlmanacMap struct {
	from int
	to   int
	span int
}

func (this AlmanacMap) isInRange(x int) bool {
	if x >= this.from && x <= (this.from+this.span) {
		return true
	}
	return false
}

func (this AlmanacMap) mapValue(x int) (int, bool) {
	if this.isInRange(x) {
		offset := x - this.from

		return (this.to + offset), true
	}
	return 0, false
}
